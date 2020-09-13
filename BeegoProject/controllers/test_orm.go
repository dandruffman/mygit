package controllers

import (
	"BeegoProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestOrmListController struct {
	beego.Controller
}

type TestOrmAddController struct {
	beego.Controller
}

func (t *TestOrmListController) Get()  {
	o :=orm.NewOrm()
	var users []models.Users
	_,err:=o.QueryTable("Users").All(&users)
	if err!=nil{
		beego.Info("查询失败", err)
		return
	}
	//users :=models.Users{Id:1}
	//o.Read(&users)

	//users :=models.ReadAllUser()

	t.Data["users"] = users
	t.TplName = "list.html"
}

func (t *TestOrmAddController) Get()  {
	t.TplName = "add.html"
}

func (t *TestOrmAddController) Post()  {
	username :=t.GetString("username")
	fmt.Print(username)

	o :=orm.NewOrm()
	users :=models.Users{Username:username}
	o.Insert(&users)
	t.TplName = "list.html"
}