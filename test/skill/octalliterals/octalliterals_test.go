package octalliterals

import (
	"fmt"
	"testing"
)

//在GO中以0或者0o开头的数字被认为是8进制数
/**
file, err := os.OpenFile("foo", os.O_RDONLY, 0o644)
二进制-使用0b或者0B作为前缀（0b100代表4）
16进制-使用0x或者0X前缀（0xF代表15）
虚数-使用i作为后缀（3i）
除此之外我们可以在数字中穿插下划线提升可读性：1_000_000_000
*/

func TestR(t *testing.T) {
	sum := 100 + 010
	fmt.Println(sum)

	// output
	// 108
}
func TestR2(t *testing.T) {
	sum := 100 + 0o10
	fmt.Println(sum)

	// output
	// 108
}
