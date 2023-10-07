
package bytesconv
 
import (
"unsafe"
)
 
// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {  
    return *(*[]byte)(unsafe.Pointer(
        &struct {
        string
        Cap int
        }{s, len(s)},
    ))
}
 
// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}
来源于gin框架: https://pkg.go.dev/github.com/gin-gonic/gin/internal/bytesconv

原理:使用unsafe.Pointer 强制转换字符串或者字符数组的类型，从而减少内存占用空间

