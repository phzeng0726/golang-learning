package helpers

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func toJson(jsonStr string) []Person {
	var results []Person

	err := json.Unmarshal([]byte(jsonStr), &results)

	if err != nil {
		log.Println("Error", err)
	}
	return results
}
