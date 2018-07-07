package db

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var ORM_ orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:QWQ920403@ty@tcp(127.0.0.1:3306)/albumserver?charset=utf8", 10)

	orm.RegisterModel(new(User), new(Album), new(Photo))

	orm.RunSyncdb("default", false, true)
	ORM_ = orm.NewOrm()
}
