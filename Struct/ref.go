package main

import "fmt"

type Person struct {
	Name string
	Surname string
}

func main() {
	fmt.Println(Person{"Dhruv","Jaradi"});
}