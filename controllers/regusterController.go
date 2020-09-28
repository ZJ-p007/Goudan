package controllers

import (
	"HelloBeego0604/dbmysql"
	"HelloBeego0604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
   //解析前端提交的用户注册信息
	dataBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
    if err != nil{
    	r.Ctx.WriteString("数据接错误，请重试")
		return
    }
    var user models.User
	err =json.Unmarshal(dataBytes,&user)
	if err != nil{
		r.Ctx.WriteString("数据解析错误，请重试")
		return
	}
    //一切正常，将用户信息保存到数据库中
    //直接调用保存数据的一个函数，并判断保存后的结果
    row, err := dbmysql.AddUser(user)
	if err != nil{
		r.Ctx.WriteString("注册用户信息失败，请重试")
	}
	fmt.Println(row)
	r.Ctx.WriteString("恭喜注册用户信息成功")

}
