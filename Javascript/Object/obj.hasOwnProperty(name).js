// 检查对象是否有某个属性 
// https://www.cnblogs.com/kongxianghai/archive/2013/04/12/3015803.html

obj.hasOwnProperty("name")

 // 没有传递 explain
if( !('explain' in query) || query.explain.length<1 ){

	// 自己调用第三方API获取单词的释义
	explain = "";
}
// 1.使用in关键字。

// 该方法可以判断对象的自有属性和继承来的属性是否存在。

// 2.使用对象的hasOwnProperty()方法。

// 该方法只能判断自有属性是否存在，对于继承属性会返回false。