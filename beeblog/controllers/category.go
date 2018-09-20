package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController)Get()  {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		err := models.AddCategory(name)
		if err != nil{
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		err := models.DelCategory(id)
		if err != nil{
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}
	c.TplName = "category.html"
	c.Data["IsCategory"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	var err error
	c.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
}