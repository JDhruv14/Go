package main

import "fmt"

fun main() {
	var num int;
	fmt.Println("Enter a number: ");
	fmt.Scanln(&num);
	if num % 2 == 0 {
		fmt.Println("Even");
	} else {
		fmt.Println("Odd");
	}
}