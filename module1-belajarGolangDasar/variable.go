package main

import "fmt"

func main() {
	var name = "Test" //you don't need to initialize the variable data type if you already put in the value of the variable
	var nema string   //but you haven't put in any value in the variable, you need to initialize the data type
	nom := "Hey!"     //you could also use this if you don't want to use var, but only use :=  during initialization.

	var (
		firstName = "Daffa"
		lastName  = "Tarigan"
	) // you could also use this btw, to declare multiple variables

	fmt.Println(name)

	nema = "Hello!"

	fmt.Println(nema)

	fmt.Println(nom)

	fmt.Println(firstName, lastName)

}
