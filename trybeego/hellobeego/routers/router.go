package routers

import (
	"playground/trybeego/hellobeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:id(\\d+)", &controllers.UserController{}, "get:GetNum")
	beego.Router("/user/?:id", &controllers.UserController{}, "get:GetInfo")
	beego.Router("/file/*.*", &controllers.UserController{}, "get:GetFile")
}
