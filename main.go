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
			charOne.Location.xAxis += moveW
			fmt.Println(charOne.Location.xAxis, "x", charOne.Location.yAxis, "y")
			continue
		}

		if inpTest == "scan" {
			scanLocation()
		}

		if inpTest == "stop" {
			break
		}

	}
	// }
}

func scanLocation() {
	var objectArray []GameObject = CreateStaticObjects()
	charOne, _ := playerChar()

	if charOne.Location.xAxis+2 >= objectArray[0].xAxis && charOne.Location.xAxis-2 <= objectArray[0].xAxis { // fixa if logicen med greater eller smaller then det kommer va mek
		fmt.Println("\n", objectArray[0].name, objectArray[0].discription)
	} else {
		fmt.Printf("you see something in the distance ... go closer to investigate! %vXaxis %vYaxis", objectArray[0].xAxis, objectArray[0].yAxis)
	}
	if charOne.Location.xAxis+2 >= objectArray[1].xAxis && charOne.Location.xAxis-2 <= objectArray[1].xAxis {

	} else {
		fmt.Printf("you see something in the distance ... go closer to investigate! %vXaxis %vYaxis", objectArray[1].xAxis, objectArray[1].yAxis)
	}
}

func CreateStaticObjects() []GameObject {
	rockObject := NewStaticObjects("rock", "The stone looks aged. \n the sun has bleached the surface. \n you can se a small streak of gold in crack. \n maybe this can be harvested...", Location{xAxis: 15, yAxis: 20})
	treeObject := NewStaticObjects("tree", "The tree stand tall. \n you can see the leaves russtle in the wind. \n at the bottom you can se that some one has etched in some cordinateds \n xAxis 40 yAxis 50", Location{xAxis: 10, yAxis: 20})
	housObject := NewStaticObjects("Red Tavern", "The House has been standing for many years \n you can hear music and laughter coming from the windows \n take drink with the patrons", Location{xAxis: 45, yAxis: 30})
	snakeOilSalesMan := NewStaticObjects("Jimmy sketch", "You see a man with a table and a sign \n on the sign it says \n Jimmys sketchy stuff \n he seems untrust worthy ...", Location{xAxis: 25, yAxis: 25})
	objectArray := [...]GameObject{rockObject, treeObject, housObject, snakeOilSalesMan}

	return objectArray[:]
}

func NewStaticObjects(name, discription string, stats Location) GameObject {
	basicstaticObejctArc := GameObject{
		name:        name,
		discription: discription,
		Location:    stats,
	}

	return basicstaticObejctArc
}

func ObjectToCharinteract() {
	var test []GameObject = CreateStaticObjects()
	var charOne, _ Char = playerChar()
	fmt.Println(test[3])
	fmt.Println(charOne)

	// här vill jag koppla ihop char och snakesalesman så att dom gör något eller ändrar något
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
type GameObject struct {
	Location
	name        string
	discription string
}

// type StaticObjectlocation struct {
// 	xAxis int
// 	yAxis int
// }
