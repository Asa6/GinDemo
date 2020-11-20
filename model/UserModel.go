package model

type User struct {
	ID       int `gorm:"AUTO_INCREMENT"` // 自增
	UserName string
	Email    string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	PassWord string
}

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "users"
}
