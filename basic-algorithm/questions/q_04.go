package questions

import "math"

// 編寫一個Go函數，接受一個整數n作為參數，返回n的階乘（n!）。
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	factorial := n
	for i := 1; i < n; i++ {
		factorial *= i
	}
	return factorial
}

// 創建一個Go結構體表示圓柱體，包含半徑和高度屬性。實現一個方法來計算並返回圓柱體的體積。
type Cylinder struct {
	Radius float64
	Height float64
}

func (c Cylinder) Volume() float64 {
	return c.Radius * c.Radius * math.Pi * c.Height
}

// 編寫一個Go函數，接受一個整數n作為參數，返回第n個斐波那契數。
// 斐波那契數列是一個數列，每個數字是前兩個數字之和。例如，數列的前幾個數字為：0, 1, 1, 2, 3, 5, 8, 13, 21, ...
// 所以，如果您的函數被傳遞了n=5，則應該返回斐波那契數列中的第5個數字，即3
// func Fibonacci(n int) int {
// 	if n == 1 {
// 		return 0
// 	}
// 	if n == 2 {
// 		return 1
// 	}
// 	total := Fibonacci(n-1) + Fibonacci(n-2)
// 	return total
// }

// 03優化版，效能至少差600倍
// https://chat.openai.com/share/045d252f-aa1f-4983-bb54-4de7470decef
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev, curr = curr, next
	}

	return curr
}
