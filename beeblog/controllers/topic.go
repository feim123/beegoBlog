package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	var err error
	c.Data["Topics"], err = models.GetAllTopic("", false)
	if err != nil{
		beego.Error(err)
	}
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "topic.html"
}

func (c *TopicController) Post() {
	tid := c.Input().Get("tid")
	if !checkAccount(c.Ctx){
		c.Redirect("/login", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	label := c.Input().Get("label")
	var err error
	if tid != "" {
		err = models.ModifyTopic(tid, title, content, category, label)
	}else {
		err = models.AddTopic(title, content, category, label)
	}
	if err != nil{
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}

func (c *TopicController) Add() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Categories"], _ = models.GetAllCategory()
	c.TplName = "topic_add.html"
}

func (c *TopicController) View() {

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	var err error
	tid := c.Input().Get("id")
	c.TplName = "topic_view.html"

	c.Data["Topic"], err = models.GetTopic(tid)
	if err != nil{
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	comments, err := models.GetAllReply(tid)
	if err != nil{
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Comments"] = comments
}

func (c *TopicController) Modify() {
	tid := c.Input().Get("tid")
	var err error
	c.Data["Topic"], err = models.GetTopic(tid)
	if err != nil{
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Categories"], err = models.GetAllCategory()
	if err != nil{
		beego.Error(err)
	}
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "topic_modify.html"
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx){
		c.Redirect("/login", 302)
		return
	}
	tid := c.Input().Get("tid")
	var err error
	err = models.DeleteTopic(tid)
	if err != nil{
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}