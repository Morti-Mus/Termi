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

// func (c Char) UseItem(itemName string, target Char) {
// 	selectedItem, has := c.BackPack[itemName]
// 	if !has {
// 		return
// 	}
// 	selectedItem.UseItemOn(target)
// }

func (c Char) Move(deltaX, deltaY int) {
	if c.Agility < 5 {
		deltaX = min(deltaX, 1)
		deltaY = min(deltaY, 1)
	}
	c.Location.XAxis += deltaX
	c.Location.YAxis += deltaY
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
