package routers

import (
	"BeegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//注册路由
    beego.Router("/", &controllers.MainController{})
	//增删改查
	beego.Router("/test_orm_all_list", &controllers.TestOrmAllListController{})
	beego.Router("/test_orm_active_list", &controllers.TestOrmActiveListController{})
	beego.Router("/test_orm_completed_list", &controllers.TestOrmCompletedListController{})
	beego.Router("/test_orm_clear_list", &controllers.TestOrmClearListController{})

	beego.Router("/test_orm_add", &controllers.TestOrmAddController{})
	beego.Router("/test_orm_delete", &controllers.TestOrmDeleteController{})
	beego.Router("/test_orm_update", &controllers.TestOrmCompleteController{})
}
