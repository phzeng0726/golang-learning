package main

import (
	"log"
)

func main() {
	myJson := `
[
    {
        "first_name" : "Clark",
        "last_name" : "Kent",
        "hair_color" : "black",
        "has_dog" : true
    },
    {
        "first_name" : "RRA",
        "last_name" : "AASQ",
        "hair_color" : "blue",
        "has_dog" : false
    }
]`

	results := helpers.toJson(myJson)

	log.Println("Successed", results)
}
