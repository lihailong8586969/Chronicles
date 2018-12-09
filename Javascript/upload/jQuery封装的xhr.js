
// 原链接：https://segmentfault.com/a/1190000008791342

// jQuery封装了xhr的实现, 也可以使用jQuery的ajax获得上传进度，示例代码：

var formData = new FormData(); 
formData.append("file", document.getElementById('file').files[0]); 
formData.append("token", token_value); 

$.ajax({ 
    url: "/uploadurl", 
    type: "POST", 
    data: formData, 
    processData: false, // 不要对data参数进行序列化处理，默认为true
    contentType: false, // 不要设置Content-Type请求头，因为文件数据是以 multipart/form-data 来编码
    xhr: function(){
        myXhr = $.ajaxSettings.xhr();
        if(myXhr.upload){
          myXhr.upload.addEventListener('progress',function(e) {
            if (e.lengthComputable) {
              var percent = Math.floor(e.loaded/e.total*100);
              if(percent <= 100) {
                $("#J_progress_bar").progress('set progress', percent);
                $("#J_progress_label").html('已上传：'+percent+'%');
              }
              if(percent >= 100) {
                $("#J_progress_label").html('文件上传完毕，请等待...');
                $("#J_progress_label").addClass('success');
              }
            }
          }, false);
        }
        return myXhr;
    },
    success: function(res){ 
        // 请求成功
    },
    error: function(res) {
        // 请求失败
        console.log(res);
    }
}); 
 