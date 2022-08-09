package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice value is", diceNumber)

	// If we run this on a virtual playground then this will not work properly
	// because the random number in those VM's will be hardcoded and will give same output everytime

	switch diceNumber {
	case 1:
		fmt.Println("You can open the game")
	case 2:
		fmt.Println("You can move your piece by 2 spots")
	case 3:
		fmt.Println("You can move your piece by 3 spots")
		fallthrough // If case 3 is succeeded then case 4 will also be executed because of fallthrough command
	case 4:
		fmt.Println("You can move your piece by 4 spots")
	case 5:
		fmt.Println("You can move your piece by 5 spots")
	case 6:
		fmt.Println("You can move by 6 spots and roll the dice once again")
	default:
		fmt.Println("What was that?????")

	}
}
