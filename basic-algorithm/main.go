package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {
	router := gin.Default()

	// 新增用戶
	router.POST("/users", createUser)

	// 讀取所有用戶
	router.GET("/users", getUsers)

	// 讀取單個用戶
	router.GET("/users/:id", getUser)

	// 更新用戶
	router.PUT("/users/:id", updateUser)

	// 刪除用戶
	router.DELETE("/users/:id", deleteUser)

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	router.Run(":" + port)
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = user.ID
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// func main() {
// 	// // util_01
// 	// // 1
// 	// questions.Sum(12, 30)
// 	// // 2
// 	// data := []int{1, 2, 3, 4, 5}
// 	// questions.FindMaxAndMin(data)
// 	// // 3
// 	// car01 := questions.Car{Brand: "Nissan", Category: "c750"}
// 	// car01.Introduce()
// 	// // 4
// 	// questions.ReverseStr("Hot dog")

// 	// // util_02
// 	// // 1
// 	// questions.FindEven(10)
// 	// // 2
// 	// rectangle01 := questions.Rectangle{Length: -10, Width: 20}
// 	// fmt.Println(rectangle01.Area())
// 	// // 3
// 	// strSlice := []string{"pipi", "mina", "omicorn"}
// 	// fmt.Println(questions.SumStrLen(strSlice))

// 	// // util_03
// 	// // 01
// 	// fmt.Println(questions.SumN(3))
// 	// // 02
// 	// triangle := questions.Triangle{SideA: 10, SideB: 10, SideC: 10}
// 	// fmt.Println(triangle.JudgeEqual())
// 	// // 03
// 	// fmt.Println(questions.JudgePalindrome("iui"))

// 	// // util_04
// 	// // 01
// 	// fmt.Println(questions.Factorial(3))
// 	// // 02
// 	// cylinder := questions.Cylinder{Radius: 10, Height: 10}
// 	// fmt.Println(cylinder.Volume())
// 	// // 03
// 	// fmt.Println(questions.Fibonacci(40))

// 	// // util_05
// 	// // 01
// 	// fmt.Println(questions.FindPrime(21))
// 	// // 02
// 	// employee := questions.Employee{Name: "Rita", Age: 93}
// 	// fmt.Println(employee.JudgeRetire())
// 	// // 03
// 	// fmt.Println(questions.CountVowel("Rita"))

// 	// util_06
// 	// 01
// 	questions.HttpServer()
// 	// // 02
// 	// employee := questions.Employee{Name: "Rita", Age: 93}
// 	// fmt.Println(employee.JudgeRetire())
// 	// // 03
// 	// fmt.Println(questions.CountVowel("Rita"))
// }
