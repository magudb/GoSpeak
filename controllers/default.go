package controllers

import (
	"fmt"

	"../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

type QouteModel struct {
	Quote  string
	Author string
}

type QuoteController struct {
	beego.Controller
}

func (this *QuoteController) Get() {
	m := &models.Qoute{1, "i'm a fin qoute", "me"}

	this.Data["json"] = m
	this.ServeJSON()
}

func (this *QuoteController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	quote := QouteModel{}
	if err := this.ParseForm(&quote); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		beego.Debug("Qoute supplied:", quote.Quote)
		m := &models.Qoute{ID: -1, Text: quote.Quote, Author: quote.Author}
		id, err := o.Insert(&m)
		if err == nil {
			msg := fmt.Sprintf("Qoute inserted with id:", id)
			beego.Debug(msg)
		} else {
			msg := fmt.Sprintf("Couldn't insert new quote. Reason: ", err)
			beego.Debug(msg)
		}
	}
	this.Data["json"] = quote
	this.ServeJSON()
}
