package main

import (
	"fmt"
)

func main() {
	sum := 10
	for sum != 0{
		sum -= 1;
		fmt.Printf("%v ",sum);
	}
}