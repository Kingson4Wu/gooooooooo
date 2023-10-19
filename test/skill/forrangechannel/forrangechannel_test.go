package forrangechannel

import (
	"bytes"
	"fmt"
	"testing"
)

func TestForrangechannel(t *testing.T) {
	var monitorsChan = make(chan *bytes.Buffer, 3)

	/*go func() {
		for {
			select {
			case buf := <-monitorsChan:
				fmt.Println(string(buf.Bytes()))
			}
		}
	}()*/

	go func() {
		for buf := range monitorsChan {
			fmt.Println(buf.String())
		}

	}()

	for i := 0; i < 100; i++ {
		inf := []byte("xx")
		buf := bytes.NewBuffer(inf)
		monitorsChan <- buf
	}

}
