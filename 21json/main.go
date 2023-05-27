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
	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`
		{
			"name": "Mohit",
			"subjects": ["cs","valorant"]
		}
	`)
	checkValid := json.Valid(jsonData)

	// when you want to store the data in a struct
	var SingleStudent student
	if checkValid {
		json.Unmarshal(jsonData, &SingleStudent)
		fmt.Printf("%#v\n", SingleStudent) // when using %v with a struct, you need to use # as well
	} else {
		fmt.Println("json data was not valid")
	}

	// when you want to store the data in a key value pair(map)
	/*
	The key is always going to be a string, but the value might be
	an integer, string, array or anything else. For this reason,
	when dealing with web requests and mapping, we use string for key
	and interface(empty) for values
	*/

	var mappedData map[string]interface{} 
	json.Unmarshal(jsonData,&mappedData)
	fmt.Printf("%#v\n",mappedData)

	// looping through the map
	for key,value := range mappedData {
		fmt.Printf("Key : %v, Value : %v\n",key,value)
	}
}
