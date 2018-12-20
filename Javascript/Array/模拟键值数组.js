// https://blog.csdn.net/chenssy/article/details/7366160


var arr =[];

// 你以为在定义键，其实只是在 arr 对象上 定义属性
arr["a"] = 1;
arr["b"] = 2;

console.log(arr.length); // 输出 0

console.log( Array.isArray(arr) ); // 输出 true

var count = 0;
for(key in arr){

	++count;
	conosle.log(key); // 只输出我刚才定义的属性
}

console.log(count); // 输出 2