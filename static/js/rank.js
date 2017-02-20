//common
function htmlencode(html) {
    return $(html)
}

function query(platform) {
    var query_array = query_box.value.split(/\r\n|\r|\n/g)
    var max_page = max_page_input.value <= 0 ? 3 : max_page_input.value
    var platform = platform_select.value
    max_page = max_page>10 ? 10 : max_page
    console.log(Date.parse(new Date()))
    for (var i=0; i< query_array.length; i++){
        for (var page = 0; page< max_page; page++){
            console.log(query_array[i]+"i:"+i+" page:"+page)
            console.log(Date.parse(new Date()))
            var done = false
            switch  (platform){
                case "tmall":
                    done = query_tmall_rank(query_array[i],page,max_page)
                    break
                case "taobao":
                    break
                default:
                    alert("平台暂未开发查询")
                    return
            }
            if (done){
                break
            }
        }
    }

}

//taobao
function query_taobao_rank(keyword, page) {
    var is_done = false
    // var url = combine_taobao_pc_url()
}
function combine() {
    
}
//tmall
function query_tmall_rank(keyword, page,max_page) {
    var is_done = false
    var url = combine_tmall_pc_url(page, encodeURI(keyword))
    var html = transmitRequest(url)
    var trytimes = 10
    while (html.length<1000){
        trytimes = trytimes-1
        if(trytimes==0){
            console.log("erro:"+url)
            return true
        }
        html = transmitRequest(url)
    }
    var htmlElement = $(html)
    var c_page = htmlElement.find("b.ui-page-s-len")
    if(c_page.length == 0){
        tmall_match(htmlElement, keyword, page)
        is_done = true
    }else{
        var current = c_page.text().split("/")[0]
        var total_page = c_page.text().split("/")[1]
        var match_flag = tmall_match(htmlElement,keyword,page)
        if(parseInt(current)> parseInt(max_page)|| parseInt(current) == parseInt(total_page)){
            is_done = true
        }
        if (match_flag){
            is_done = true
        }
    }

    return is_done
}

function tmall_match(htmlElement, keyword, current_page) {
    var match_flag = false
    ids = id_box.value.split(/\r\n|\r|\n/g)
    if(ids.length>=1&& ids[0]!=""){
        for (var index = 0;index<ids.length; index++ ){
            var data = htmlElement.find("div[data-id='"+ids[index]+"']")
            if(data.length>0){
                var data = extract_tmall_item(data[0])
                if(data){
                    result_area.value += keyword+","+data.id+","+data.title+","+data.shop_name+","+(parseInt(data["rank"])+current_page*60)+"\r\n"
                    continue
                }
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
        match_flag = false
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
    if(tmall_item.querySelector(".productStatus > span>em").text){
        data["shop_name"] = tmall_item.querySelector(".productStatus > span>em").text.trim().replace("笔 ","").replace("")
    }
    data["rank"] = tmall_item.querySelector(".productTitle>a").getAttribute("data-p").split("-")[0]
    return data
}


function combine_tmall_pc_url(page,keyword) {
    var url = "https://list.tmall.com/search_product.htm?s="+(page*60)+"&q="+keyword+"&sort=s&style=g&smAreaId=320500&type=pc#J_Filter"
    return url
}
