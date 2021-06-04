package main

import (
	"fmt"
	"time"
)

/**
	问题：求1-20000中哪些是素数
*/
func main() {

	sourceChan := make(chan int64, 1000)
	resultChan := make(chan int64, 1000)
	exitChan := make(chan bool, 8)
	n := int64(2000)
	go buildRawData(n, sourceChan)

	for i := 0; i < 8; i++ {
		go processData(sourceChan, resultChan, exitChan)
	}
	go func() {
		unix := time.Now()
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
		since := time.Since(unix)
		fmt.Printf("time cost: %v", since)
	}()
	fmt.Printf("求[1-%v] 中素数结果如下\n", n)
	for {
		val, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Printf("素数：%v\n", val)
	}
}

//将原始数据写入sourceChan中
func buildRawData(n int64, sourceChan chan int64) {
	for i := int64(1); i <= n; i++ {
		sourceChan <- i
	}
	close(sourceChan)
}

//判断素数，并存入primeChan中
func processData(sourceChan chan int64, resultChan chan int64, exitChan chan bool) {
	for {
		val, ok := <-sourceChan
		if !ok {
			break
		}
		flag := true
		for i := int64(2); i < val; i++ {
			if val%i == 0 {
				flag = false
				break
			}
		}
		if flag && val != int64(1) {
			//向resultChan放入值
			resultChan <- val
		}

	}
	//任务完成，向exitChan管道中放入完成标志
	exitChan <- true
}
