package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Object struct {
	//gorm.Model
	ServiceName   string  `gorm:"column:service_name"`
	ObjectName    string  `gorm:"column:object_name;primary_key"`
	//ObjectType    string
	//MinNum        int
	//MaxNum        int
	//SupportOps    string
	//OpsWithScript string
	//AdmVisible    string
	//EngName       string
	//ChnName       string
	ObjAttr []ObjAttr `gorm:"ForeignKey:Object;AssociationForeignKey:ObjectName"`

}

type ObjAttr struct{
	Object  string    `gorm:"primary_key"`
	AttrName string    `gorm:"primary_key"`
}


type Profile struct {
	gorm.Model
	Name   string
	UserID uint
}

type User struct {
	gorm.Model
	Refer   uint
	Profiles []Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
}

func main() {
	db, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/config?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Object{},&ObjAttr{},&Profile{},&User{})


	obj:=&Object{ObjectName:"ip"}
	objattr:=&ObjAttr{}



	db.Create(obj)
	db.Create(objattr)



	db.Model(&obj).Association("ObjAttr").Append(ObjAttr{"ip","desc2"})


	db.Model(&obj).Association("ObjAttr").Find(&objattr)

	fmt.Println(objattr)
	fmt.Println(db.Model(&obj).Association("ObjAttr").Count())

}

