package questions

import (
	"errors"
	"fmt"
)

// 編寫一個Go函數，接受一個整數n作為參數，打印出從1到n的所有偶數。
func FindEven(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// 創建一個Go結構體表示矩形，包含長和寬屬性。實現一個方法來計算並返回矩形的面積。
type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Length <= 0 || r.Width <= 0 {
		return 0, errors.New("length and width must be non-negative")
	}
	return r.Length * r.Width, nil
}

// 編寫一個Go函數，接受一個字串切片作為參數，返回包含所有字串的總長度。
func SumStrLen(strSlice []string) int {
	totalLen := 0
	for _, s := range strSlice {
		totalLen += len(s)
	}
	return totalLen
}
