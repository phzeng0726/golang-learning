package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	// 獲取當前 Goroutine 的獨立種子
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	randomNumber := rand.Intn(n)
	return randomNumber
}
