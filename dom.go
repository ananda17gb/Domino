package main

import (
	"fmt"
	"math/rand"
)

type tileNum struct {
	d1, d2 int
}

type arrTile [4]tileNum

type player struct {
	name         string
	score, round int
}

type arrPlayer [101]player

func header() {
	fmt.Println("Welcome to the game of Ceme-4tile")
	fmt.Println("---------------------------------")
}

func pipsGenerator(T *arrTile) {
	for i := 0; i < 4; i++ {
		T[i].d1 = rand.Intn(13)
		T[i].d2 = rand.Intn(13)
	}
	for i := 0; i < 4; i++ {
		fmt.Print("(", T[i].d1, ",", T[i].d2, ") ")
	}
}

func mainmenu(T arrTile, P arrPlayer) {
	var n int
	var name string
	fmt.Println("Insert name")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	fmt.Println()
	fmt.Println("Please select: ")
	fmt.Println("1. Start")
	fmt.Println("2. How to play")
	fmt.Println("3. Leaderboard")
	fmt.Println("4. Exit")
	fmt.Scan(&n)
	if n == 1 || n == 2 || n == 3 || n == 4 {
		if n == 1 {
			play()
		} else if n == 2 {
			howtoplay()
		} else if n == 3 {
			leaderboard()
		} else {
			fmt.Println("Thanks for playing :)")
		}
	} else {
		fmt.Println("Please input any number that are on the screen")
		fmt.Scan(&n)
	}
}

func howtoplay() {
	fmt.Println("These are the number that you can input to make a move")
	fmt.Println("1 is for changing tile")
	fmt.Println("2 is for tiling")
	fmt.Println("0 is for done")
	fmt.Println("9 is for exiting current game")
}

func play(T *arrTile) {
	var decision int
	fmt.Println("Dealing ...")
	for decision != 0 {
		pipsGenerator(&*T)
		fmt.Print("Your tiles: ", T)
		fmt.Println("Decision? ")
		fmt.Scan(&decision)
		if decision == 1 
	}

}

func main() {
	var T arrTile
	var P arrPlayer
	header()
	mainmenu(T, P)
}
