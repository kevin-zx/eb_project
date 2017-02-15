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
	url := "https://list.tmall.com/search_product.htm?s=0&q=%E9%A3%9E%E9%B8%9F%E5%92%8C%E6%96%B0%E9%85%92%E6%97%97%E8%88%B0%E5%BA%97&sort=s&style=g&smAreaId=320500&type=pc#J_Filter"
	header := make(map[string]string)
	header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36"
	c,e :=httpUtil.GetWebConFromUrlWithHeader(url,header)
	enc := mahonia.NewDecoder("GBK")
	output := enc.ConvertString(c)
	fmt.Println(output)
	fmt.Println(e)

}
