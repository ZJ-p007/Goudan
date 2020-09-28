package routers

import (
	"HelloBeego0604/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Router:路由
    beego.Router("/", &controllers.MainController{})

    //用户注册功能接口
    beego.Router("/register",&controllers.RegisterController{})




}
