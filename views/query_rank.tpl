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
      #result_area{
          width: 645px;
          height: 300px;
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
      <textarea id="result_area" placeholder="结果"></textarea>
  </div>
</body>

{{range .js}}
<script src="/static/js/{{.}}"></script>
{{end}}
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
        result_area.value = ""
      }else{
          alert("请填写你要查询的排名")
          return
      }

      var platform = platform_select.value
      query(platform)
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

    
</script>
</html>
