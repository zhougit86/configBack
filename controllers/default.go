package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

type CfgController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *CfgController) Get() {
	o:= orm.NewOrm()
	Objs := make([]models.Object,0)

	qs := o.QueryTable("Object")
	qs.All(&Objs)

	for _,v:=range Objs{
		fmt.Println(v)
	}

	c.Data["json"] = &Objs
	c.ServeJSON()

	//User1:=&models.Object{"Test","IP",
	//"Peizhi",0,10,"get;set",
	//	"get;set","Yes","IP Config",
	//	"C IP Config",
	//}
	//fmt.Println(o.Insert(User1))

	c.Ctx.WriteString("hello")
}

func (c *CfgController) Post() {
	Obj:=models.Object{}
	RB:= c.Ctx.Input.RequestBody
	json.Unmarshal(RB,&Obj)

	o:= orm.NewOrm()

	_,err:=o.Insert(&Obj)
	if err!=nil{
		c.Abort("500")
	}

	c.Ctx.WriteString("received")
}
