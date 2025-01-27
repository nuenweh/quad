package piscine

import "fmt"


// Main function, it is being called by its varints, each giving their own symbols
func QuadCustom(x, y int, TopLeft, TopRight, BotLeft, BotRight, Hor, Ver rune) {
	// Returns nothing, if any of the function arguments are negative or equals to 0
	if x <= 0 || y <= 0 {
		return
	}

	// Prints top side, or just the corner
	if x == 1 {
		fmt.Println(string(TopLeft))
	} else {
		fmt.Print(string(TopLeft))
		for i := 0; i < x-2; i++ {
			fmt.Print(string(Hor))
		}
		fmt.Println(string(TopRight))
	}

	// Prints left and right side, or only left side
	for i := 0; i < y-2; i++ {
		if x == 1 {
			fmt.Println(string(Ver))
		} else {
			fmt.Print(string(Ver))
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Println(string(Ver))
		}
	}

	// Prints bottom side, or just the corner
	if y > 1 {
		if x == 1 {
			fmt.Println(string(BotLeft))
		} else {
			fmt.Print(string(BotLeft))
			for i := 0; i < x-2; i++ {
				fmt.Print(string(Hor))
			}
			fmt.Println(string(BotRight))
		}
	}
}

// Variants of this function, each uses different symbols
func QuadA(x, y int) {
	QuadCustom(x, y, 'o', 'o', 'o', 'o', '-', '|')
}

func QuadB(x, y int) {
	QuadCustom(x, y, '/', '\\', '\\', '/', '*', '*')
}

func QuadC(x, y int) {
	QuadCustom(x, y, 'A', 'A', 'C', 'C', 'B', 'B')
}

func QuadD(x, y int) {
	QuadCustom(x, y, 'A', 'C', 'A', 'C', 'B', 'B')
}

func QuadE(x, y int) {
	QuadCustom(x, y, 'A', 'C', 'C', 'A', 'B', 'B')
}
