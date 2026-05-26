package main

type Char struct {
	Stats
	Location
	Name             string
	LastName         string
	Age              int
	Wheight          int
	Length           int
	testItemInteract bool
	BackPack         map[string]itemObject
}

type Stats struct {
	Health       int
	Mana         int
	Strength     int
	Agility      int
	Intelligence int
	Faith        int
}

type Location struct {
	XAxis int
	YAxis int
}
type GameObject struct {
	Location
	Stats
	Name        string
	Discription string
	Interaction string
}

type itemObject struct {
	Location
	Stats
	Name        string
	Discription string
	Interaction string
	Damage      int
}
