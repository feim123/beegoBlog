package models

import (
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
	"log"
	"strings"
)

type Topic struct {
	Id 				int64
	Uid 			int64
	Category		string
	Title 			string
	Labels			string
	Content 		string 		`orm:"size(5000)"`
	Attachment 		string
	CreatedTime		time.Time	`orm:"index"`
	UpdatedTime		time.Time	`orm:"index"`
	Views			int64		`orm:"index"`
	Author			string
	ReplyTime		time.Time	`orm:"index"`
	ReplyCount		int64
	ReplyLastUserID	int64
}

func AddTopic(title, content, category, label string) error {

	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()
	timeNow := time.Now()
	topic := &Topic{
		Title:			title,
		Content:		content,
		Labels:			label,
		Category:		category,
		CreatedTime:	timeNow,
		UpdatedTime:	timeNow,
		Views:			0,
		ReplyTime:		timeNow}

	_, err := o.Insert(topic)
	if err != nil{
		return err
	}
	//更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("Title", category).One(cate)
	if err == nil{
		cate.TopicCount++
		o.Update(cate)
	}
	return err
}

func GetAllTopic(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc{
		if cate != ""{
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-createdTime").All(&topics)
	}else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		log.Println("get topic error")
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", "", -1)
	return topic, err
}

func ModifyTopic(tid, title, content, category, label string) error {

	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	id, err := strconv.ParseInt(tid, 10, 64)
	o := orm.NewOrm()

	//更新topic时修改分类
	oldTopic := new(Topic)
	oldTopic, err = GetTopic(tid)
	if err != nil{
		log.Println("get old topic error")
	}
	oldCate := oldTopic.Category
	if oldCate != category{
		err = DelCategoryNum(oldCate)
		if err != nil{
			log.Println("invoke DelCategory error")
		}
		err = AddCategoryNum(category)
		if err != nil{
			log.Println("invoke AddCategory error")
		}
	}

	timeNow := time.Now()
	topic := &Topic{Id:	id}
	err = o.Read(topic)
	if err == nil{
		topic.Title = 		title
		topic.Content = 	content
		topic.UpdatedTime = timeNow
		topic.Category =	category
		topic.Labels=		label
		o.Update(topic)
	}
	return err
}

func DeleteTopic(tid string) error {

	o := orm.NewOrm()

	//更新category数
	topic2, err := GetTopic(tid)
	cateName := topic2.Category
	err = DelCategoryNum(cateName)

	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return err
	}
	topic := &Topic{Id:id}
	_, err = o.Delete(topic)

	return err
}

func DelTopicReplyCount(tidNum int64) error {
	var err error
	o := orm.NewOrm()

	replies := make([]*Comment, 0)
	qs := o.QueryTable("Comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-createdTime").All(&replies)
	if err != nil{
		return err
	}

	ReplyTime := time.Now()
	if len(replies) != 0{
		ReplyTime = replies[0].CreatedTime
	}
	topic := &Topic{Id:tidNum}
	if o.Read(topic) == nil{
		topic.ReplyTime = ReplyTime
		topic.ReplyCount--
		_, err = o.Update(topic)
	}
	return err
}

func AddTopicReplyCount(tidNum int64) error {
	var err error
	o := orm.NewOrm()
	topic := &Topic{Id:tidNum}
	if o.Read(topic) == nil{
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}