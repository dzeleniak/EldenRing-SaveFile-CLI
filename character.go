package main

import (
	"fmt"
)

var (
	pattern = []byte{
		0xB0, 0xAD, 0x01, 0x00, 0x01, 0xFF, 0xFF, 0xFF,
	}

	patternDlc = []byte{
		0xB0, 0xAD, 0x01, 0x00, 0x01,
	}

	itemData = LoadItemData()
)

type character struct {
	name       string
	data       []byte
	isDlcFile  bool
	inventory  []Item
	Armaments  []*armament
	Armor      []*armor
	AshesOfWar []*ashOfWar
	Magic      []*magic
	Talisman   []*talisman
	SpritAshes []*spiritAsh
}

func (c *character) Load() {
	c.LoadInventory()
}

func (c *character) LoadInventory() []byte {
	index := subfinder(c.data, pattern)
	if index != -1 {
		index += len(pattern) + 8
	}

	if index == -1 {
		index = subfinder(c.data, patternDlc)
		if index != -1 {
			index += len(patternDlc) + 3
			c.isDlcFile = true
		}
	}

	if index == -1 {
		// Pattern not found, handle error appropriately
		return nil
	}

	// Create a new slice of 50 zeros for comparison
	emptyPattern := make([]byte, 50)
	index1 := subfinder(c.data[index:], emptyPattern)
	if index1 != -1 {
		index1 += index + 6
	}

	if index1 == -1 {
		// Empty pattern not found, handle error appropriately
		return nil
	}

	chunkSize := 8
	if c.isDlcFile {
		chunkSize = 16
	}

	fmt.Println(c.isDlcFile)

	c.inventory = make([]Item, 0)

	c.Armaments = make([]*armament, 0)
	c.Armor = make([]*armor, 0)
	c.AshesOfWar = make([]*ashOfWar, 0)
	c.Magic = make([]*magic, 0)
	c.Talisman = make([]*talisman, 0)
	c.SpritAshes = make([]*spiritAsh, 0)

	for _, itemData := range split(c.data[index:index1], chunkSize) {
		itemId := getIdReversed(itemData)
		item, err := getItemById(itemId)
		if err != nil {
			continue
		}

		c.inventory = append(c.inventory, item)

		switch t := item.(type) {
		case *armament:
			x := item.(*armament)
			c.Armaments = append(c.Armaments, x)
		case *ashOfWar:
			x := item.(*ashOfWar)
			c.AshesOfWar = append(c.AshesOfWar, x)
		case *armor:
			x := item.(*armor)
			c.Armor = append(c.Armor, x)
		case *magic:
			x := item.(*magic)
			c.Magic = append(c.Magic, x)
		case *spiritAsh:
			x := item.(*spiritAsh)
			c.SpritAshes = append(c.SpritAshes, x)
		case *talisman:
			x := item.(*talisman)
			c.Talisman = append(c.Talisman, x)
		default:
			panic(fmt.Sprintf("unexpected main.Item: %#v", t))
		}
	}

	return c.data[index:index1]
}
