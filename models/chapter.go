package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Chapter struct {
	Id         int
	Title      string
	Content    string `orm:"type(text)"`
	Status     int
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
	Updated_at time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Chapter))
}

func (this *Chapter) Theme() map[int]Chapter {
	theme := make(map[int]Chapter, 0)
	o := orm.NewOrm()
	var chaps []Chapter
	o.Raw("SELECT id,title,content from chapter where status = 0").QueryRows(&chaps)

	for _, v := range chaps {
		theme[v.Id] = v
	}
	return theme
}
