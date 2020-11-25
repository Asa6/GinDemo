package model

type User struct {
	ID       uint // 字段`ID`为默认主键
	UserName string
	Email    string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	PassWord string
	FullName string
}

// 设置User的表名为`users`
func (User) TableName() string {
	return "users"
}
