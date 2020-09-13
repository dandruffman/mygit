package routers

import (
	"BeegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//增删改查
	beego.Router("/test_orm_list", &controllers.TestOrmListController{})
	beego.Router("/test_orm_add", &controllers.TestOrmAddController{})
}
