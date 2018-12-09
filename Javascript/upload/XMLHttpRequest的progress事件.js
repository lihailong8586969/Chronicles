

// 原链接：https://segmentfault.com/a/1190000008791342


var formData = new FormData(); 
formData.append("file", document.getElementById('file').files[0]); 
formData.append("token", token_value); // 其他参数按这样子加入

var xhr = new XMLHttpRequest();
xhr.open('POST', '/uploadurl');
// 上传完成后的回调函数
xhr.onload = function () {
  if (xhr.status === 200) {
　　console.log('上传成功');
  } else {
  　console.log('上传出错');
  }
};
// 获取上传进度
xhr.upload.onprogress = function (event) {
  if (event.lengthComputable) {
    var percent = Math.floor(event.loaded / event.total * 100) ;
    // 设置进度显示
    $("#J_upload_progress").progress('set progress', percent);
  }
};
xhr.send(formData);
