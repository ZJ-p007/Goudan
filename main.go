package main

import (
	"HelloBeego0604/dbmysql"
	_ "HelloBeego0604/routers"
	"github.com/astaxie/beego"
	//mysql驱动
	_"github.com/go-sql-driver/mysql"
)

func main() {
    //获取配置选项
   /** appName :=config.String("appname")
	fmt.Println("项目应用名称:",appName)
    port,err := config.Int("httpport")
    if err != nil{
    	//配置信息解析错误
    	panic("项目配置信息解析错误，请查验后重试")
	}
    fmt.Println("应用的监听端口:",port)
    */
	//定义conifg变量，接收并赋值全局配置变量

	//1.连接数据库
	dbmysql.Connect()



    //其他配置

    //启动程序
    beego.Run()
}


