

var arr = [];

// 把 arr 当成了对象，age 和 name 都变成了 arr 对象的属性，而不是下标
arr["age"] = 19;
arr["name"] = "ChenSX";

// 其实就是这个
// arr.age = 19;
// arr.name = "ChenSX";



