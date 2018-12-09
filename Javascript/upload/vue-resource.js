

// 原链接：https://segmentfault.com/a/1190000008791342

var formData = new FormData();
formData.append('token', token_value);  // csrf token
formData.append("works", document.getElementById('file').files[0]); // file
var url = $("#R_batch_upload_url").val();

vm.$http.post(url, formData, {
  progress: (e) => {
    if (e.lengthComputable) {
      var percent = Math.floor(e.loaded/e.total*100);
      if(percent <= 100) {
        $("#J_progress_bar").progress('set progress', percent);
        $("#J_progress_label").html('已上传：'+percent+'%');
      }
      if(percent >= 100) {
        $("#J_progress_label").html('文件上传完毕，提交表单中，请等待...');
        $("#J_progress_label").addClass('success');
      }
    }
  }
})
.then((res) => {
  if(res.ok && res.status === 200) {
    window.location.href = window.location.href;
  }
}, (res) => {
  if(res.status === 400) {
      $("#J_progress_label").html('文件格式错误，请修改后重试');
      $("#J_progress_label").addClass('warning');
      console.log(res);
      vm.errMsg.show = true;
      vm.errMsg.msg = res.body.msg;
      vm.canSend = true;
      // TODO hide the loader dimmer
      $("#J_upload_batch").dimmer("hide");
    } else {
      $("#J_progress_label").html(res.statusText);
      $("#J_progress_label").addClass('warning');
    }
});