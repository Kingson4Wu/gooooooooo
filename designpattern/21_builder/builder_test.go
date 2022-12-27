package builder

import (
	"fmt"
	"testing"
)

/**
生成器模式

（一）概念

生成器是一种创建型设计模式，使你能够分步骤创建复杂对象。


与其他创建型模式不同，生成器不要求产品拥有通用接口。这使得用相同的创建过程生成不同的产品成为可能。
*/

func TestBuilder(t *testing.T) {
	pancakeCook := NewPancakeCook(NewNormalPancakeBuilder())
	fmt.Printf("摊一个普通煎饼 %#v\n", pancakeCook.MakePancake())

	pancakeCook.SetPancakeBuilder(NewHealthyPancakeBuilder())
	fmt.Printf("摊一个健康的加量煎饼 %#v\n", pancakeCook.MakeBigPancake())
}
