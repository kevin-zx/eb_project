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
	url := "https://list.tmall.com/search_product.htm?s=0&q=%E4%B8%AD%E6%96%87&sort=s&style=g&smAreaId=320500&type=pc#J_Filter"
	header := make(map[string]string)
	header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36"
	c,_ :=httpUtil.GetWebConFromUrlWithHeader(url,header)
	enc := mahonia.NewDecoder("GBK")
	output := enc.ConvertString(c)
	fmt.Println(output)

}
