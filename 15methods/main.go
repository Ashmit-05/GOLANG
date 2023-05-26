package main

import "fmt"

func main() {
	ashmit := User{"Ashmit", "ashmit@go.dev", false, 20, "walden"}
	fmt.Println("User : ", ashmit)
	ashmit.GetStatus()
	ashmit.changeMail("keating@helton") // note that we are passing it by reference. see the function
	fmt.Println("New mail ", ashmit.Email)
}

type User struct {
	Name     string
	Email    string
	Status   bool
	Age      int
	password string
}

func (u User) GetStatus() {
	fmt.Println("User status is : ", u.Status)
}

func (u *User) changeMail(newMail string) {
	u.Email = newMail
	fmt.Println("email changed successfully")
}
