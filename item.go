package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"strings"
)

type Item interface {
	ItemId() string
	ItemType() string
	ItemName() string
}

type ashOfWar struct {
	Id       string
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (i *ashOfWar) ItemId() string {
	return i.Id
}

func (i *ashOfWar) ItemType() string {
	return "Ash Of War"
}

func (i *ashOfWar) ItemName() string {
	return i.Name
}

type armor struct {
	Id       string
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (i *armor) ItemId() string {
	return i.Id
}

func (i *armor) ItemType() string {
	return "Armor"
}

func (i *armor) ItemName() string {
	return i.Name
}

type armament struct {
	Id    string
	Name  string `json:"name"`
	Class string `json:"class"`
}

func (i *armament) ItemId() string {
	return i.Id
}

func (i *armament) ItemType() string {
	return "Armament"
}

func (i *armament) ItemName() string {
	return i.Name
}

type magic struct {
	Id       string
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (i *magic) ItemId() string {
	return i.Id
}

func (i *magic) ItemType() string {
	return "Magic"
}

func (i *magic) ItemName() string {
	return i.Name
}

type spiritAsh struct {
	Id   string
	Name string `json:"name"`
}

func (i *spiritAsh) ItemId() string {
	return i.Id
}

func (i *spiritAsh) ItemType() string {
	return "Spirit Ash"
}

func (i *spiritAsh) ItemName() string {
	return i.Name
}

type talisman struct {
	Id   string
	Name string `json:"name"`
}

func (i *talisman) ItemId() string {
	return i.Id
}

func (i *talisman) ItemType() string {
	return "Talisman"
}

func (i *talisman) ItemName() string {
	return i.Name
}

type ItemDataFile struct {
	Armament    map[string]armament  `json:"armament"`
	Armor       map[string]armor     `json:"armor"`
	AshesOfWar  map[string]ashOfWar  `json:"ashesOfWar"`
	Magic       map[string]magic     `json:"magic"`
	SpiritAshes map[string]spiritAsh `json:"spiritAshes"`
	Talisman    map[string]talisman  `json:"talisman"`
}

func LoadItemData() *ItemDataFile {
	itemData := loadBaseitems()
	loadDlcItems(itemData)

	return itemData
}

func loadBaseitems() *ItemDataFile {
	file, err := os.Open("data/baseItems.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var itemData ItemDataFile
	err = json.NewDecoder(file).Decode(&itemData)

	if err != nil {
		panic(err)
	}

	return &itemData
}

func loadDlcItems(itemData *ItemDataFile) {
	file, err := os.Open("data/dlcItems.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var dlcItemData ItemDataFile
	err = json.NewDecoder(file).Decode(&dlcItemData)

	if err != nil {
		panic(err)
	}

	maps.Copy(itemData.Armament, dlcItemData.Armament)
	maps.Copy(itemData.Armor, dlcItemData.Armor)
	maps.Copy(itemData.AshesOfWar, dlcItemData.AshesOfWar)
	maps.Copy(itemData.Magic, dlcItemData.Magic)
	maps.Copy(itemData.SpiritAshes, dlcItemData.SpiritAshes)
	maps.Copy(itemData.Talisman, dlcItemData.Talisman)
}

func getItemById(id string) (Item, error) {
	// id = strings.TrimLeft(id, "0")
	id = strings.ToUpper(id)

	// fmt.Println(id)

	arm, ok := itemData.Armament[id]
	if ok {
		return &armament{
			Id:    id,
			Name:  arm.Name,
			Class: arm.Class,
		}, nil
	}

	armr, ok := itemData.Armor[id]
	if ok {
		return &armor{
			Id:       id,
			Category: armr.Category,
			Name:     armr.Name,
		}, nil
	}

	aow, ok := itemData.AshesOfWar[id]
	if ok {
		return &ashOfWar{
			Id:       id,
			Category: aow.Category,
			Name:     aow.Name,
		}, nil
	}

	ash, ok := itemData.SpiritAshes[id]
	if ok {
		return &spiritAsh{
			Id:   id,
			Name: ash.Name,
		}, nil
	}

	tal, ok := itemData.Talisman[id]
	if ok {
		return &talisman{
			Id:   id,
			Name: tal.Name,
		}, nil
	}

	return nil, fmt.Errorf("could not find item")
}
