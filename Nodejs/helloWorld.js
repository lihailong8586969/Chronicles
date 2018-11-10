var http = require('http');

http.createServer(function (request, response) {

    // 发送 HTTP 头部 
    // HTTP 状态值: 200 : OK
    // 内容类型: text/html;charset=utf-8
    response.writeHead(200, {'Content-Type': 'text/html;charset=utf-8'});


	// 发送响应数据
	var fs = require("fs");
	var data = fs.readFileSync('./txts/introduce.txt');
    response.end(data.toString()+'\n');
}).listen(8888);


// 终端打印如下信息
console.log('Server running at http://127.0.0.1:8888/');