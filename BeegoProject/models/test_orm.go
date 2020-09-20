package models

import "github.com/astaxie/beego/orm"

type Todos struct {
	Id int
	Matter string
	Status bool
}//创建结构体

func init()  {
	orm.RegisterModel(new(Todos))//注册模型
}