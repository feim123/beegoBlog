package models

import (
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id 			int64
	Tid 		int64
	Name		string
	Content		string		`orm:"size(5000)"`
	CreatedTime time.Time	`orm:"index"`
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return err
	}
	reply := &Comment{
		Tid:			tidNum,
		Name:			nickname,
		Content:		content,
		CreatedTime:	time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)

	err = AddTopicReplyCount(tidNum)

	return err
}

func GetAllReply(tid string) ([]*Comment, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return nil, err
	}

	comments := make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&comments)
	return comments, nil
}

func DelReply(cid, tid string) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil{
		return err
	}
	o := orm.NewOrm()
	comment := &Comment{Id:cidNum}
	_, err = o.Delete(comment)

	err = DelTopicReplyCount(tidNum)

	return err
}