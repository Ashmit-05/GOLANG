package main

import (
	"fmt"
	"net/url"
)

const test_url string = "https://lco.dev:3000/learn?coursename=nodejs&quality=great"

func main() {
	// deconstructing a url
	result, _ := url.Parse(test_url)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams)

	for key, val := range qparams {
		fmt.Printf("Key : %v, Value : %v\n", key, val)
	}

	// constructing a url
	partsOfUrl := &url.URL{ // remember that the & sign is very important
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=ashmit",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
