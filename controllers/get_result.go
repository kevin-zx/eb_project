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
	jsonData := models.GetResult(token, ids)
	this.Ctx.WriteString(jsonData)
}