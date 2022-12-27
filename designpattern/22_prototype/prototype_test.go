package prototype

import (
	"fmt"
	"reflect"
	"testing"
)

/**
原型模式

（一）概念

原型是一种创建型设计模式，使你能够复制对象，甚至是复杂对象，而又无需使代码依赖它们所属的类。


所有的原型类都必须有一个通用的接口， 使得即使在对象所属的具体类未知的情况下也能复制对象。原型对象可以生成自身的完整副本， 因为相同类的对象可以相互访问对方的私有成员变量。



（二）示例

纸质文件可以通过复印机轻松拷贝出多份，设置Paper接口，包含读取文件内容和克隆文件两个方法。同时声明两个类报纸（Newspaper）和简历（Resume）实现了Paper接口，通过复印机（Copier）复印出两类文件的副本，并读取文件副本内容。


*/

func TestPrototype(t *testing.T) {
	copier := NewCopier("云打印机")
	oneNewspaper := NewNewspaper("Go是最好的编程语言", "Go语言十大优势")
	oneResume := NewResume("小明", 29, "5年码农")

	otherNewspaper := copier.copy(oneNewspaper)
	copyNewspaperMsg := make([]byte, 100)
	byteSize, _ := otherNewspaper.Read(copyNewspaperMsg)
	fmt.Println("copyNewspaperMsg:" + string(copyNewspaperMsg[:byteSize]))

	otherResume := copier.copy(oneResume)
	copyResumeMsg := make([]byte, 100)
	byteSize, _ = otherResume.Read(copyResumeMsg)
	fmt.Println("copyResumeMsg:" + string(copyResumeMsg[:byteSize]))
}

// Copier 复印机
type Copier struct {
	name string
}

func NewCopier(n string) *Copier {
	return &Copier{name: n}
}

func (c *Copier) copy(paper Paper) Paper {
	fmt.Printf("copier name:%v is copying:%v ", c.name, reflect.TypeOf(paper).String())
	return paper.Clone()
}
