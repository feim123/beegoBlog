package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController)Add()  {

	if !checkAccount(c.Ctx){
		c.Redirect("/login", 302)
		return
	}

	tid := c.Input().Get("tid")
	nickname := c.Input().Get("nickname")
	content := c.Input().Get("content")
	err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view?id="+tid, 302)
}

func (c *ReplyController)Del()  {

	if !checkAccount(c.Ctx){
		c.Redirect("/login", 302)
		return
	}

	cid := c.Input().Get("cid")
	tid := c.Input().Get("tid")
	var err error
	err = models.DelReply(cid, tid)
	if err != nil{
		beego.Error(err)
	}

	c.Redirect("/topic/view?id="+tid, 302)
}
