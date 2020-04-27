package routers

import (
	"github.com/beautytiger/go-playground/trybeego/ihome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
