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
	jsonData := models.GetResult(token, ids, plateform)
	this.Ctx.WriteString(jsonData)
}