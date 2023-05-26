package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Seize the day, Gather ye Rosebuds while ye may!"
	file, err := os.Create("./test.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length of the file content is : ", length)
	defer file.Close()
	readFile(file.Name())
}

func readFile(filename string)  {
	data,err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("File data is :\n",string(data))
}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}
