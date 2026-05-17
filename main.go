package main

import (
	"fmt"
)

var house = GameObject{}

func main() {
	CreateStaticObjects()
	// // staticObjects()
	// // fmt.Println()

	// scanLocation()
	picker()
	charMovment()
}

func picker() {
	charOne, charTwo := playerChar()

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

func playerChar() (Char, Char) {
	char1 := Char{
		name:     "John",
		lastName: "Doe",
		age:      30,
		wheight:  180,
		length:   75,
		Stats: Stats{
			health:       100,
			mana:         50,
			strength:     10,
			agility:      8,
			intelligence: 7,
			faith:        5,
		},
		Location: Location{
			xAxis: 1,
			yAxis: 1,
		},
	}
	char2 := Char{

		name:     "Jane",
		lastName: "Doel",
		age:      25,
		wheight:  165,
		length:   60,
		Stats: Stats{
			health:       80,
			mana:         70,
			strength:     8,
			agility:      10,
			intelligence: 9,
			faith:        6,
		},
		Location: Location{
			xAxis: 1,
			yAxis: 1,
		},
	}

	return char1, char2
}

// connect this func to picker() so that the char you use get used in this func instead of pre determined as of right now charOne

func charMovment() {
	charOne, _ := playerChar()
	var moveW int = charMovmentForward(0)
	var moveS int = charMovmentBackward(0)
	var inpTest string
	// var stuff bool

	fmt.Println("Pleas input w for North")
	fmt.Println("Pleas input s for South")
	fmt.Println("Pleas input a for West")
	fmt.Println("Pleas input d for East")
	fmt.Println("if you want to scan for any locations or objects write scan")
	fmt.Println("if you want it to stop write stop")
	// add error handeling to this code.
	// if inpTest == "char1" {
	for {
		fmt.Scanln(&inpTest)

		if inpTest == "w" {
			charOne.Location.yAxis += moveW
			fmt.Println(charOne.Location.xAxis, "x", charOne.Location.yAxis, "y")
			continue
		}
		if inpTest == "s" {
			charOne.Location.yAxis += moveS
			fmt.Println(charOne.Location.xAxis, "x", charOne.Location.yAxis, "y")
			continue
		}
		if inpTest == "a" {
			charOne.Location.xAxis += moveS
			fmt.Println(charOne.Location.xAxis, "x", charOne.Location.yAxis, "y")
			continue
		}
		if inpTest == "d" {
			charOne.location.xAxis = moveW
			fmt.Println(charOne.location.xAxis, charOne.location.yAxis)
		}
		sum += i
		// if stuff = (false){
		// 	break
		// }
	}
}

func charMovmentForward(i int) int {
	i++
	return i
}
func charMovmentBackward(i int) int {
	i--
	return i
}

type Char struct {
	Stats
	Location
	name     string
	lastName string
	age      int
	wheight  int
	length   int
}

type Stats struct {
	health       int
	mana         int
	strength     int
	agility      int
	intelligence int
	faith        int
}

type Location struct {
	xAxis int
	yAxis int
}
