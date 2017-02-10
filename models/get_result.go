package models

import (
	"github.com/kevin-zx/go-util/regexpUtil"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"strings"
	"encoding/json"
)

func GetResult(token string, ids string)string{
	if is_query_done(token){
		return get_query_data(token, ids)
	}else{
		return `{"message":"还没查好"}`
	}

}
//获取查排名的结果
func get_query_data(token string, ids string) string {
	sql := "select * from tmall_view where `batch` = ?"
	id_array := regexpUtil.SplitString2Line(ids)
	add := 1
	if len(id_array)<=1 && id_array[0]==""{
		add = 0
	}
	param := make([]interface{},len(id_array)+add)
	param[0] = token
	question_mark_array := []string{}
	//如果是根据id去查排名则需要装配url
	if len(id_array) > 0 && id_array[0]!=""{
		sql += " and `p_id` in ("
		for i,id := range id_array{
			question_mark_array = append(question_mark_array, "?")
			param[i+1] = id
		}
		sql += strings.Join(question_mark_array, ",")+")"
	}


	data,_ :=  mysqlutil.SelectAll(sql,param...)
	json_data,_:=json.Marshal(data)
	return string(json_data)
}

//查看这个批次的排名是否查询完成
func is_query_done(token string)  (bool){
	flag := false
	//查看任务库中是否含有
	c,_ := mysqlutil.SelectAll("select 1 from spider_keyword where `batch` = ? and `status` = 1 limit 1", token)
	//这里没有结果有可能是token不对
	if len(*c)==0{
		flag = true
	}else
	if (*c)[0]["1"] == "1"{
		flag = false
	}

	return flag
}