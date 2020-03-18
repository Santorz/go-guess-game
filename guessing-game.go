package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number")

	//*generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) //generates numbers between zero and n(10)

	var guess int
	var trials = 5
	var playcount = 1
	for {

		fmt.Println("You have ", +trials, "trial(s) ! \n",
			"Please input your guess: ")
		fmt.Scan(&guess)
		fmt.Println("N/B:You've played ",
			+playcount, " time(s) ! ")
		if guess > secretNumber {
			fmt.Println("Too big!  ")
		} else if guess < secretNumber {
			fmt.Println("Too small!  ")
		} else if guess == secretNumber {
			fmt.Println("You Win!  ")
		}
		playcount++
		trials--
		if playcount > 5 || trials < 1 {
			fmt.Println("Out of trials !")
			break
		}
		if guess == secretNumber {
			break
		}
	}
}
