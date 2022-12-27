
https://mp.weixin.qq.com/s/p572g5KcSwy2ri40d1cPTg

Context 是如何影响 grpc 通信超时控制的

图片

﻿上述场景是我在实际开发应用过程中抽象出来的 grpc 通信过程，这也是一个较为通用的过程，client 端将带有超时时间的 context 传递到 server 端，server 端在超时时间内需要完成请求处理并返回响应给 client 端，若超过超时请求时间，那么链接将会断开，client 端将不会收到任何响应。

然而在实际开发应用中，发现即便 server 端的 context 超时了，但是其请求响应仍会偶发性地传递到 client 端，导致我们的一个功能出现了不可预期的情况，为了用代码描述对应的交互流程，我在这里放了简化后的示例代码，描述了当时的交互逻辑。﻿﻿
https://github.com/git-qfzhang/hello-golang/tree/master/grpc-go/helloworld

图片

grpc 超时传递流程

在 Golang grpc 通信过程中，超时信息会在不同通信端进行传递的，传递的介质是 Http2 Request Frame。grpc client 在发送请求之前，会将信息封装在不同的的 Frame 中，例如 Data Frame 用来存放请求的 response payload；Header Frame 用户存在一些跨 goroutine 传递的数据，例如路径信息。而超时信息就存放在 Header Frame 中，其源码如下所示：

// NewStream 方法的调用链路：grpc.Invoke -> invoke -> sendRequest -> NewStream

// NewStream creates a stream and register it into the transport as "active"
// streams.
func (t *http2Client) NewStream(ctx context.Context, callHdr *CallHdr) (_ *Stream, err error) {
  // 省略 ...
  
  // HPACK encodes various headers. Note that once WriteField(...) is
  // called, the corresponding headers/continuation frame has to be sent
  // because hpack.Encoder is stateful.
  t.hBuf.Reset()
  t.hEnc.WriteField(hpack.HeaderField{Name: ":method", Value: "POST"})
  t.hEnc.WriteField(hpack.HeaderField{Name: ":scheme", Value: t.scheme})
  t.hEnc.WriteField(hpack.HeaderField{Name: ":path", Value: callHdr.Method})
  t.hEnc.WriteField(hpack.HeaderField{Name: ":authority", Value: callHdr.Host})
  t.hEnc.WriteField(hpack.HeaderField{Name: "content-type", Value: "application/grpc"})
  t.hEnc.WriteField(hpack.HeaderField{Name: "user-agent", Value: t.userAgent})
  t.hEnc.WriteField(hpack.HeaderField{Name: "te", Value: "trailers"})

  if callHdr.SendCompress != "" {
    t.hEnc.WriteField(hpack.HeaderField{Name: "grpc-encoding", Value: callHdr.SendCompress})
  }
  if dl, ok := ctx.Deadline(); ok {
    // Send out timeout regardless its value. The server can detect timeout context by itself.
    timeout := dl.Sub(time.Now())
    t.hEnc.WriteField(hpack.HeaderField{Name: "grpc-timeout", Value: encodeTimeout(timeout)})
  }
  // 省略 ...
}
client server 端在收到超时信息后，将 grpc-timeout 字段从 Header 中取出，基于该超时信息新建一个 context 实例，其源码如下所示：

// processHeaderField 方法调用链：grpc.Server -> handleRawConn -> serveNewHTTP2Transport -> serveStreams -> HandleStreams -> operateHeaders -> processHeaderField
// operateHeader takes action on the decoded headers.
func (t *http2Server) operateHeaders(frame *http2.MetaHeadersFrame, handle func(*Stream)) (close bool) {
  buf := newRecvBuffer()
  s := &Stream{
    id:  frame.Header().StreamID,
    st:  t,
    buf: buf,
    fc:  &inFlow{limit: initialWindowSize},
  }
  var state decodeState
  for _, hf := range frame.Fields {
    state.processHeaderField(hf)
  }
  // 省略 ...
  s.recvCompress = state.encoding
  if state.timeoutSet {
    s.ctx, s.cancel = context.WithTimeout(context.TODO(), state.timeout)
  } else {
    s.ctx, s.cancel = context.WithCancel(context.TODO())
  }
  // 省略 ...
}
func (d *decodeState) processHeaderField(f hpack.HeaderField) {
  switch f.Name {
  // 省略 ...
  case "grpc-timeout":
    d.timeoutSet = true
    var err error
    d.timeout, err = decodeTimeout(f.Value)
    if err != nil {
      d.setErr(streamErrorf(codes.Internal, "transport: malformed time-out: %v", err))
      return
    }
  // 省略 ...
  }
}
在 grpc client 端，会去不断检查 context.Done() 来判断 context 是否超时，若超时，则会断开链接。然而，也会存在 context timeout races 的情况，例如，client 端 context 已经超时，但是此时下一轮检查还未开始，同时 server 端恰好返回了响应信息，此时虽然 client 端 context 超时了，但是仍然会接收到 server 端的响应并处理；更普遍的情况是 select { case <- ctx; ...; case <- response; ...}，这就会导致有 50% 的概率未检测到 context 超时，详情请参考我之前在 grpc-go 中提的 issue。


确保 grpc 响应超时错误

在我之前经历的错误场景中， server 端 context 出现超时，并返回响应给 client 端，此时 client 端预期应该也会超时并断开链接，但实际是会成功接收到 client 端的响应，由于处理逻辑的问题，当时的响应并不包含超时错误，因此 client 端在接收到请求后会重新发送一次请求，重新发送完成后，才检测到 context 超时，最终断开链接，导致了错误的出现。﻿﻿

图片

因此，在应用过程中，需要在 server 端 context timeout 时，保证返回的 response 中的错误信息是 grpc.DeadlineExceeded，让 client 端也感知到 timeout 的发生，避免不必要逻辑的发生。﻿

参考：
[1] https://github.com/grpc/grpc-go﻿
[2] https://github.com/grpc/grpc-go/issues/5206#issuecomment-1058564271﻿
[3] https://xiaomi-info.github.io/2019/12/30/grpc-deadline/


---


