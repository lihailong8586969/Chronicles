package main
import "fmt"

func sum(s []int, c chan int) {

	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum // 将和送入 c
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[3:], c)
	go sum(s[:3], c)

	// 接收顺序是按后进先出的方式(栈) ? 为什么输出是反的啊 ?
	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y)


	// 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
	// 修改示例填满缓冲区，然后看看会发生什么。
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
