/*
	Hello, in this game you must guess the number, which pc generates.

You have 10 attempts.
Good luck.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	number := rand.Intn(101)
	fmt.Println(number)
	attempt := 0

	for attempt = 0; attempt <= 10; attempt++ {
		fmt.Println("Please enter the number")
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.Trim(guess, "\n")
		guessNum, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if guessNum < number {
			fmt.Println("Your number is LOW")
		} else if guessNum > number {
			fmt.Println("Your number is HIGH")
		} else {
			fmt.Println("Congatulations, you win this game!!")
			break
		}
	}
}
