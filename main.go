package main

import (
	"./controllers"
	"./models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")

	orm.RegisterModel(new(models.Qoute))
}

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/quotes", &controllers.QuoteController{})
	beego.Run()
}
