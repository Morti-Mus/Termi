package main

import (
	"fmt"
)

var house = GameObject{}

func main() {
	attackwithItemChar()
	// CreateStaticItemObject()
	// CreateStaticObjects()
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
		Name:             "John",
		LastName:         "Doe",
		Age:              30,
		Wheight:          180,
		Length:           75,
		testItemInteract: true,
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

		if inpTest == "inspect" {
			ItemCharInspect(charOne)
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

	if charOne.Location.XAxis+2 >= objectArray[0].XAxis &&
		charOne.Location.XAxis-2 <= objectArray[0].XAxis &&
		charOne.Location.YAxis+2 >= objectArray[0].YAxis &&
		charOne.Location.YAxis-2 <= objectArray[0].YAxis {

		fmt.Println("\n", objectArray[0].Name, objectArray[0].Discription)

	} else {
		fmt.Printf("you see something in the distance ... go closer to investigate! %vXaxis %vYaxis", objectArray[0].XAxis, objectArray[0].YAxis)
	}
	if charOne.Location.XAxis+2 >= objectArray[1].XAxis &&
		charOne.Location.XAxis-2 <= objectArray[1].XAxis &&
		charOne.Location.YAxis+2 >= objectArray[1].YAxis &&
		charOne.Location.YAxis-2 <= objectArray[1].YAxis {

	} else {
		fmt.Printf("you see something in the distance ... go closer to investigate! %vXaxis %vYaxis", objectArray[1].XAxis, objectArray[1].YAxis)
	}
}

func LocationChecker() bool { // kolla om detta går att göra bättre med en loop som kollar allt istället?
	var objectArray []GameObject = CreateStaticObjects()
	charOne, _ := playerChar()

	locationCharCheck := [2]int{charOne.Location.XAxis, charOne.Location.YAxis}
	locationObjectCheck := [2]int{objectArray[0].Location.XAxis, objectArray[0].Location.YAxis}
	locationCombinationCheckPlus := locationCharCheck[0]+2 >= locationObjectCheck[0]
	locationCombinationCheckMinus := locationCharCheck[0]-2 <= locationObjectCheck[0]

	finalLocationCombinationCheck := locationCombinationCheckPlus && locationCombinationCheckMinus
	fmt.Println(finalLocationCombinationCheck)
	return finalLocationCombinationCheck
}

func CreateStaticObjects() []GameObject {
	rockObject := NewStaticObjects("rock",
		"The stone looks aged. \n the sun has bleached the surface. \n you can se a small streak of gold in crack. \n maybe this can be harvested...",
		"",
		Location{XAxis: 15, YAxis: 20},
		Stats{Health: 5})
	treeObject := NewStaticObjects("tree",
		"The tree stand tall. \n you can see the leaves russtle in the wind. \n at the bottom you can se that some one has etched in some cordinateds \n XAxis 40 YAxis 50",
		"",
		Location{XAxis: 10, YAxis: 20},
		Stats{Health: 5})
	housObject := NewStaticObjects("Red Tavern",
		"The House has been standing for many years \n you can hear music and laughter coming from the windows \n take drink with the patrons",
		"",
		Location{XAxis: 45, YAxis: 30},
		Stats{Health: 5})
	snakeOilSalesMan := NewStaticObjects("Jimmy sketch",
		"You see a man with a table and a sign \n on the sign it says \n Jimmys sketchy stuff \n he seems untrust worthy ...",
		"",
		Location{XAxis: 4, YAxis: 4},
		Stats{Health: 30})

	objectArray := [...]GameObject{rockObject, treeObject, housObject, snakeOilSalesMan}

	return objectArray[:]
}

func CreateStaticItemObject() map[string]itemObject { // detta måste gå att göra på något bättre sätt en tidigare
	itemArray := map[string]itemObject{}

	swordItem := itemObject{
		Name:        "Zweihander",
		Discription: "A big sword from German origin",
		Interaction: "You have picked upp the sword \n its remarcably heavy.",
		Damage:      30,
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}
	stickItem := itemObject{
		Name:        "Stick",
		Discription: "of truth",
		Interaction: "This stick radiats unspeakble power",
		Damage:      10,
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}

	itemArray[stickItem.Name] = stickItem // kolla hur fan detta igentligen funkar.
	itemArray[swordItem.Name] = swordItem

	// fmt.Printf("%+v", testItem)
	// fmt.Printf("swordItem\t%v\n", swordItem.Name)
	// fmt.Println(testItem[swordItem.Name])
	return itemArray

}

func NewStaticObjects(Name, Discription, Interaction string, Location Location, Health Stats) GameObject {
	basicstaticObejctArc := GameObject{
		Name:        Name,
		Discription: Discription,
		Location:    Location,
		Interaction: Interaction,
		Stats:       Health,
	}

	return basicstaticObejctArc
}

//https://www.geeksforgeeks.org/go-language/delete-elements-in-a-slice-in-golang/ länk till slice remove func

func attackwithItemChar(charOne Char) {
	staticObjects := CreateStaticObjects()

	var index int = 3
	elem := staticObjects[index]

	fmt.Println(elem)

	if charOne.XAxis == staticObjects[3].XAxis && charOne.BackPack["Zweihander"].Damage < staticObjects[3].Health {
		staticObjects = append(staticObjects[:index], staticObjects[index+1:]...)
	}
}

func itemToBackPack() {
	charOne, _ := playerChar()

	testItemArray := CreateStaticItemObject()
	Zweihander := testItemArray["Zweihander"]

	if charOne.XAxis == Zweihander.XAxis {
		charOne.BackPack[Zweihander.Name] = Zweihander

		fmt.Printf("you have added %v to your backpack", charOne.BackPack["Zweihander"].Name)
	}
	fmt.Println(testItemArray["Zweihander"])
	fmt.Println(charOne)
	// CreateStaticItemObject()
}

func teststuff() {
	testArr := []int{}
	testArr = append(testArr, 1)
	fmt.Println(testArr[0])
}

func ItemCharInspect(charOne Char) {
	var itemArray = CreateStaticItemObject()

	if charOne.testItemInteract == true &&
		charOne.Location.XAxis == itemArray["Zweihander"].Location.XAxis &&
		charOne.Location.YAxis == itemArray["Zweihander"].Location.YAxis {
		fmt.Println(itemArray["Zweihander"])
	} else {
		fmt.Println("well you arent in the correct space")
	}

}

func ObjectToCharinteract(charOne Char) {
	var GameObject []GameObject = CreateStaticObjects()
	dialogInteraction := "Hello there sir... \n You look like you have some coin to offer for my snake oil"

	fmt.Println(GameObject[3].YAxis)
	if charOne.Location.XAxis == GameObject[3].XAxis && charOne.Location.YAxis == GameObject[3].YAxis {
		fmt.Println(dialogInteraction)
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
