package main

import (
    "fmt"
    "os"
    "path/filepath"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    f, err := ioutil.TempFile("", "sample")
    check(err)

    fmt.Println("Temp file name:", f.Name())

    defer os.Remove(f.Name())

    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

    dname, err := ioutil.TempDir("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    fname := filepath.Join(dname, "file1")
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}
//https://github.com/golang/go/issues/42026