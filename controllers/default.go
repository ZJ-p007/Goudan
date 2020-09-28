package controllers

import (
	"HelloBeego0604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	//_"github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//name1 :=c.GetString("name")
	//age1,err := c.GetInt("age")

	//获取get类型的请求参数
	name :=c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	//以admin,18为正确数据进行验证
	if name != "admcin" || age !="18"{
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
	/**c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "zj636500@qq.com"
	c.TplName = "index.tpl"

	 */
}


//该post方法是处理post类型的请求的时候要调用的方法
func (c *MainController) Post() {
	/**fmt.Println("post类型请求")
	user :=c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为:",user)
	psd :=c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是:",psd)

	//与固定值比较 用户名为admin 密码是123456
	/*if user != "admin" || psd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起数据不正确"))
		return
	}
    c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据正确"))

	 */

	//body := c.Ctx.Request.Body




	//json包解析
	var person models.Person
	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
    if err != nil{
    	c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名",person.User,"年龄",person.Age)
	c.Ctx.WriteString("用户名是："+person.User)

}
//连接数据库：database,err :=sql.Open
//_ "github.com/go-sql-driver/mysql"