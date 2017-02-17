package main

import (
	_ "eb_project/routers"
	"github.com/astaxie/beego"
	"github.com/kevin-zx/go-util/mysqlUtil"
)

func main() {
	mysqlutil.GlobalMysqlUtil.InitMySqlUtil("115.159.3.51",3306,"remote","Iknowthat","eb_spider")
	beego.Run()
}
