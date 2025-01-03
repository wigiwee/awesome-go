package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"mode"`
	Password string   `json:"-"`              // - suggests the field to be ignored in json
	Tags     []string `json:"tags,omitempty"` //omitempty suggests that if the value is nil the field will not be shown
}

func main() {
	fmt.Println("Working with json in go")

	EncodeJson()
}

func EncodeJson() {

	course := []course{
		{"java", 344, "online", "abcaa", []string{"basic java", "beginner"}},
		{"python advance", 3465, "offline", "abcasdfa", nil},
		{"full stack", 789, "online", "alkjsdflakjsf", []string{"web-dev", "javascript", "react", "backend-frontend"}},
	}

	// finalJson, err := json.Marshal(course)
	finalJson, err := json.MarshalIndent(course, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)

}
