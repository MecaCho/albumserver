package db

import (
	"beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var ORM_ orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "root:Huawei@123@tcp(100.114.234.225:3306)/albumserver?charset=utf8", 10)

	orm.RegisterModel(new(User), new(Album), new(Photo))

	orm.RunSyncdb("default", false, true)
	ORM_ = orm.NewOrm()
}


type User struct {
	IncID     int    `orm:"pk;column(inc_id)"json:"inc_id"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreateAt  string `json:"create_at"`
	UpdatedAt string `json:"updated_at"`
}

type Album struct {
	IncID      int64    `orm:"pk;column(inc_id)";json:"inc_id"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	UploadTime string   `json:"upload_time"`
	Size       int      `json:"size"`
	PhotoNum   int      `json:"photo_num"`
	Photos     []*Photo `orm:"rel(m2m)";json:"photos"`
	UserID     string   `json:"user_id"`
}

type Photo struct {
	IncID      int    `orm:"pk;column(inc_id)";json:"inc_id"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	UploadTime string `json:"upload_time"`
	Size       int    `json:"size"`
	CreateAt   string `json:"create_at"`
	UpdatedAt  string `json:"updated_at"`
	Album      *Album `orm:"rel(fk)";json:"album"`
}
