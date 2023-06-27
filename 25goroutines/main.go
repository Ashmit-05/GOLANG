package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"signals"}

var wg sync.WaitGroup // usually this is a pointer
var mut sync.Mutex // usually this is a pointer

func main() {
	//go greeter("Hello")
	//greeter("World")
    websiteList := []string{
        "https://lco.dev",
        "https://google.com",
        "https://api.sellular.club",
        "https://vercel.sellular.club",
    }

    for _, website := range websiteList {
        go getStatusCode(website)
        wg.Add(1)
    }

    wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
    defer wg.Done()
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
    
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d is the status for %s\n", res.StatusCode, endpoint)
}
