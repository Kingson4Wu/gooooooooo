package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkHi(b *testing.B) {
	b.ReportAllocs()
	req, err := http.ReadRequest(
		bufio.NewReader(strings.NewReader("GET /hi HTTP/1.0\r\n\r\n")))
	if err != nil {
		b.Fatal(err)
	}
	rw := httptest.NewRecorder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handleHi(rw, req)
	}
}

//go test -v -run=^$ -bench=.
/**

goos: darwin
goarch: amd64
pkg: git.kugou.net/cupid/gooooooooo/command/cmd/hello/pprof
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
BenchmarkHi
BenchmarkHi-4             204402              5313 ns/op
PASS
ok      git.kugou.net/cupid/gooooooooo/command/cmd/hello/pprof  1.750s

*/

func BenchmarkHiParallel(b *testing.B) {
	r, err := http.ReadRequest(bufio.NewReader(strings.NewReader("GET /hi HTTP/1.0\r\n\r\n")))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		rw := httptest.NewRecorder()
		for pb.Next() {
			handleHi(rw, r)
		}
	})
}
