package chanlearn

import "fmt"

func Send(ch chan int) {
	ch <- 666
	fmt.Println("123123")
}

func Fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
