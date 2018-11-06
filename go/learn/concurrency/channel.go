// v, ok := <-ch
// 它可以用来检查Channel是否已经被关闭了。
// 从一个被close的channel中接收数据不会被阻塞，而是立即返回，接收完已发送的数据后会返回元素类型的零值(zero value)。
// 如前所述，你可以使用一个额外的返回参数来检查channel是否关闭。
// 如果OK 是false，表明接收的x是产生的零值，这个channel被关闭了或者为空。

x, ok := <-ch
x, ok = <-ch
var x, ok = <-ch

