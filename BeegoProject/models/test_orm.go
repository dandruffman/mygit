package models

import "github.com/astaxie/beego/orm"

type Users struct {
	Id int
	Username string
}

func (a * Users) TableName() string{
	return "users"
}

func init()  {
	orm.RegisterModel(new(Users))
}