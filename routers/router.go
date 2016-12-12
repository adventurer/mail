package routers

import (
	"mail/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mail/chapter", &controllers.ChapterController{}, "*:Index")
	beego.Router("/mail/new", &controllers.ChapterController{}, "get:New")
	beego.Router("/mail/new", &controllers.ChapterController{}, "post:Add")
	beego.Router("/mail/review", &controllers.ChapterController{}, "get:Review")
	beego.Router("/mail/edit", &controllers.ChapterController{}, "get:Edit")
	beego.Router("/mail/editsub", &controllers.ChapterController{}, "post:Editsub")
	beego.Router("/mail/user", &controllers.ChapterController{}, "get:User")
	beego.Router("/chapter/newuser", &controllers.ChapterController{}, "get:Newuser")
	beego.Router("/chapter/adduser", &controllers.ChapterController{}, "post:Adduser")
}
