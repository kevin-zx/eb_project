package routers

import (
	"eb_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/query_rank/", &controllers.QueryRankController{})
	beego.Router("/query/", &controllers.QueryApiController{})
	beego.Router("/getresult/", &controllers.GetResultController{})
	beego.Router("/httptransmit/", &controllers.HttpTransmitController{})
}
