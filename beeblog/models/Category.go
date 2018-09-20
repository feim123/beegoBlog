package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"log"
)

type Category struct {
	Id 				int64
	Title			string
	CreatedTime		time.Time	`orm:"index"`
	Views			int64		`orm:"index"`
	TopicTime		time.Time	`orm:"index"`
	TopicCount		int64
	TopicLastUserId	int64
}

func AddCategory(name string) error {

	o := orm.NewOrm()

	timeNow := time.Now()

	cate := &Category{Title:name, CreatedTime:timeNow, TopicTime:timeNow}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil{
		return err
	}
	_, err = o.Insert(cate)
	if err != nil{
		log.Println("insert error")
		return err
	}
	log.Println("insert success")
	return nil
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id:cid}
	_, err = o.Delete(cate)
	return err
}

func GetCategory(cateName string) (*Category, error) {
	Cate := new(Category)
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	err := qs.Filter("Title", cateName).One(Cate)
	return Cate, err
}

func DelCategoryNum(cateName string) error {
	o := orm.NewOrm()
	Category, _ := GetCategory(cateName)
	Category.TopicCount--
	 _, err := o.Update(Category)
	 return err
}


func AddCategoryNum(cateName string) error {
	o := orm.NewOrm()
	Category, _ := GetCategory(cateName)
	Category.TopicCount++
	_, err := o.Update(Category)
	return err
}