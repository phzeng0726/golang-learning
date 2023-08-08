package questions

// 編寫一個Go函數，接受一個整數n作為參數，返回1到n的所有數字的總和。
func SumN(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

// 創建一個Go結構體表示三角形，包含三個邊的長度屬性。實現一個方法來判斷三角形是否為等邊三角形。
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) JudgeEqual() bool {
	if t.SideA == t.SideB && t.SideC == t.SideB {
		return true
	}
	return false
}

// 編寫一個Go函數，接受一個字串作為參數，判斷該字串是否是回文（正向和反向讀取相同），並返回布爾值。
func JudgePalindrome(s string) bool {
	return s == ReverseStr(s)
}
