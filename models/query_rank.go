package models

import (
	"fmt"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"github.com/kevin-zx/go-util/regexpUtil"
	"encoding/base64"
	"time"
)
//查询排名
func QueryRank(query_info string, platform string) (string) {
	//获取时间戳变成批次串
	timestamp := time.Now().UnixNano()
	t := fmt.Sprintf("%d",timestamp)
	batch := base64.StdEncoding.EncodeToString([]byte(t))
	//task插入库中
	query_data := regexpUtil.SplitString2Line(query_info)
	for _,keyword :=  range query_data{
		if keyword != ""{
			mysqlutil.Exec("insert into spider_keyword (`keyword`,`platform`,`batch`,`create_time`,`status`) VALUES (?,?,?,now(),1)",keyword,platform,batch)
		}
	}
	return batch
}
