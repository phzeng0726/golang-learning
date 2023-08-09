package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用 Go 的並行（goroutines）和通道（channels）功能，建立一個簡單的多執行緒程式，例如計算並列印一個數字切片中各數字的平方。
// channel 控制通道數量，在 goroutine 內使用了 ch <- true 之後，又立即使用 <-ch 來接收，這樣做的效果和直接遞增 wg.Add(1) 差不多，而不一定需要使用通道。你可以將這兩行代碼刪除，直接使用 wg.Add(1)，效果會更直觀。

func main() {
	ch := make(chan bool, 10)
	defer close(ch)

	var wg sync.WaitGroup
	slice := []int{1, 2, 3, 4, 5}

	for _, s := range slice {
		ch <- true
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ms := rand.Int31n(1000)
			time.Sleep(time.Duration(ms) * time.Millisecond)
			fmt.Println(n*n, ":", ms)
			<-ch
		}(s)
	}
	wg.Wait()
}
