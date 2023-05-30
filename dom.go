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
	fmt.Println()
}

func pipsGenerator(T *arrTile, n int) {
	T[n].d1 = rand.Intn(13)
	T[n].d2 = rand.Intn(13)
}

func mainmenu(T arrTile, P arrPlayer) {
	var n int
	var i int = 0
	fmt.Print("Please insert your name: ")
	fmt.Scan(&P[i].name)
	fmt.Println()
	fmt.Println("Hello", P[i].name)
	i++
	fmt.Println()
	fmt.Println("Please select: ")
	fmt.Println("1. Start")
	fmt.Println("2. Leaderboard")
	fmt.Println("3. Exit")
	fmt.Scan(&n)
	if n == 1 || n == 2 || n == 3 {
		if n == 1 {
			fmt.Println()
			play(&T)
		} else if n == 2 {
			fmt.Println()
			// leaderboard()
		} else if n == 3 {
			fmt.Println("Thanks for playing :)")
		} else {
			fmt.Println("Please input any number that are on the screen")
			fmt.Scan(&n)
		}
	}

}

func howtoplay() {
	fmt.Println("These are the number that you can input to make a move")
	fmt.Println("1 is for changing tile 1")
	fmt.Println("2 is for changing tile 2")
	fmt.Println("3 is for changing tile 3")
	fmt.Println("4 is for changing tile 4")
	fmt.Println("0 is for done")
	fmt.Println("9 is for exiting current game")
}

func play(T *arrTile) {
	var player, comp arrTile
	var decision int = -1
	var dec1, dec2, dec3, dec4 bool = false, false, false, false
	var score, round int = 0, 0
	howtoplay()
	fmt.Println()
	fmt.Println("Dealing ...")
	fmt.Println("Your score is", score, "/", round)
	for decision != 0 {
		for i := 0; i < 4; i++ {
			pipsGenerator(&*T, i)
		}
		player = *T
		for i := 0; i < 4; i++ {
			pipsGenerator(&*T, i)
		}
		comp = *T
		fmt.Print("Your tiles: ", player)
		fmt.Println()
		fmt.Println("Decision? ")
		fmt.Scan(&decision)
		if decision == 1 && dec1 {
			pipsGenerator(&player, 1)
			dec1 = true
		} else if decision == 2 && dec2 {
			pipsGenerator(&player, 2)
			dec2 = true
		} else if decision == 3 && dec3 {
			pipsGenerator(&player, 3)
			dec3 = true
		} else if decision == 4 && dec4 {
			pipsGenerator(&player, 4)
			dec4 = true
		} else if decision == 9 {
			fmt.Println("Your last score is", score, "/", round)
			fmt.Println("Thank you for playing with us.")
			fmt.Print("Your winning rate is ", score/round, "%")
			fmt.Println()
		} else if decision == 0 {
			ppoints := 0
			cpoints := 0
			for i := 0; i < 4; i++ {
				ppoints = player[i].d1 + player[i].d2
			}
			for i := 0; i < 4; i++ {
				cpoints = comp[i].d1 + comp[i].d2
			}
			if ppoints > cpoints {
				fmt.Println("You won")
				score++
				round++
				fmt.Println("Your score is", score, "/", round)
			} else if ppoints < cpoints {
				fmt.Println("You lost")
				round++
				fmt.Println("Your score is", score, "/", round)
			} else {
				fmt.Println("It's a draw, current game score will not be submitted to the total score")
			}
		} else {
			fmt.Println("You already input that one, please input other number displayed")
		}
	}
}

// func changeTile(T *arrTile, n int) {

// }

// func leaderboard() {

// }

func main() {
	var T arrTile
	var P arrPlayer
	header()
	mainmenu(T, P)
}
