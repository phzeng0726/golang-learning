package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	User{ID: 1, Name: "Rita", Age: 26},
	User{ID: 2, Name: "Yong", Age: 27},
	User{ID: 3, Name: "Mona", Age: 23},
}

// C
func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, users)
}

// R
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// U
// TODO error handling
func updateUser(c *gin.Context) {
	var updateUser User

	if err := c.BindJSON(&updateUser); err != nil {
		return
	}

	for i, u := range users {
		if updateUser.ID == u.ID {
			users = append(users[:i], users[i+1:]...)
			users = append(users, updateUser)
		}
	}

	c.IndentedJSON(http.StatusCreated, users)
}

// D

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.PATCH("/users", updateUser)

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	router.Run(":" + port)
}
