package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// var users = []User{
// 	User{ID: "1", Name: "Rita", Age: 26},
// 	User{ID: "2", Name: "Yong", Age: 27},
// 	User{ID: "3", Name: "Mona", Age: 23},
// }

// // C
// func createUser(c *gin.Context) {
// 	var newUser User

// 	if err := c.BindJSON(&newUser); err != nil {
// 		return
// 	}

// 	users = append(users, newUser)
// 	c.IndentedJSON(http.StatusCreated, users)
// }

// // R
// func getUsers(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, users)
// }

// // U
// func updateUser(c *gin.Context) {
// 	var updateUser User

// 	if err := c.BindJSON(&updateUser); err != nil {
// 		return
// 	}

// 	for i, u := range users {
// 		if updateUser.ID == u.ID {
// 			users = append(users[:i], users[i+1:]...)
// 			users = append(users, updateUser)
// 		}
// 	}

// 	c.IndentedJSON(http.StatusCreated, users)
// }

// // D
// func getUserIndexById(id string) (int, error) {
// 	for i, u := range users {
// 		if u.ID == id {
// 			return i, nil
// 		}
// 	}

// 	return -1, errors.New("User not found")
// }

// func deleteUser(c *gin.Context) {
// 	id, ok := c.GetQuery("id")
// 	if !ok {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
// 		return
// 	}

// 	i, err := getUserIndexById(id)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
// 		return
// 	}

// 	users = append(users[:i], users[i+1:]...)

// 	c.IndentedJSON(http.StatusOK, users)
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/users", getUsers)
// 	router.POST("/users", createUser)
// 	router.PATCH("/users", updateUser)
// 	router.DELETE("/users", deleteUser)

// 	port := "8080"
// 	fmt.Printf("Listening on port %s...\n", port)
// 	router.Run(":" + port)
// }
