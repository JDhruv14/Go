package main

import "fmt"

func split(sum int) (x, y int) {
	x = 2 * sum;
	y = 3 * sum;
	return;
}

func main() {
	fmt.Println(split(17));
}