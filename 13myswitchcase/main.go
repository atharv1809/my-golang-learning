package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice number is:", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can open.")
	case 2:
		fmt.Println("Dice value is 2, you can move 2 spot.")
	case 3:
		fmt.Println("Dice value is 3, you can move 3 spot.")
	case 4:
		fmt.Println("Dice value is 4, you can move 4 spot.")
		// fallthrough
		// fallthrough anywhere in switch case will cause switch case to execute the current case and all the below cases
	case 5:
		fmt.Println("Dice value is 5, you can move 5 spot.")
	case 6:
		fmt.Println("Dice value is 6, you can roll again.")
	default:
		fmt.Println("What was that?")
	}
}
