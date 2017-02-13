package controllers

import (
	"github.com/astaxie/beego"
)

type QueryRankController struct {
	beego.Controller
}

func (c *QueryRankController) Get() {
	c.Data["PlatForm"] = []string{"tmall","taobao","tmall_m","taobao_m"}
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "query_rank.tpl"
}
