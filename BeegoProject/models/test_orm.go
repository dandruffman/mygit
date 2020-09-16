package models

import "github.com/astaxie/beego/orm"

type Users struct {
	Id int
	Username string
}//创建结构体

func init()  {
	orm.RegisterModel(new(Users))//注册模型
}