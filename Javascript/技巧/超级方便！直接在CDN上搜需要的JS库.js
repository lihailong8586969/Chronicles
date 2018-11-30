<script src="https://cdn.bootcss.com/blueimp-md5/2.10.0/js/md5.js"></script>
<script src="https://cdn.bootcss.com/js-cookie/latest/js.cookie.js"></script>
<script type="text/javascript">
	
	// 我这个操作是每个2秒后刷新页面(因为有些文件有缓存)
	try{

		var last_refresh_time = Cookies.get("last_refresh_time");

		var time = new Date().getTime();

		if( time - last_refresh_time >= 2000 ){

			Cookies.set("last_refresh_time", time);
			
			var url = (window.location.href.split('?'))[0]
			window.location.href = url + "?refresh=" + md5(Math.random());
		}

	}catch(e){

		console.log("URL changing with a error !");
	}

</script>