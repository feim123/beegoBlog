package models


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/Unknwon/com"
	"os"
	"path"
	"github.com/astaxie/beego/orm"
)

const (
	_DB_NAME		= "data/beeblog.db"
	_MYSQL_DRIVER = "mysql"
)

func RegisterDB()  {
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:12345@tcp(120.25.220.178)/beeblog?parseTime=true&loc=Local", 10)
}
