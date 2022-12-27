package abstractfactory

import "fmt"

// Cook 厨师接口，抽象工厂
type Cook interface {
	// MakeFood 制作主食
	MakeFood() Food
	// MakeDrink 制作饮品
	MakeDrink() Drink
}

// Food 主食接口
type Food interface {
	// Eaten 被吃
	Eaten() string
}

// Drink 饮品接口
type Drink interface {
	// Drunk 被喝
	Drunk() string
}

//-----三餐不同厨师接口的实现

// breakfastCook 早餐厨师
type breakfastCook struct{}

func NewBreakfastCook() *breakfastCook {
	return &breakfastCook{}
}

func (b *breakfastCook) MakeFood() Food {
	return &cakeFood{"切片面包"}
}

func (b *breakfastCook) MakeDrink() Drink {
	return &gruelDrink{"小米粥"}
}

// lunchCook 午餐厨师
type lunchCook struct{}

func NewLunchCook() *lunchCook {
	return &lunchCook{}
}

func (l *lunchCook) MakeFood() Food {
	return &dishFood{"烤全羊"}
}

func (l *lunchCook) MakeDrink() Drink {
	return &sodaDrink{"冰镇可口可乐"}
}

// dinnerCook 晚餐厨师
type dinnerCook struct{}

func NewDinnerCook() *dinnerCook {
	return &dinnerCook{}
}

func (d *dinnerCook) MakeFood() Food {
	return &noodleFood{"大盘鸡拌面"}
}

func (d *dinnerCook) MakeDrink() Drink {
	return &soupDrink{"西红柿鸡蛋汤"}
}

//----不同吃的

// cakeFood 蛋糕
type cakeFood struct {
	cakeName string
}

func (c *cakeFood) Eaten() string {
	return fmt.Sprintf("%v被吃", c.cakeName)
}

// dishFood 菜肴
type dishFood struct {
	dishName string
}

func (d *dishFood) Eaten() string {
	return fmt.Sprintf("%v被吃", d.dishName)
}

// noodleFood 面条
type noodleFood struct {
	noodleName string
}

func (n *noodleFood) Eaten() string {
	return fmt.Sprintf("%v被吃", n.noodleName)
}

//-----不同喝的

// gruelDrink 粥
type gruelDrink struct {
	gruelName string
}

func (g *gruelDrink) Drunk() string {
	return fmt.Sprintf("%v被喝", g.gruelName)
}

// sodaDrink 汽水
type sodaDrink struct {
	sodaName string
}

func (s *sodaDrink) Drunk() string {
	return fmt.Sprintf("%v被喝", s.sodaName)
}

// soupDrink 汤
type soupDrink struct {
	soupName string
}

func (s *soupDrink) Drunk() string {
	return fmt.Sprintf("%v被喝", s.soupName)
}
