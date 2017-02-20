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
	url := "https://list.tmall.com/search_product.htm?q=asics&type=p"
	header := make(map[string]string)
	header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36"
	header["referer"] = "https://www.tmall.com/"
	//header["cookie"] = "_tb_token_=EFq7PZKxEPIt; cookie2=77e88cf7685737fe7a37abd895b10f4e; t=742d1a7b1a55992fd566b2878354cbdc;_nk_=%5Cu5C1A%5Cu9648%5Cu9633; cookie17=UNX8iIJwYIcl; "
	//header["cookie"] = "cna=l5MxEZ/OwRoCAbR2NJkdj+wg; _tb_token_=Gs8HJ1cXiJ24; cookie2=ad7359db0f652cacf41a007d485effe5; t=4d53b98003cee20df8f03e59de03733f; tt=tmall-main; _med=dw:1920&dh:1080&pw:1920&ph:1080&ist:0; pnm_cku822=098UW5TcyMNYQwiAiwQRHhBfEF8QXtHcklnMWc%3D%7CUm5Ockt%2BRHxGc0t3T3NMdSM%3D%7CU2xMHDJ7G2AHYg8hAS8XLAIiDFAxVztcIlh2IHY%3D%7CVGhXd1llXGlTa1FkXGBYZFtkU25MdEt%2BRn9Kc0xxSnRNcUp3Q207%7CVWldfS0QMAw3AiIeIgIsSixtCWAQfgFsKwAueC4%3D%7CVmhIGCUFOBgkEC0UNAw0DjsbJxMsETENMAU4GCQQLxIyDjMNNmA2%7CV25Tbk5zU2xMcEl1VWtTaUlwJg%3D%3D; res=scroll%3A1904*5578-client%3A1904*970-offset%3A1904*5578-screen%3A1920*1080; cq=ccp%3D1; l=AgoK4OyYn220N6HfZMU-Z70t2vqs-45V; isg=AgkJZNPrtRJyYEnmvrRpbT08GDX1e_2IdHyBaqt-hfAv8ikE86YNWPcgUuE-"
	header["cookie"] = "_med=dw:375&dh:667&pw:750&ph:1334&ist:1; cna=5hYoEYWpLisCATrS6epjtC5D; sm4=320506; hng=; _tb_token_=ee73ed3387775; uc1=cookie15=URm48syIIVrSKA%3D%3D&existShop=false; uc3=nk2=UoW2s8jnnp6vqIdJAQ%3D%3D&id2=UNGTofEhG8JMHw%3D%3D&vt3=F8dARVfRR4oAmuTAl0Y%3D&lg2=VFC%2FuZ9ayeYq2g%3D%3D; uss=ACuzmzIOgjSrL6Sp6S4%2FGVtLou6wQl4J30oUV%2BHCsljCiuKRpDmXpgBtXDA%3D; lgc=141pmyg859823; tracknick=141pmyg859823; cookie2=3c89643915d02b4fa802cf57332052c3; cookie1=UtABctaSzCi3BaoyrCrIRR%2F%2B67MGJiSnSWk%2BfjEWt3o%3D; unb=3158188536; skt=d3d29c820dea88ba; t=0a0b6a0a7f82c2b76218a76b915f9b2c; _l_g_=Ug%3D%3D; _nk_=141pmyg859823; cookie17=UNGTofEhG8JMHw%3D%3D; login=true; tt=login.tmall.com; pnm_cku822=067UW5TcyMNYQwiAiwQRHhBfEF8QXtHcklnMWc%3D%7CUm5Ockt%2BRHxGck1wTXBFeiw%3D%7CU2xMHDJ7G2AHYg8hAS8WIgwsAl4%2FWTVSLFZ4Lng%3D%7CVGhXd1llXGlTa1FlWmdaZ1JtWmdFfUN6Qn9KdEh2THVBdEF8Rmg%2B%7CVWldfS0SMgc7GyccPBInUm9TbVRoTWRQbElsHmEIJnAm%7CVmhIGCUFOBgkEC0TMwszBjsbJxMsETENMAU4GCQQLxIyDjMKN2E3%7CV25Tbk5zU2xMcEl1VWtTaUlwJg%3D%3D; res=scroll%3A1903*5580-client%3A1903*955-offset%3A1903*5580-screen%3A1920*1080; Hm_lvt_73a389ead45490aaf952f750657e830f=1487302542,1487323511,1487553774,1487579201; Hm_lpvt_73a389ead45490aaf952f750657e830f=1487579201; cq=ccp%3D0; l=Au7uPCSeN7u81JOsYVhSS3Qtvk6x67Ll; isg=AuvrtjUBN5jBAmtHZXuqpknWeg9TWO0M_AXqSV1pwyqN_Ape5dVy0m0aIIto"
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
