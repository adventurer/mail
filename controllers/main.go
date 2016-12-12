package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Layout = "layouts/common.html"
	this.TplName = "blank/index.html"
}
