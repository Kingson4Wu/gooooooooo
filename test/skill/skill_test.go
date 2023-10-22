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

// 为什么这个函数的返回值会是 -1
func demo1() int {
	ret := -1
	defer func() {
		ret = 1
	}()
	return ret
}

// output: -1

// 为什么这个函数的返回值会是 1
func demo2() (ret int) {
	defer func() {
		ret = 1
	}()
	return ret
}

// output: 1

/**作者：yakumioto
链接：https://juejin.cn/post/6968737029756551176
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func TestXxx(t *testing.T) {
	fmt.Println(demo1())
	fmt.Println(demo2())
}

/**

以上就是两个方法的汇编源码解析，从两个栗子中可以得到结果。
demo1 中的 ret 是临时变量，虽然 defer 确实改了 ret 的值，但这个值跟返回值没一毛钱关系，而且在汇编中 demo1 的返回值在还没调用 demo1.func1 的时候就已经确定了，所以 demo1 返回了 -1。
demo2 中的 ret 则直接指向了返回值的地址，defer 也改了返回值的值， 所以 demo2 就返回了 1

作者：yakumioto
链接：https://juejin.cn/post/6968737029756551176
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
