package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main()  {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response is %T\n",response);
	fmt.Println(response)

	defer response.Body.Close()

	data,err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body : ",string(data))
}