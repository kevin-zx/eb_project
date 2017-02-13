package controllers

import (
	"github.com/astaxie/beego"
)

type QueryRankController struct {
	beego.Controller
}

func (c *QueryRankController) Get() {
	c.Data["PlatForm"] = []string{"tmall","taobao","jd"}
	c.Data["js"] = []string{"rank.js"}
	c.TplName = "query_rank.tpl"
}
