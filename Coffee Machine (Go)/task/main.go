package main

import (
	"fmt"
)

const (
	WATER = iota
	MILK
	BEANS
	MONEY
	CUPS
)

var (
	espresso   = [4]int{250, 0, 16, 4}
	latte      = [4]int{350, 75, 20, 7}
	cappuccino = [4]int{200, 100, 12, 6}
)

var machine = [5]int{400, 540, 120, 550, 9}

func writeRemaining() {
	fmt.Printf(`
The coffee machine has:
%d ml of water
%d ml of milk
%d g of coffee beans
%d disposable cups
$%d of money
`, machine[WATER], machine[MILK], machine[BEANS], machine[CUPS], machine[MONEY])
}

func checkForAvailability(machine *[5]int, coffee *[4]int) (bool, string) {
	switch {
	case machine[WATER] < coffee[WATER]:
		return false, "Sorry, not enough water!"
	case machine[MILK] < coffee[MILK]:
		return false, "Sorry, not enough milk!"
	case machine[BEANS] < coffee[BEANS]:
		return false, "Sorry, not enough coffee beans!"
	case machine[CUPS] < 1:
		return false, "Sorry, not enough cups!"
	default:
		return true, "I have enough resources, making you a coffee!"

	}
}

func buy() {
	var ch [4]int
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		ch = espresso
	case "2":
		ch = latte
	case "3":
		ch = cappuccino
	case "back":
		return
	}
	available, msg := checkForAvailability(&machine, &ch)
	if available {
		machine[WATER] -= ch[WATER]
		machine[MILK] -= ch[MILK]
		machine[BEANS] -= ch[BEANS]
		machine[CUPS]--
		machine[MONEY] += ch[MONEY]
	}
	fmt.Println(msg)

}

func fill() {
	fmt.Println("Write how many ml of water you want to add:")
	var amount int
	fmt.Scan(&amount)
	machine[WATER] += amount
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&amount)
	machine[MILK] += amount
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&amount)
	machine[BEANS] += amount
	fmt.Println("Write how many disposable cups you want to add: ")
	fmt.Scan(&amount)
	machine[CUPS] += amount
}

func take() {
	fmt.Printf("I gave you $%d\n", machine[MONEY])
	machine[MONEY] = 0
}

func main() {

L:
	for {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)
		switch action {
		case "buy":
			buy()
		case "fill":
			fill()
		case "take":
			take()
		case "remaining":
			writeRemaining()
		case "exit":
			break L
		}
	}

}
