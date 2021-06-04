package main

import "fmt"
/*
	问题：
	1. 启动一个协程，将1-2000写入到一个channel中，如numChan
	2. 启动8个协程，从numChan中读取数，比如n，并计算1+...+n的值，并存入到resChan
	3. 最后8个协程完成工作后，再遍历resChan，显示结果（如res[1]=1, ... res[10]=55....）
*/
func main() {

	numChan := make(chan int64, 100)
	resChan := make(chan map[int64]int64, 100)
	flagChan := make(chan bool, 8)
	go writeData(numChan, 2000000)
	for i := 0; i < 8; i++ {
		go calculate(numChan, resChan, flagChan)
	}
	go func() {
		for i := 0; i < 8; i++ {
			<- flagChan
		}
		close(resChan)
	}()
	for {
		v, ok := <- resChan
		if !ok {
			break
		}
		for key, val := range v {
			fmt.Printf("读取数据res[%v]=%v\n", key, val)
		}

	}
	/*for {
		if len(flagChan) == 8 {
			close(resChan)
			break
		}
	}
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("main读到数据: [%v]\n", v)
	}*/

}

func writeData(numChan chan int64, num int64) {
	for i := int64(1); i < num; i++ {
		fmt.Println("写入数据：", i)
		numChan <- i
	}
	close(numChan)
}

/*func calculate(source chan int64, resChan chan map[int64]int64, flagChan chan bool) {
	for v := range source {
		fmt.Println("读取到numChan原始数据：", v)
		total := processCal(v)
		mp := make(map[int64]int64, 1)
		mp[v] = total
		resChan <- mp
	}
	flagChan <- true
}*/
func calculate(source chan int64, resChan chan map[int64]int64, flagChan chan bool) {
	for {
		v, ok := <-source
		if !ok {
			break
		}
		fmt.Println("读取到numChan原始数据：", v)
		//计算
		total := processCal(v)
		mp := make(map[int64]int64, 1)
		mp[v] = total
		resChan <- mp
	}
	flagChan <- true
}

func processCal(n int64) int64 {
	count := int64(0)
	for i := int64(1); i <= n; i++ {
		count += i
	}
	return count
}
