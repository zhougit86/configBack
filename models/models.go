package models

import (
	"github.com/astaxie/beego/orm"
)

type Object struct {
	ServiceName   string
	ObjectName    string `orm:"pk"`
	ObjectType    string
	MinNum        int
	MaxNum        int
	SupportOps    string
	OpsWithScript string
	AdmVisible    string
	EngName       string
	ChnName       string

}

func init() {
	orm.RegisterModel(new(Object))
}
