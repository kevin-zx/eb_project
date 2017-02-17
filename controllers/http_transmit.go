package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/axgle/mahonia"
)

type HttpTransmitController struct {
	beego.Controller
}

func (this *HttpTransmitController)  Post(){
	url := this.GetString("url")
	webcon,err := httpUtil.GetWebConFromUrl(url)
	encoding := this.GetString("encoding")

	enc := mahonia.NewDecoder(encoding)
	if err != nil {
		println(err.Error())
		this.Ctx.WriteString(`{"message":"request_erro","status":"0"}`)
	}else {
		this.Ctx.WriteString(enc.ConvertString(webcon))
	}
}