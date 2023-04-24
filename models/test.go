package models

import (
	"fmt"
)

func GetSquare(in chan int, out chan int) {
	// 开启一个goroutine，把0-24之间的数发送到ch1
	go func() {
		for i := 0; i < 25; i++ {
			in <- i
		}
		close(in)
	}()

	// 从ch1中取出数据，计算平方，放到ch2
	go func() {

		for value := range in {
			out <- value * value
		}
		close(out)
	}()
}

func TestGetSquare() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	GetSquare(ch1, ch2)
	count := 0
	for i := range ch2 {
		fmt.Println(i)
		count++
	}
	fmt.Printf("\ncount is [%d]", count)
}
