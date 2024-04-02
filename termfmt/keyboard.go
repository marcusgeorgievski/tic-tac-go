package termfmt

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Extends functionality with keyboard input handling and the detection of
//  certain key presses

const (
	UP    keyboard.Key = 65517
	DOWN  keyboard.Key = 65516
	LEFT  keyboard.Key = 65515
	RIGHT keyboard.Key = 65514
	ENTER keyboard.Key = 13
	QUIT  keyboard.Key = 113
)

func GetKey() keyboard.Key {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	char, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("You pressed: rune %q, key %v\r\n", char, key)
	_ = keyboard.Close()

	if char == 113 {
		return keyboard.Key(char)
	}
	return key
}

func GetPlayerInfo() (name, symbol string) {
	fmt.Printf("Name: ")
	fmt.Scan(&name)

	fmt.Printf("Symbol: ")
	fmt.Scan(&symbol)

	fmt.Printf("\n")
	return
}
