package main

import (
	"flag"
	"fmt"
)

// 實作一個簡單的命令列計算器，支持加法、減法、乘法和除法操作，使用 flag 庫來解析命令列參數。
// go build main.go
// .\main.exe -num1 10 -op add -num2 20
func main() {
	// 定義命令列參數
	operator := flag.String("op", "", "Operation: add, subtract, multiply, divide")
	num1 := flag.Float64("num1", 0.0, "First number")
	num2 := flag.Float64("num2", 0.0, "Second number")

	// 解析命令列參數
	flag.Parse()

	// 執行相應的計算
	switch *operator {
	case "add":
		result := *num1 + *num2
		fmt.Printf("Result: %f\n", result)
	case "subtract":
		result := *num1 - *num2
		fmt.Printf("Result: %f\n", result)
	case "multiply":
		result := *num1 * *num2
		fmt.Printf("Result: %f\n", result)
	case "divide":
		if *num2 != 0 {
			result := *num1 / *num2
			fmt.Printf("Result: %f\n", result)
		} else {
			fmt.Println("Cannot divide by zero.")
		}
	default:
		fmt.Println("Invalid operation.")
	}
}
