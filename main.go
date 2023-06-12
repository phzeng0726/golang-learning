package main

import (
	"learning/helpers"
	"log"
	"sync"
)

const numPool = 1000

func CalculateValue(intChan chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done() // 執行結束後通知等待組完成

	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
	log.Println(i)
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	var wg sync.WaitGroup
	var results []int

	for i := 0; i < 30; i++ {
		wg.Add(1) // 增加等待組中的計數器

		go CalculateValue(intChan, &wg, i)

		num := <-intChan

		results = append(results, num)
	}

	wg.Wait() // 等待所有goroutine完成

	log.Println(results)
}
