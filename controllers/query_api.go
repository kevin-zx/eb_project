package controllers

import (
	"github.com/astaxie/beego"

	"eb_project/models"
)

type QueryApiController struct {
	beego.Controller
}

func (this *QueryApiController) Post()  {
	query_info := this.GetString("queryinfo")
	platform := this.GetString("platform")
	token := models.QueryRank(query_info,platform)
	this.Ctx.WriteString(token)
}