package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const url1 string = "https://www.google.com"
const url2 string = "https://www.google.com:8080/search?q=hello+world&language=English&browser=firefox"
const url3 = "http://localhost:8080/api/v1/components"
const url4 = "http://localhost:8080/api/v1/auth/login"

func main() {
	fmt.Println("")

	//sending a GET request
	respones, err := http.Get(url1)
	if err != nil {
		panic(err)
	}
	defer respones.Body.Close() //callers responsibility to clsoe the connection

	fmt.Printf("Response is of type :%T \n ", respones)

	databytes, err := io.ReadAll(respones.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(databytes)

	fmt.Println(string(databytes))

	//handling/parsing url
	result, _ := url.Parse(url2)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) //returns all query parameters in unstructured manner

	// to get query parameteres in structrued manner
	query_params := result.Query() //query_params is a map
	fmt.Printf("Type of query_param is %T \n", query_params)
	fmt.Println(query_params)
	fmt.Println("query broswer is ", query_params.Get("browser"))

	//building url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/search",
		RawPath: "test",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

	sendGetRequest()

	sendPostRequest()

	sendPostFormRequest()
}

func sendGetRequest() {

	response, err := http.Get(url3)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status : ", response.Status)
	fmt.Println("status code : ", response.StatusCode)
	fmt.Println("content length : ", response.ContentLength)

	var responesString strings.Builder
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	byteCount, _ := responesString.Write(content)
	fmt.Println("byte Count is : ", byteCount)
	//another way to get response data ( this method provieds more functionality, responseString supports a lot of methods)
	fmt.Println("response data is : ", responesString.String())

	fmt.Println(string(content)) //one way to get response data

}

// sends json body
func sendPostRequest() {

	requestBody := strings.NewReader(`
		{
			"username":"tl@pravaraengg.org.in",
    		"password":"admin"
		}
	`)

	response, err := http.Post(url4, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

// sends form body
func sendPostFormRequest() {

	data := url.Values{}
	data.Add("firstName", "Bob")
	data.Add("lastname", "Builder")
	data.Add("email", "bob@gmailcom")

	response, err := http.PostForm(url4, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
