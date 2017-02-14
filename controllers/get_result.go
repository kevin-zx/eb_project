package controllers

import "github.com/astaxie/beego"
import (
	"eb_project/models"
)

type GetResultController struct {
	beego.Controller
}

func (this *GetResultController) Post()  {
	token := this.GetString("token")
	ids := this.GetString("ids")
	plateform := this.GetString("plateform")
	data_area := this.GetString("data_area")
	jsonData := models.GetResult(token, ids, plateform,data_area)
	this.Ctx.WriteString(jsonData)
}

func (this *GetResultController) Get()  {
	token := this.GetString("token")
	ids := this.GetString("ids")
	plateform := this.GetString("plateform")
	data_area := this.GetString("data_area")
	//println(data_area)
	//println(data_area)
	jsonData := models.GetResult(token, ids, plateform,data_area)
	this.Ctx.WriteString(jsonData)
}