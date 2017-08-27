package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:1234@tcp(localhost:3306)/config?charset=utf8")
}

func main() {
	beego.Run()
}
