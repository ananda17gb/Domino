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
	var i int = 0
	fmt.Println("Insert name")
	fmt.Scan(&P[i].name)
	fmt.Println("Hello", P[i].name)
	i++
	fmt.Println()
	fmt.Println("Please select: ")
	fmt.Println("1. Start")
	fmt.Println("2. How to play")
	fmt.Println("3. Leaderboard")
	fmt.Println("4. Exit")
	fmt.Scan(&n)
	if n == 1 || n == 2 || n == 3 || n == 4 {
		if n == 1 {
			play(&T)
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
	var T arrTile
	var P arrPlayer
	var answer string
	fmt.Println("These are the number that you can input to make a move")
	fmt.Println("1 is for changing tile 1")
	fmt.Println("2 is for changing tile 2")
	fmt.Println("3 is for changing tile 3")
	fmt.Println("4 is for changing tile 4")
	fmt.Println("0 is for done")
	fmt.Println("9 is for exiting current game")
	fmt.Println("Go to the main menu? {Yes/No}")
	fmt.Scan(&answer)
	if answer == "Yes" {
		mainmenu(T, P)
	}
}

func play(T *arrTile) {
	var player, comp arrTile
	var decision int
	var dec1, dec2, dec3, dec4 bool = false, false, false, false
	var score, round int = 0,0
	fmt.Println("Dealing ...")
	for decision != 0 {
		pipsGenerator(&*T)
		player = T
		pipsGenerator(&*T)
		comp = T
		fmt.Print("Your tiles: ", player)
		fmt.Println("Decision? ")
		fmt.Scan(&decision)
		if decision == 1 && dec1 {
			changeTile(&*T, 1)
			dec1 = true
		} else if decision == 2 && dec2 {
			changeTile(&*T, 2)
			dec2 = true
		} else if decision == 3 && dec3 {
			changeTile(&*T, 3)
			dec3 = true
		} else if decision == 4 && dec4 {
			changeTile(&*T, 4)
			dec4 = true
		} else if decision == 9 {
			fmt.Println("Your last score is",)
			fmt.Println("Thank you for playing with us.")
			fmt.Print("Your winning rate is ",,"%")
			fmt.Println()
		} else {
			decision = 0
		}
	}
}

func changeTile(T *arrTile, n int) {

}

func leaderboard() {

}

func main() {
	var T arrTile
	var P arrPlayer
	header()
	mainmenu(T, P)
}
