package main

import (
	"fmt"
)

func main() {
	picker()

}

func picker() {
	charOne, charTwo := char()
	var inpTest string
	fmt.Scanln(&inpTest)
	sum := 0
	for i := 0; i < 10; i++ {
		if inpTest == "w" {
			move := charMovment(1)

			charOne.location.xAxis = move
			fmt.Println(charOne.location.xAxis)
		}
		sum += i
	}
	// bool1 := true
	// bool2 := false

	fmt.Print("Pick a character (1/2): ")
	var input string
	fmt.Scanln(&input)

	if input == "1" {
		fmt.Printf("You picked: %v, %s\n", charOne.age, charOne.name)
	} else if input == "2" {
		fmt.Printf("You picked: %s, %s\n", charTwo.lastName, charTwo.name)
	} else {
		fmt.Println("Invalid input.")
	}
}

func char() (charInfo, charInfo) {
	char1 := charInfo{
		name:     "John",
		lastName: "Doe",
		age:      30,
		wheight:  180,
		length:   75,
		charStats: charStats{
			health:       100,
			mana:         50,
			strength:     10,
			agility:      8,
			intelligence: 7,
			faith:        5,
		},
		location: location{
			xAxis: 0,
			yAxis: 0,
		},
	}
	char2 := charInfo{

		name:     "Jane",
		lastName: "Doel",
		age:      25,
		wheight:  165,
		length:   60,
		charStats: charStats{
			health:       80,
			mana:         70,
			strength:     8,
			agility:      10,
			intelligence: 9,
			faith:        6,
		},
		location: location{
			xAxis: 1,
			yAxis: 1,
		},
	}

	return char1, char2

}

func charMovment(i int) int {
	i++
	return i
}

type charInfo struct {
	charStats
	location
	name     string
	lastName string
	age      int
	wheight  int
	length   int
}

type charStats struct {
	health       int
	mana         int
	strength     int
	agility      int
	intelligence int
	faith        int
}

type location struct {
	xAxis int
	yAxis int
}
