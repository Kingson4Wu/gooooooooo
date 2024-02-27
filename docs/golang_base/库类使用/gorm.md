+ gorm-gen: https://juejin.cn/post/7133150674400837668

https://gorm.io/zh_CN/docs/index.html


### DBResolver
+ https://gorm.io/zh_CN/docs/dbresolver.html
GORM 支持基于策略的 sources/replicas 负载均衡，自定义策略应该是一个实现了以下接口的 struct：

type Policy interface {
    Resolve([]gorm.ConnPool) gorm.ConnPool
}


+ gorm.Clause()子句分析之ON DUPLICATE KEY UPDATE 

result := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "xxx_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"xx", "give_num", "xxxx", "xxx", "xxxx", "remark", "update_time"}),
	}).Create(&rule)

