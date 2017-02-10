<!DOCTYPE html>

<html>
<head>
  <title>电商</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }
    #keyword_area{
      width: 300px;
      height: 400px;
      box-sizing: border-box;
    }
    #query_box{
      width: 650px;
      height: 1000px;
      border: 1px solid #eee;
    }
    #id_area{
      width: 300px;
      height: 400px;
      box-sizing: border-box;
    }
  </style>
  <script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
</head>

<body>

  <div id="query_box">
  <h1>查排名</h1>
  
    <textarea id="keyword_area" placeholder="输入keyword"></textarea>
    <textarea id="id_area" placeholder="输入ID"></textarea>
    <br/>
    <br/>
<select id="platform">
  {{range .PlatForm}}
  <option value="{{.}}">{{.}}</option>
  {{end}}
</select>
      <input id="max_page_input" type="number" placeholder="查询最大页数">

    <br/>
    <br/>
    <button id="submit">提交</button>
      <textarea id="result_area">

      </textarea>
  </div>
</body>
<script>
    //获取要用的元素
    var query_box = document.querySelector("#keyword_area")
    var platform_select = document.querySelector("#platform")
    var id_box = document.querySelector("#id_area")
    var submit = document.querySelector("#submit")
    var max_page_input = document.querySelector("#max_page_input")
    var token = ""
    var query_keywords = ""
    var query_product_ids = ""
    var result_area = document.querySelector("#result_area")
    //提交按钮点击事件的处理
    submit.onclick = function () {
        var query_arry = query_box.value.split(/\r\n|\r|\n/g)
       if(query_arry.length >= 1 && query_arry[0]!=""){
        query_keywords  = query_box.value
        query_product_ids = id_box.value
      }else{
          alert("请填写你要查询的排名")
          return
      }

      var platform = platform_select.value
      switch (platform){
          case "tmall":
            query_tmall_rank()
            break;
      }

//      $.ajax({
//              url: "/query/",
//              type:'POST',
//              data:{queryinfo:query_keywords, platform:platform_select.value},
//              success: function(data,textStatus,jqXHR){
//                  console.log(data)
//              },
//              error:function (xhr,textStatus) {
//                  console.log(xhr)
//                  console.log(textStatus)
//              }
//          }
//      )
    }

    function query_tmall_rank() {
        var query_array = query_box.value.split(/\r\n|\r|\n/g)
        var max_page = max_page_input.value == 0 ? 3 : max_page_input.value
        max_page = max_page>10 ? 10 : max_page
        for (var i=0; i< query_array.length; i++){
            for (var page = 0; page< max_page; page++){
                var url = combine_tmall_pc_url(page, encodeURI(query_array[i]))
                var html = transmitRequest(url)
                var htmlElement = $(html)
                var c_page = htmlElement.find("b.ui-page-s-len")
                if(c_page.length == 0){
                    tmall_match(htmlElement, keyword, page)
                }else{
                    var current = c_page.text().split("/")[0]
                    var total_page = c_page.text().split("/")[1]
                    var match_flag = tmall_match(htmlElement,query_array[i],page)
                    if(current > max_page || current == total_page){
                        break
                    }
                    if (match_flag){
                        break
                    }
                }
            }
        }
    }
    
    
    
    function tmall_match(htmlElement, keyword, current_page) {
        var match_flag = false
        ids = id_box.value.split(/\r\n|\r|\n/g)
        if(ids.length>=1&& ids[0]!=""){
            for (var index = 0;index<ids.length; index++ ){
                var data = htmlElement.find("div[data-id='"+ids[index]+"']")
                if(data.length>0){
                    var data = extract_tmall_item(data[0])
                    if(data)
                        result_area.value += keyword+","+data.id+","+data.title+","+data.shop_name+","+(parseInt(data["rank"])+current_page*60)+"\r\n"
                    match_flag = true
                }
            }
        }else{
            var item_list = htmlElement.find("#J_ItemList > div")
            for(var i=0; i<item_list.length; i++){
                var il = item_list[i]
                var data = extract_tmall_item(il)
                if(data)
                    result_area.value += keyword+","+data.id+","+data.title+","+data.shop_name+","+(parseInt(data["rank"])+current_page*60)+"\r\n"

            }
            match_flag = true
        }
        return match_flag
    }
    
    function extract_tmall_item(tmall_item) {
        var data = {}
        if(!tmall_item.querySelector(".productTitle>a")){
            return null
        }
        data["id"] = tmall_item.getAttribute("data-id")
        data["title"] = tmall_item.querySelector(".productTitle>a").getAttribute("title")
        data["shop_name"] = tmall_item.querySelector(".productShop >a").text.trim()
        data["rank"] = tmall_item.querySelector(".productTitle>a").getAttribute("data-p").split("-")[0]
        return data
    }
    
    //服务器端转发
    function transmitRequest(url) {
        var html = ""
        $.ajax({
            url:"/httptransmit/",
            type:'POST',
            async:false,
            data:{url:url,encoding:"GBK"},
            success: function (data,textStatus,jqXHR) {
                html = data
            },
            error:function (xhr,textStatus) {
                console.log(xhr)
                console.log(textStatus)
            }
        })
        return html
    }
    
    function htmlencode(html) {
        return $(html)
    }

    function combine_tmall_pc_url(page,keyword) {
        var url = "https://list.tmall.com/search_product.htm?s="+(page*60)+"&q="+keyword+"&sort=s&style=g&smAreaId=320500&type=pc#J_Filter"
        return url
    }

    //获取result
    function getResult() {
        $.ajax({
            url: "/getresult/",
            type: 'POST',
            data:{token:token, ids:query_product_ids},
            async:false,
            success: function (data,textStatus,jqXHR) {
                console.log(data)
            },
            error:function (xhr,textStatus) {
                console.log(xhr)
                console.log(textStatus)
            }

        })
    }
    
</script>
</html>
