package main

import "fmt"

func main() {
	fmt.Println(sayHi("Marco"))
	fmt.Println("say hello")
	fmt.Println("say hello222")
}

func sayHi(person string) string {
	return fmt.Sprintf("Hi %s", person)
}
