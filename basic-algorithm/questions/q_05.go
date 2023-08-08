package questions

import "strings"

// 編寫一個Go函數，接受一個整數n作為參數，返回n以內的所有質數。
func FindPrime(n int) []int {
	prime := []int{2} // 第一個質數是2

	// 除了2的偶數都非質數
	// 質數不會被小於自己的質數整除
	for i := 3; i <= n; i += 2 {
		// 所有目前找到的質數都不整除（有餘數）的時候，代表為質數
		temp := []int{2}
		for _, p := range prime {
			if i%p != 0 {
				temp = append(temp, i)
			}
		}
		if len(temp) == len(prime)+1 {
			prime = append(prime, i)
		}

	}

	return prime
}

// 創建一個Go結構體表示員工，包含姓名和年齡屬性。實現一個方法來判斷員工是否退休（年齡大於等於65）。
type Employee struct {
	Name string
	Age  int
}

func (e Employee) JudgeRetire() bool {
	return e.Age >= 65
}

// 編寫一個Go函數，接受一個字串作為參數，返回該字串中所有元音字母的數量。
func CountVowel(str string) int {
	vowelCount := 0

	for _, s := range str {
		// 使用 strings.ContainsAny() 來檢查是否包含元音字母
		if strings.ContainsAny(string(s), "aeiouAEIOU") {
			vowelCount++
		}
	}

	return vowelCount
}
