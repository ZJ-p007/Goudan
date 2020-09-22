package routers

import (
	"HelloBeego0604/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Router:路由
    beego.Router("/index", &controllers.MainController{})

}
