package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//https://cloud.tencent.com/developer/article/1869961

// 简单使用
func TestSimpleRand(t *testing.T) {
	t.Log("start ...")
	assert := assert.New(t)
	assert.Equal(1, 1)
	assert.NotEqual(1, 2)
	assert.NotNil("123")
	assert.IsType([]string{}, []string{""})

	assert.Contains("Hello World", "World")
	assert.Contains(map[string]string{"Hello": "World"}, "Hello")
	assert.Contains([]string{"Hello", "World"}, "Hello")
	assert.True(true)
	//assert.True(false)
	t.Log("next ...")
	var s []string
	assert.Empty(s)
	assert.Nil(s)
	t.Log("end ...")
}

// 一般用的更多的是表驱动方式把同一个单元的测试用例都放在一起
func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}

func Calculate(num int) int {
	return num + 2
}

// suite套件包
type _Suite struct {
	suite.Suite
}

// SetupSuite() 和 TearDownSuite() 全局只会执行一次
// SetupTest() TearDownTest() BeforeTest() AfterTest() 对套件中的每个测试执行一次
func (s *_Suite) AfterTest(suiteName, testName string) {
	fmt.Printf("AferTest: suiteName=%s,testName=%s\n", suiteName, testName)
}

func (s *_Suite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest: suiteName=%s,testName=%s\n", suiteName, testName)
}

// SetupSuite() 仅执行一次
func (s *_Suite) SetupSuite() {
	fmt.Printf("SetupSuite() ...\n")
}

// TearDownSuite() 仅执行一次
func (s *_Suite) TearDownSuite() {
	fmt.Printf("TearDowmnSuite()...\n")
}

func (s *_Suite) SetupTest() {
	fmt.Printf("SetupTest()... \n")
}

func (s *_Suite) TearDownTest() {
	fmt.Printf("TearDownTest()... \n")
}

func (s *_Suite) TestSimpleRand() {
	fmt.Printf("TestSimpleRand()... \n")
	//ret := SimpleRand(1, 10) // 4.
	//assert.Equal(s.T(), ret, int64(10))
}

func (s *_Suite) TestCalculate() {
	fmt.Printf("TestCalculate()... \n")
	ret := Calculate(1) //9.
	assert.Equal(s.T(), ret, 3)
}

// 让 go test 执行测试
func TestAll(t *testing.T) {
	suite.Run(t, new(_Suite))
}
