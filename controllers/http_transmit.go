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
	//webcon,err := httpUtil.GetWebConFromUrl(url)
	header := make(map[string]string)
	header["cookie"] = "_med=dw:375&dh:667&pw:750&ph:1334&ist:1; cna=5hYoEYWpLisCATrS6epjtC5D; _m_user_unitinfo_=unit|unsz; _m_unitapi_v_=1486692519720; _m_h5_tk=f159725ebb37faea0aaef10230bba43f_1486973764989; _m_h5_tk_enc=82bd19ea26bddb3f33ec6746338e7e68; sm4=320506; hng=; _tb_token_=ee73ed3387775; uc3=nk2=CspNHHAvJqNq&id2=VvuD8lShQ6I%3D&vt3=F8dARHtAySb3hVwe0TQ%3D&lg2=UIHiLt3xD8xYTw%3D%3D; uss=AH%2Br4rcQ5edWcXH%2FyoIvD7hmRKEzwHzmgEGw80fUknqQzjU6no4h7R5QvA%3D%3D; lgc=iloveopen; tracknick=iloveopen; cookie2=3c89643915d02b4fa802cf57332052c3; skt=e7ba48f26ba285b5; t=0a0b6a0a7f82c2b76218a76b915f9b2c; tt=tmall-main; pnm_cku822=186UW5TcyMNYQwiAiwZTXFIdUh1SHJOe0BuOG4%3D%7CUm5Ockt%2BRHpFfEB1THJHfyk%3D%7CU2xMHDJ7G2AHYg8hAS8WIgwsAl4%2FWTVSLFZ4Lng%3D%7CVGhXd1llXGlTbVJrV2JbZVBoX2JAdE5zTnpEfEh8RHFFeExwTGI0%7CVWldfS0TMw8wCDQUKwslTDcZTxk%3D%7CVmhIGCUFOBgkEC0TMwg0DDERLRkmGzsHOg8yEi4aJRg4BDkAPWs9%7CV25Tbk5zU2xMcEl1VWtTaUlwJg%3D%3D; res=scroll%3A1903*6150-client%3A1903*591-offset%3A1903*6150-screen%3A1920*1080; Hm_lvt_73a389ead45490aaf952f750657e830f=1487208727,1487239175,1487302542,1487323511; Hm_lpvt_73a389ead45490aaf952f750657e830f=1487324186; cq=ccp%3D1; l=AqGhl2rVUBLnhSyVWlnVfrM/MWe6YhUJ; isg=ApGRxuuVrQbbB8EVyzXgdN8woJ0R-hs8KmOA-3Mnodg-GrBsvU2PQUk8ylkG"
	webcon,err := httpUtil.GetWebConFromUrlWithHeader(url,header)
	encoding := this.GetString("encoding")

	enc := mahonia.NewDecoder(encoding)
	if err != nil {
		println(err.Error())
		this.Ctx.WriteString(`{"message":"request_erro","status":"0"}`)
	}else {
		this.Ctx.WriteString(enc.ConvertString(webcon))
	}
}