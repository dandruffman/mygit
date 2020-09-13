package main

import (
	_ "BeegoProject/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres",orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=admin dbname=test port=5432 sslmode=disable")
}

func main() {
	beego.Run()
}

