package main

//import "fmt"
import (
	//"strings"
	//"fmt"
	//"regexp"
	//"time"
	//"encoding/base64"
	//"eb_project/mysqlutil"
	//"strconv"
	//"strings"
	//"eb_project/util"
	//"eb_project/models"
	//"fmt"
)
import (
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/axgle/mahonia"
	"fmt"
)

func main()  {
	url := "https://list.tmall.com/search_product.htm?q=asics&type=p&vmarket=&spm=875.7931836%2FB.a2227oh.d100&from=mallfp..pc_1_searchbutton"
	header := make(map[string]string)
	header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36"
	//header["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	header["cookie"] = "_med=dw:375&dh:667&pw:750&ph:1334&ist:1; cna=5hYoEYWpLisCATrS6epjtC5D; _m_user_unitinfo_=unit|unsz; _m_unitapi_v_=1486692519720; _m_h5_tk=f159725ebb37faea0aaef10230bba43f_1486973764989; _m_h5_tk_enc=82bd19ea26bddb3f33ec6746338e7e68; sm4=320506; hng=; _tb_token_=ee73ed3387775; uc3=nk2=CspNHHAvJqNq&id2=VvuD8lShQ6I%3D&vt3=F8dARHtAySb3hVwe0TQ%3D&lg2=UIHiLt3xD8xYTw%3D%3D; uss=AH%2Br4rcQ5edWcXH%2FyoIvD7hmRKEzwHzmgEGw80fUknqQzjU6no4h7R5QvA%3D%3D; lgc=iloveopen; tracknick=iloveopen; cookie2=3c89643915d02b4fa802cf57332052c3; skt=e7ba48f26ba285b5; t=0a0b6a0a7f82c2b76218a76b915f9b2c; pnm_cku822=118UW5TcyMNYQwiAiwQRHhBfEF8QXtHcklnMWc%3D%7CUm5Ockt%2BRHpFfkB4QnZDeS8%3D%7CU2xMHDJ7G2AHYg8hAS8WIgwsAl4%2FWTVSLFZ4Lng%3D%7CVGhXd1llXGlTbVJpV29VYVRuWWRGeUV8RnNIdk5zTHRMeU10TXNdCw%3D%3D%7CVWldfS0RMQU%2BCioWKwslDmAdewpnAn4edVsNWw%3D%3D%7CVmhIGCUFOBgkEC0TMwg1DzMTLxskGTkFOA0wECwYJxo6BjsCP2k%2F%7CV25Tbk5zU2xMcEl1VWtTaUlwJg%3D%3D; res=scroll%3A1903*6071-client%3A1903*955-offset%3A1903*6071-screen%3A1920*1080; Hm_lvt_73a389ead45490aaf952f750657e830f=1487208727,1487239175,1487302542,1487323511; Hm_lpvt_73a389ead45490aaf952f750657e830f=1487326359; cq=ccp%3D1; l=Ak9PlP/1RiRlD5IrePMDyHn4X-lZcaOQ; isg=AnR0o2GokAFG-QRKliatEwqTRTJsuZg3X2AFBA7VBv-EeRTDNl1oxyo7jwZb"
	//header[":path"] = "/search_product.htm?s=0&q=%E7%AB%A5%E8%A3%85&sort=s&style=g&smAreaId=320500&type=pc"
	//header["content-encoding"]="gzip"
	header["content-language"]="zh-CN"
	//header["content-type"] = "text/html;charset=GBK"
	c,e :=httpUtil.GetWebConFromUrlWithHeader(url,header)
	enc := mahonia.NewDecoder("GBK")
	output := enc.ConvertString(c)
	fmt.Println(output)
	fmt.Println(e)

}
