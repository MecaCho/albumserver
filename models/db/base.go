package db

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var ORM_ orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:Huawei@123@tcp(100.114.234.225:3306)/albumserver?charset=utf8", 10)

	orm.RegisterModel(new(User), new(Album), new(Photo))

	orm.RunSyncdb("default", false, true)
	ORM_ = orm.NewOrm()
}
