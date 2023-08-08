package questions

import "fmt"

// 編寫一個Go函數，計算兩個整數的和並返回結果。
func Sum(a int, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

// 創建一個Go切片，包含一些整數。編寫一個函數，找到切片中的最大值和最小值，並返回它們。
func FindMaxAndMin(data []int) (int, int) {
	if len(data) == 0 {
		return 0, 0 // 如果切片为空，返回默认值
	}

	min := data[0]
	max := data[0]

	for _, n := range data {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	fmt.Println(min, max)
	return min, max
}

// 創建一個Go結構體表示汽車，包含品牌和型號屬性。實例化幾輛汽車並打印它們的信息。
type Car struct {
	Brand    string
	Category string
}

func (c Car) Introduce() {
	fmt.Printf("This car is made by %s and %s category\n", c.Brand, c.Category)
}

// 編寫一個Go函數，接受一個字符串作為參數，將其逆序並返回。
func ReverseStr(input string) string {
	output := ""
	inputLen := len(input)
	for i, _ := range input {
		output += string(input[inputLen-i-1])
	}

	fmt.Println(output)
	return output
}
