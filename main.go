package main

import (
	"fmt"
	"strings"
)

func main() {

	// file, err := os.ReadFile("ER0000.sl2")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// names, err := getCharacterNames(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for i, n := range names {
	// 	fmt.Printf("%d -> %v\n", i+1, n)
	// }

	save := &saveFile{
		filename: "ER0000.sl2",
	}

	save.Load()
	save.LoadCharacters()

	// fmt.Println(len(save.data))

	for _, c := range save.characters {
		if strings.Trim(c.name, "") != "" {
			fmt.Printf("Name: %v\n", c.name)
			fmt.Printf("Armaments:\n")
			for _, item := range c.Armaments {
				fmt.Printf("\t- %v\n", item.ItemName())
			}
			fmt.Printf("Talisman:\n")
			for _, item := range c.Talisman {
				fmt.Printf("\t- %v\n", item.ItemName())
			}
		}
	}
}
