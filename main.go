package main

import (
    "fmt"
)

func main() {
    scanLocation()
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

func playerChar() (charInfo, charInfo) {
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
            xAxis: 1,
            yAxis: 1,
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
    fmt.Println("if you want it to stop write stop")
    sum := 0
    fmt.Scanln(&inpTest)
    // if inpTest == "char1" {
    for i := 0; i < 100; i++ {

        if inpTest == "w" {
            charOne.location.yAxis += moveW
            fmt.Println(charOne.location.xAxis, charOne.location.yAxis)
        }
        if inpTest == "s" {
            charOne.location.yAxis += moveS
            fmt.Println(charOne.location.xAxis, charOne.location.yAxis)
        }
        if inpTest == "a" {
            charOne.location.xAxis += moveS
            fmt.Println(charOne.location.xAxis, charOne.location.yAxis)
        }
        if inpTest == "d" {
            charOne.location.xAxis += moveW
            fmt.Println(charOne.location.xAxis, charOne.location.yAxis)
        }
        sum += i
        if inpTest == "stop" {
            break
        }

    }
    // }
}


func scanLocation() {
    scan := staticObjects()

    if scan <
    fmt.Println(scan.xAxis)
}

func staticObjects() staticObjectInfo {
    rock := staticObjectInfo{
        name:        "Rock",
        discription: "has been untuched for many years, moss is starting to grow.",
        staticObjectlocation: staticObjectlocation{
            xAxis: 20,
            yAxis: 15,
        },
    }
    return rock
}


func charMovmentForward(i int) int {
    i++
    return i
}
func charMovmentBackward(i int) int {
    i--
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
type staticObjectInfo struct {
    staticObjectlocation
    name        string
    discription string
}

type staticObjectlocation struct {
    xAxis int
    yAxis int
}
