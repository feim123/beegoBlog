package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var err error
	cate := c.Input().Get("cate")
	c.Data["Topics"], _ = models.GetAllTopic(cate, true)
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Categories"], err = models.GetAllCategory()
	if err != nil{
		beego.Error(err)
	}
	c.TplName = "home.html"
}