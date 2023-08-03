package skill

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	// 示例：如果条件满足，返回 x 的值，否则返回 y 的值
	x := 10
	y := 20
	max := ternary(x > y, x, y)
	fmt.Println("The greater number is:", max)

}

// ternary 是一个函数，用于模拟三目运算
func ternary(condition bool, x, y interface{}) interface{} {
	if condition {
		return x
	}
	return y
}

func stringToFixedNumber(input string) uint64 {
	// 创建一个FNV-1哈希对象
	hash := fnv.New64a()

	// 将字符串转换成字节数组并计算哈希值
	hash.Write([]byte(input))

	// 获取哈希值并返回
	return hash.Sum64()
}

func TestIdConvert(t *testing.T) {

	// 要转换的字符串
	str := "hello"

	// 将字符串转换成固定的数字
	num := stringToFixedNumber(str)

	// 打印转换后的数字
	fmt.Println("转换后的数字:", num)
	fmt.Println("type:", reflect.TypeOf(num))
}
