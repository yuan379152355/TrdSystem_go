package tables

import (
	//此处注意“_”表示引用mysql函数中init的方法而无需使用函数
	. "TrdSystem/public"
	_ "github.com/go-sql-driver/mysql"
)

type SY_user struct {
	SYBase
	//对应id表字段
	Userid int `db:"userid"`
	//对应name表字段
	Username string `db:"username"`
	//对应passwword表字段
	Password string `db:"password"`
}

func NewSY_user(sYBase SYBase) *SY_user {
	var result = new(SY_user)
	result.SYBase = sYBase
	return result
}

func Insert() {
	//ySql := NewSYSql(sYBase SYBase)
}
