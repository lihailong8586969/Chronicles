// 错误使用
var obj = { 
    
    // 可以以这种方式定义
    introduce(){  
		
		console.log("你好");
    } 

	hello(){
		
		console.log("Hello World");
	}

}


// 正确使用：三大框架等经常见到这种使用方法
var obj = { 
    
    // 可以以这种方式定义
    introduce(){  
		
		console.log("你好");
    }, 

	hello(){
		
		console.log("Hello World");
	}

}

// 正确使用：
var obj = { 
    
    // 可以以这种方式定义
    introduce(){  
		
		console.log("你好");
    }, 

	hello: function(){
		
		console.log("Hello World");
	}

}