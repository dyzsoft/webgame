package routers

import (
	"github.com/astaxie/beego"
	"webgame/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
