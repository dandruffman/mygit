package controllers

import (
	"BeegoProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//创建结构体
type TestOrmListController struct {
	beego.Controller
}

type TestOrmAddController struct {
	beego.Controller
}

type TestOrmDeleteController struct {
	beego.Controller
}

type TestOrmUpdateController struct {
	beego.Controller
}

//通过get（）用列表显示数据信息
func (t *TestOrmListController) Get()  {
	o :=orm.NewOrm()//创建ORM对象
	var users []models.Users//创建数组
	_,err:=o.QueryTable("Users").All(&users)//取出全部users表数据
	if err!=nil{
		beego.Info("查询失败", err)
		return
	}

	//将取出的数据传到list页面
	t.Data["users"] = users
	t.TplName = "list.html"
}

func (t *TestOrmAddController) Get()  {
	t.TplName = "add.html"
}

//通过post方式获取数据插入数据库
func (t *TestOrmAddController) Post()  {
	username :=t.GetString("username")//获取前端传过来的用户名
	fmt.Print(username)

	//将获取的用户名插入数据库（主键id为serial自增integer类型）
	o :=orm.NewOrm()
	users :=models.Users{Username:username}
	o.Insert(&users)
	t.Redirect("/test_orm_list",302)
}

func (t *TestOrmDeleteController) Get()  {
	id, _ :=t.GetInt("id")
	fmt.Print(id)

	o :=orm.NewOrm()
	users :=models.Users{Id:id}
	o.Delete(&users)
	t.Redirect("/test_orm_list",302)
}

//在编辑页面显示要更改的旧用户名
func (t *TestOrmUpdateController) Get()  {
	id, _ :=t.GetInt("id")
	fmt.Print(id)

	o :=orm.NewOrm()
	users :=models.Users{Id:id}
	o.Read(&users)

	t.Data["users"] = users
	t.TplName = "edit.html"
}

func (t *TestOrmUpdateController) Post()  {
	//拿到新用户名
	username :=t.GetString("username")
	fmt.Print(username)
	//拿到需要修改的用户名对应的id
	id, _ :=t.GetInt("id")
	fmt.Print(id)

	o :=orm.NewOrm()
	//找到数据库中id对应的数据
	users :=models.Users{Id:id}
	o.Read(&users)
	//更新用户名
	users.Username=username
	o.Update(&users)
	t.Redirect("/test_orm_list",302)
}

