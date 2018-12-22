// 实例化委托（Delegate）
// 一旦声明了委托类型，委托对象必须使用 new 关键字来创建，且与一个特定的方法有关。
// 当创建委托时，传递到 new 语句的参数就像方法调用一样书写，但是不带有参数。
// 例如：

public delegate void printString(string s);

printString ps1 = new printString(WriteToScreen);

printString ps2 = new printString(WriteToFile);