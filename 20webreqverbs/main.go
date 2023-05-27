package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// GetRequest("http://localhost:8080/")
	// PostJsonRequest("http://localhost:8080/post")
	PostFormRequest("http://localhost:8080/postform")
}

func GetRequest(test_url string) {
	response, err := http.Get(test_url)
	CheckNilError(err)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	CheckNilError(err)
	fmt.Println(response.Status)
	fmt.Println(response.ContentLength)
	fmt.Println(string(data))

	// another way of converting your data to string
	// although not beginner friendly, this method is much more powerful
	var responseString strings.Builder
	byteCount, _ := responseString.Write(data)
	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String())
}

func PostJsonRequest(test_url string) {
	// this is a fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"Henry David Thoreau",
			"message":"What devil possessed me that I behaved so well?"
		}
	`)
	response, err := http.Post(test_url, "application/json", requestBody)
	CheckNilError(err)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	CheckNilError(err)
	fmt.Println(string(data))
}

func PostFormRequest(test_url string) {
	data := url.Values{}
	data.Add("firstname", "Vijay")
	data.Add("lastname", "Chauhan")
	data.Add("middlename", "Dinanath")

	response, err := http.PostForm(test_url, data)
	CheckNilError(err)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
