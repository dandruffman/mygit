package controllers

import (
	"BeegoProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//创建结构体
type TestOrmAllListController struct {
	beego.Controller
}

type TestOrmActiveListController struct {
	beego.Controller
}

type TestOrmCompletedListController struct {
	beego.Controller
}

type TestOrmClearListController struct {
	beego.Controller
}

type TestOrmAddController struct {
	beego.Controller
}

type TestOrmDeleteController struct {
	beego.Controller
}

type TestOrmCompleteController struct {
	beego.Controller
}

//通过get（）用列表显示数据信息
func (t *TestOrmAllListController) Get()  {
	o :=orm.NewOrm()//创建ORM对象
	var todos []models.Todos//创建数组
	_,err:=o.QueryTable("todos").All(&todos)//取出全部users表数据
	if err!=nil{
		beego.Info("查询失败", err)
		return
	}

	//将取出的数据传到list页面
	t.Data["todos"] = todos
	t.TplName = "list.html"
}

func (t *TestOrmActiveListController) Get()  {
	o :=orm.NewOrm()
	var todos []models.Todos
	_,err:=o.QueryTable("todos").Filter("status",false).All(&todos)//取出全部users表数据
	if err!=nil{
		beego.Info("查询失败", err)
		return
	}
	//将取出的数据传到list页面
	t.Data["todos"] = todos
	t.TplName = "list.html"
}

func (t *TestOrmCompletedListController) Get()  {
	o :=orm.NewOrm()
	var todos []models.Todos
	_,err:=o.QueryTable("todos").Filter("status",true).All(&todos)
	if err!=nil{
		beego.Info("查询失败", err)
		return
	}
	//将取出的数据传到list页面
	t.Data["todos"] = todos
	t.TplName = "list.html"
}

func (t *TestOrmClearListController) Get()  {
	o :=orm.NewOrm()
	_,err:=o.QueryTable("todos").Filter("status",true).Delete()//取出全部users表数据
	if err!=nil{
		beego.Info("查询失败", err)
		return
	}
	t.Redirect("/test_orm_all_list",302)
}

//通过post方式获取数据插入数据库
func (t *TestOrmAddController) Post()  {
	matter :=t.GetString("matter")//获取前端传过来的用户名
	fmt.Print(matter)

	//将获取的用户名插入数据库（主键id为serial自增integer类型）
	o :=orm.NewOrm()
	todos :=models.Todos{Matter:matter}
	o.Insert(&todos)
	t.Redirect("/test_orm_all_list",302)
}

func (t *TestOrmDeleteController) Get()  {
	id, _ :=t.GetInt("id")
	fmt.Print(id)

	o :=orm.NewOrm()
	todos :=models.Todos{Id:id}
	o.Delete(&todos)
	t.Redirect("/test_orm_all_list",302)
}

func (t *TestOrmCompleteController) Post() {
	id := t.GetStrings("matter")
	beego.Info(id)

	o :=orm.NewOrm()
	for _,x :=range id {
		_, err := o.QueryTable("todos").Filter("id", x).Update(orm.Params{
			"status": "true",
		})
		if err != nil {
			t.Ctx.WriteString("查询出错!")
			return
		}
	}
	t.Redirect("/test_orm_all_list",302)
}