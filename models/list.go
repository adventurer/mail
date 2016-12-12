package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type List struct {
	Id         int
	Pid        int
	Email      string
	Status     int
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
	Updated_at time.Time `orm:"auto_now;type(datetime)"`
}

type Email struct {
	Id      int
	To      string
	Title   string
	Content string
}

func init() {
	orm.RegisterModel(new(List))
}

func (this *List) Tagsend(id, tag int) {
	o := orm.NewOrm()
	list := List{Id: id, Status: tag}

	o.Update(&list, "status")
}

func (this *List) Alive() []Email {
	o := orm.NewOrm()
	var lists []List
	o.Raw("SELECT id,pid,email from list where status = 0").QueryRows(&lists)
	chaps := new(Chapter)
	themes := chaps.Theme()
	var emails []Email
	var email Email
	for _, v := range lists {
		email.Id = v.Id
		email.Title = themes[v.Pid].Title
		email.Content = themes[v.Pid].Content
		email.To = v.Email
		emails = append(emails, email)
	}
	return emails
}
