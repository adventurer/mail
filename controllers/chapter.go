package controllers

import (
	"fmt"
	"log"
	"mail/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ChapterController struct {
	beego.Controller
}

//列表页面
func (this *ChapterController) Index() {
	o := orm.NewOrm()
	var chaps []models.Chapter
	o.Raw("SELECT * from chapter").QueryRows(&chaps)

	this.Data["List"] = chaps
	this.Layout = "layouts/common.html"
	this.TplName = "mail/chapter.html"
}

//预览页
func (this *ChapterController) Review() {
	_id := this.Input().Get("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		panic("id not found..")
	}
	o := orm.NewOrm()
	chap := models.Chapter{Id: id}
	err = o.Read(&chap)
	this.Ctx.WriteString(chap.Content)
}

//新增邮件页
func (this *ChapterController) New() {
	this.Layout = "layouts/common.html"
	this.TplName = "mail/new.html"
}

//新增邮件提交
func (this *ChapterController) Add() {
	this.Ctx.Output.Body([]byte("add chapter"))
	o := orm.NewOrm()
	chapter := new(models.Chapter)
	chapter.Title = this.GetString("title", "none")
	chapter.Content = this.GetString("content", "none")
	_, err := o.Insert(chapter)
	if err != nil {
		log.Println("chapter inert error..")
	}
	this.Redirect("/mail/chapter", 302)
}

//添加收件人页
func (this *ChapterController) Newuser() {
	pid := this.GetString("pid", "none")
	this.Data["Pid"] = pid
	this.Layout = "layouts/common.html"
	this.TplName = "mail/newuser.html"
}

// 添加收件人提交
func (this *ChapterController) Adduser() {
	o := orm.NewOrm()
	list := new(models.List)
	var lists []models.List

	users := this.GetString("person", "none")
	pid := this.Input().Get("pid")
	ipid, _ := strconv.Atoi(pid)

	usersArr := strings.Split(users, "\r\n")
	for _, u := range usersArr {
		list.Email = u
		list.Pid = ipid
		lists = append(lists, *list)
	}
	successNums, err := o.InsertMulti(len(lists), lists)
	if err != nil {
		fmt.Println(successNums)
	}

	this.Redirect("/mail/chapter", 302)
}

// 修改页面
func (this *ChapterController) Edit() {
	_id := this.Input().Get("id")
	id, _ := strconv.Atoi(_id)
	o := orm.NewOrm()
	chap := new(models.Chapter)
	chap.Id = id
	o.Read(chap)
	this.Data["mail"] = chap
	this.Layout = "layouts/common.html"
	this.TplName = "mail/edit.html"

}

// 修改提交
func (this *ChapterController) Editsub() {
	id, _ := strconv.Atoi(this.Input().Get("id"))
	content := this.GetString("content", "none")
	title := this.GetString("title", "none")
	chap := new(models.Chapter)
	chap.Id = id
	chap.Content = content
	chap.Title = title
	o := orm.NewOrm()
	o.Update(chap, "Title", "Content")
	this.Redirect("/mail/chapter", 302)
}

// 收件用户列表
func (this *ChapterController) User() {
	id := this.GetString("id", "none")
	o := orm.NewOrm()
	var lists []models.List
	o.Raw("SELECT id,email,status from list where pid = " + id).QueryRows(&lists)
	this.Data["List"] = lists
	this.Layout = "layouts/common.html"
	this.TplName = "mail/userlist.html"

}
