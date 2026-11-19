package main

import "fmt"

func main() {
	fmt.Println("Test")
	fmt.Println("Hi")
	fmt.Println("TestTest"[5]) // this will fetch the byte instead of string, so don't get confused on why it became a number
	fmt.Println(len("Test"))
}
