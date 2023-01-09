package test

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
https://cloud.tencent.com/developer/article/1881962
*/

/*
  gorm.io/gorm 指的是gorm V2版本，详细可参考 https://gorm.io/zh_CN/docs/v2_release_note.html
  github.com/jinzhu/gorm 一般指V1版本
*/

type OrderRepo struct {
	db *gorm.DB
}

// 将gorm.DB作为一个参数，在初始化时赋值：方便测试时，放一个mock的db
func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

// Order针对的是 orders 表中的一行数据
type Order struct {
	Id    int64
	Name  string
	Price float32
}

// OrderFields 作为一个 数据库Order对象+fields字段的组合
// fields用来指定Order中的哪些字段生效
type OrderFields struct {
	order  *Order
	fields []interface{}
}

func NewOrderFields(order *Order, fields []interface{}) *OrderFields {
	return &OrderFields{
		order:  order,
		fields: fields,
	}
}

func (repo *OrderRepo) AddOrder(order *Order) (err error) {
	err = repo.db.Create(order).Error
	return
}

func (repo *OrderRepo) QueryOrders(pageNumber, pageSize int, condition *OrderFields) (orders []Order, err error) {
	db := repo.db
	// condition非nil的话，追加条件
	if condition != nil {
		// 这里的field指定了order中生效的字段，这些字段会被放在SQL的where条件中
		db = db.Where(condition.order, condition.fields...)
	}
	err = db.
		Limit(pageSize).
		Offset((pageNumber - 1) * pageSize).
		Find(&orders).Error
	return
}

func (repo *OrderRepo) UpdateOrder(updated, condition *OrderFields) (err error) {
	if updated == nil || len(updated.fields) == 0 {
		return errors.New("update must choose certain fields")
	} else if condition == nil {
		return errors.New("update must include where condition")
	}

	err = repo.db.
		Model(&Order{}).
		// 这里的field指定了order中被更新的字段
		Select(updated.fields[0], updated.fields[1:]...).
		// 这里的field指定了被更新的where条件中的字段
		Where(condition.order, condition.fields...).
		Updates(updated.order).
		Error
	return
}

// 注意，我们使用的是gorm 2.0，网上很多例子其实是针对1.0的
var (
	DB1  *gorm.DB
	mock sqlmock.Sqlmock
)

// TestMain是在当前package下，最先运行的一个函数，常用于初始化
func TestMain(m *testing.M) {
	var (
		db  *sql.DB
		err error
	)

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	DB1, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// m.Run 是真正调用下面各个Test函数的入口
	os.Exit(m.Run())
}

/*
 sqlmock 对语法限制比较大，下面的sql语句必须精确匹配（包括符号和空格）
*/

func TestOrderRepo_AddOrder(t *testing.T) {
	var order = &Order{Name: "order1", Price: 1.1}
	orderRepo := NewOrderRepo(DB1)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `orders` (`name`,`price`) VALUES (?,?)").
		WithArgs(order.Name, order.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := orderRepo.AddOrder(order)
	assert.Nil(t, err)
}

func TestOrderRepo_QueryOrders(t *testing.T) {
	var orders = []Order{
		{1, "name1", 1.0},
		{2, "name2", 1.0},
	}
	page, size := 2, 10
	orderRepo := NewOrderRepo(DB1)
	condition := NewOrderFields(&Order{Price: 1.0}, []interface{}{"price"})

	mock.ExpectQuery(
		"SELECT * FROM `orders` WHERE `orders`.`price` = ? LIMIT 10 OFFSET 10").
		WithArgs(condition.order.Price).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(orders[0].Id, orders[0].Name, orders[0].Price).
				AddRow(orders[1].Id, orders[1].Name, orders[1].Price))

	ret, err := orderRepo.QueryOrders(page, size, condition)
	assert.Nil(t, err)
	assert.Equal(t, orders, ret)
}

func TestOrderRepo_UpdateOrder(t *testing.T) {
	orderRepo := NewOrderRepo(DB1)
	// 表示要更新的字段为Order对象中的id,name两个字段
	updated := NewOrderFields(&Order{Id: 1, Name: "test_name"}, []interface{}{"id", "name"})
	// 表示更新的条件为Order对象中的price字段
	condition := NewOrderFields(&Order{Price: 1.0}, []interface{}{"price"})

	mock.ExpectBegin()
	mock.ExpectExec(
		"UPDATE `orders` SET `id`=?,`name`=? WHERE `orders`.`price` = ?").
		WithArgs(updated.order.Id, updated.order.Name, condition.order.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := orderRepo.UpdateOrder(updated, condition)
	assert.Nil(t, err)
}
