package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10 // less is more difficult
	usage    = ` 
Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is,the harder it gets.
Wanna play?
`
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("Not a number")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick up a positive number")
		return
	}

	t := time.Now().Second()
	rand.Seed(int64(t))

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1
		fmt.Printf(" num%d \n", n)

		if n != guess {
			continue
		}
		if turn == 0 {
			fmt.Printf(`🥇 FIRST TIME WINNER!!! you won at turn%d `, turn+1)
		} else {
			fmt.Printf(" 🎉 Congratulation YOU WON at turn%d", turn+1)
		}
		return

	}

	fmt.Println("☠️ YOU LOST... Try again?")
}
