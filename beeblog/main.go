package main

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	_ "beeblog/routers"
)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}