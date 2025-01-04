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

	DecodeJson()
}

func EncodeJson() {

	course := []course{
		{"java", 344, "online", "abcaa", []string{"basic java", "beginner"}},
		{"python advance", 3465, "offline", "abcasdfa", nil},
		{"full stack", 789, "online", "alkjsdflakjsf", []string{"web-dev", "javascript", "react", "backend-frontend"}},
	}

	finalJson, err := json.Marshal(course)
	// finalJson, err := json.MarshalIndent(course, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`{"coursename":"java","Price":344,"mode":"online","tags":["basic java","beginner"]}`)
	jsonDataFromWeb2 := []byte(`[{"coursename":"java","Price":344,"mode":"online","tags":["basic java","beginner"]},{"coursename":"python advance","Price":3465,"mode":"offline"},{"coursename":"full stack","Price":789,"mode":"online","tags":["web-dev","javascript","react","backend-frontend"]}]`)
	//verify json integrity
	var courses course

	checkValid := json.Valid(jsonDataFromWeb)
	fmt.Println(checkValid)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v \n ", courses)
	} else {
		fmt.Println("invalid json")
	}

	//some cases where you just want to add data to json
	// var courseMap map[string]interface{} //map for holding json data with single obj
	var courseMap []map[string]interface{} //slice of map for holding json with multiple obj
	json.Unmarshal(jsonDataFromWeb2, &courseMap)
	fmt.Printf("%#v \n ", courseMap)

}
