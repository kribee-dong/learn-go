package main

import (
	"fmt"
	"time"
)

//write Data
func erectData(intChan chan int) {
	for i := 1; i <= 10; i++ {
		//放入数据
		intChan <- i //
		fmt.Println("writeData ", i)
		time.Sleep(time.Second)
	}
	close(intChan) //关闭
}

//read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)

}

func main() {

	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go erectData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second * 10)
	for i := 0; ; i++ {
		value, ok := <-exitChan
		fmt.Printf("第%v次loop, value=%v, ok=%v\n", i, value, ok)
		if !ok {
			break
		}
	}

}
