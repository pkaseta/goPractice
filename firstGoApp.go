package main

import "fmt"

func main() {
	firstName := "Paul"
	lastname := "Kaseta"

	var fullName string = firstName + " " + lastname

	birthMonth := "12"
	birthDate := "24"
	birthYear := "1984"

	var fullBirthday string = birthMonth + "/" + birthDate + "/" + birthYear
	fmt.Println("Name:" + " " + fullName)
	fmt.Println("Birthday:" + " " + fullBirthday)
}
