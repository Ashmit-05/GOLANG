package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name     string   `json:"name"`
	Age      int      `json:"-"`                  // this field won't be shown in the json data
	Subjects []string `json:"subjects,omitempty"` // if empty, it won't show
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	students := []student{
		{"Ashmit", 20, []string{"aec", "dec", "ss"}},
		{"Atul", 19, nil},
		{"Mohit", 20, []string{"cs", "valorant"}},
	}

	// marshal is used to convert the given data to json format
	// it always takes an interface(a struct) as the input
	// it returns data in bytes, which needs to converted to string

	jsonData, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	// marshal indent formats the data to increase readability

	formattedJsonData, err := json.MarshalIndent(students, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(formattedJsonData))
}
