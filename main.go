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
		fmt.Printf("You picked: %v, %s\n", charOne.Age, charOne.Name)
	} else if input == "2" {
		fmt.Printf("You picked: %s, %s\n", charTwo.LastName, charTwo.Name)
	} else {
		fmt.Println("Invalid input.")
	}
}

func playerChar() (Char, Char) {
	char1 := Char{
		Name:     "John",
		LastName: "Doe",
		Age:      30,
		Wheight:  180,
		Length:   75,
		Stats: Stats{
			Health:       100,
			Mana:         50,
			Strength:     10,
			Agility:      8,
			Intelligence: 7,
			Faith:        5,
		},
		Location: Location{
			XAxis: 1,
			YAxis: 1,
		},
	}
	char2 := Char{

		Name:     "Jane",
		LastName: "Doel",
		Age:      25,
		Wheight:  165,
		Length:   60,
		Stats: Stats{
			Health:       80,
			Mana:         70,
			Strength:     8,
			Agility:      10,
			Intelligence: 9,
			Faith:        6,
		},
		Location: Location{
			XAxis: 1,
			YAxis: 1,
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
			charOne.Location.YAxis += moveW
			fmt.Println(charOne.Location.XAxis, "x", charOne.Location.YAxis, "y")
			continue
		}
		if inpTest == "s" {
			charOne.Location.YAxis += moveS
			fmt.Println(charOne.Location.XAxis, "x", charOne.Location.YAxis, "y")
			continue
		}
		if inpTest == "a" {
			charOne.Location.XAxis += moveS
			fmt.Println(charOne.Location.XAxis, "x", charOne.Location.YAxis, "y")
			continue
		}
		if inpTest == "d" {
			charOne.Location.XAxis += moveW
			fmt.Println(charOne.Location.XAxis, "x", charOne.Location.YAxis, "y")
			continue
		}

		if inpTest == "scan" {
			scanLocation()
		}

		if inpTest == "talk" {
			ObjectToCharinteract(charOne)
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

	if charOne.Location.XAxis+2 >= objectArray[0].XAxis && charOne.Location.XAxis-2 <= objectArray[0].XAxis { // fixa if logicen med greater eller smaller then det kommer va mek
		fmt.Println("\n", objectArray[0].Name, objectArray[0].Discription)
	} else {
		fmt.Printf("you see something in the distance ... go closer to investigate! %vXaxis %vYaxis", objectArray[0].XAxis, objectArray[0].YAxis)
	}
	if charOne.Location.XAxis+2 >= objectArray[1].XAxis && charOne.Location.XAxis-2 <= objectArray[1].XAxis {

	} else {
		fmt.Printf("you see something in the distance ... go closer to investigate! %vXaxis %vYaxis", objectArray[1].XAxis, objectArray[1].YAxis)
	}
}

func CreateStaticObjects() []GameObject {
	rockObject := NewStaticObjects("rock", "The stone looks aged. \n the sun has bleached the surface. \n you can se a small streak of gold in crack. \n maybe this can be harvested...", "", Location{XAxis: 15, YAxis: 20})
	treeObject := NewStaticObjects("tree", "The tree stand tall. \n you can see the leaves russtle in the wind. \n at the bottom you can se that some one has etched in some cordinateds \n XAxis 40 YAxis 50", "", Location{XAxis: 10, YAxis: 20})
	housObject := NewStaticObjects("Red Tavern", "The House has been standing for many years \n you can hear music and laughter coming from the windows \n take drink with the patrons", "", Location{XAxis: 45, YAxis: 30})
	snakeOilSalesMan := NewStaticObjects("Jimmy sketch", "You see a man with a table and a sign \n on the sign it says \n Jimmys sketchy stuff \n he seems untrust worthy ...", "", Location{XAxis: 4, YAxis: 4})
	objectArray := [...]GameObject{rockObject, treeObject, housObject, snakeOilSalesMan}

	return objectArray[:]
}

func NewStaticObjects(Name, Discription, Interaction string, stats Location) GameObject {
	basicstaticObejctArc := GameObject{
		Name:        Name,
		Discription: Discription,
		Location:    stats,
		Interaction: Interaction,
	}

	return basicstaticObejctArc
}

func ObjectToCharinteract(charOne Char) {
	var GameObject []GameObject = CreateStaticObjects()
	testOutput := "Test"
	fmt.Println("tester")
	fmt.Println(GameObject[3].YAxis)
	if charOne.Location.XAxis == GameObject[3].XAxis && charOne.Location.YAxis == GameObject[3].YAxis {
		fmt.Println(testOutput)
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

// type Char struct {
// 	Stats
// 	Location
// 	Name     string
// 	LastName string
// 	Age      int
// 	Wheight  int
// 	Length   int
// }

// type Stats struct {
// 	Health       int
// 	Mana         int
// 	Strength     int
// 	Agility      int
// 	Intelligence int
// 	Faith        int
// }

// type Location struct {
// 	XAxis int
// 	YAxis int
// }
// type GameObject struct {
// 	Location
// 	Name        string
// 	Discription string
// 	Interaction string
// }

// type StaticObjectlocation struct {
// 	XAxis int
// 	YAxis int
// }
