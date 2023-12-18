package main

import (
	"fmt"
	"math/rand"
	"time"
)

const nT int = 2
const nA int = 4
const pMAX int = 101

type arrTile [nA]tile
type tile [nT]int

type arrPlayer struct {
	player [pMAX]playerInfo
	n      int
}
type playerInfo struct {
	name                  string
	score, round, winrate int
}

// MAIN
func main() {
	var P arrPlayer
	mainmenu(&P)
}

// Main Menu and Header
func header() {
	fmt.Println()
	fmt.Println("Welcome to the game of Ceme-4tile")
	fmt.Println("---------------------------------")
}

func mainmenu(P *arrPlayer) {
	var n int
	game := true
	if P.n < 1 {
		header()
	}
	fmt.Println("Type 1	| START")
	fmt.Println("Type 2	| LEADERBOARD")
	fmt.Println("Type 3	| EXIT")
	for game {
		fmt.Println()
		fmt.Print("Please select: ")
		fmt.Scanln(&n)
		if n == 1 {
			fmt.Println()
			fmt.Print("Player name: ")
			fmt.Scanln(&P.player[P.n].name)
			P.n++
			play(&*P)
			game = false
		} else if n == 2 {
			fmt.Println()
			leaderboard(*P)
			game = false
		} else if n == 3 {
			fmt.Println()
			fmt.Println("Thanks for playing :)")
			game = false
		} else {
			fmt.Println("*Please insert given number*")
		}
	}
}

// GUIDE
func howtoplay(P arrPlayer) {
	fmt.Println()
	fmt.Printf("Hello %s, this is your moveset: \n", P.player[P.n-1].name)
	fmt.Println("Type 1 to change tile 1")
	fmt.Println("Type 2 to change tile 2")
	fmt.Println("Type 3 to change tile 3")
	fmt.Println("Type 4 to change tile 4")
	fmt.Println("Attempt for each tile change and reshuffle is 1 time only!")
	fmt.Println()
	fmt.Println("Type 0 if you are ready")
	fmt.Println("Type 9 to exit current game")
	fmt.Println()
}

// Random tiles generator
func generateTilePlayer(player *arrTile) {
	for i := 0; i < nA; i++ {
		player[i][0] = rand.Intn(13)
		player[i][1] = rand.Intn(13)
	}
}
func generateTileComp(comp *arrTile) {
	for i := 0; i < nA; i++ {
		comp[i][0] = rand.Intn(13)
		comp[i][1] = rand.Intn(13)
	}
}

// PLAY SECTION
func play(P *arrPlayer) {
	var player, comp arrTile
	var decision int
	var dec1, dec2, dec3, dec4, decshuff bool = false, false, false, false, false
	var score, round, winrate int = 0, 0, 0
	rand.Seed(time.Now().UnixNano())
	generateTilePlayer(&player)
	generateTileComp(&comp)
	game := true
	guide := true
	dealing_check := true
	for game && decision != 9 {
		if guide {
			howtoplay(*P)
			guide = false
		}
		if dealing_check {
			fmt.Printf("Dealing ... [Round %v]\n", round+1)
			dealing_check = false
		}
		fmt.Println("Your tiles: ", player)
		if !decshuff {
			joe := true
			shuff := ""
			for joe {
				fmt.Print("Reshuffle? (y/n): ")
				fmt.Scan(&shuff)
				if shuff == "y" {
					decision = -1
					joe = false
				} else if shuff == "n" {
					joe = false
					decshuff = true
				} else {

				}
			}
		}
		if decshuff {
			fmt.Print("Decision? ")
			fmt.Scan(&decision)
		}
		fmt.Println()
		if decision == 1 && !dec1 {
			player[0][0] = rand.Intn(13)
			player[0][1] = rand.Intn(13)
			dec1 = true
		} else if decision == 2 && !dec2 {
			player[1][0] = rand.Intn(13)
			player[1][1] = rand.Intn(13)
			dec2 = true
		} else if decision == 3 && !dec3 {
			player[2][0] = rand.Intn(13)
			player[2][1] = rand.Intn(13)
			dec3 = true
		} else if decision == 4 && !dec4 {
			player[3][0] = rand.Intn(13)
			player[3][1] = rand.Intn(13)
			dec4 = true
		} else if decision == 0 {
			ppoints := 0
			cpoints := 0
			psame := 0
			csame := 0
			for i := 0; i < 4; i++ {
				ppoints += player[i][0] + player[i][1]
			}
			for i := 0; i < 4; i++ {
				cpoints += comp[i][0] + comp[i][1]
			}
			for i := 0; i < 4; i++ {
				if player[i][0] == player[i][1] {
					psame += 1
				}
			}
			for i := 0; i < 4; i++ {
				if comp[i][0] == comp[i][1] {
					csame += 1
				}
			}
			if psame > csame {
				fmt.Println("!!! YOU WON !!!")
				fmt.Println("Your tiles	: ", player)
				fmt.Println("Opponent tiles	: ", comp)
				fmt.Printf("You have %v Double(s), while opponents have %v\n", psame, csame)
				score++
				round++
				fmt.Println()
				fmt.Println("Your score is ", score, "|", round)
				fmt.Println()
			} else if csame > psame {
				fmt.Println("You lost")
				fmt.Println("Your tiles	: ", player)
				fmt.Println("Opponent tiles	: ", comp)
				fmt.Printf("You have %v Double(s), while opponents have %v\n", psame, csame)
				round++
				fmt.Println()
				fmt.Println("Your score is ", score, "|", round)
				fmt.Println()
			} else {
				if ppoints > cpoints {
					fmt.Println("!!! YOU WON !!!")
					fmt.Println("Your tiles	: ", player)
					fmt.Println("Opponent tiles	: ", comp)
					fmt.Printf("You have %v pips count, while opponents have %v\n", ppoints, cpoints)
					score++
					round++
					fmt.Println()
					fmt.Println("Your score is ", score, "|", round)
					fmt.Println()
				} else if ppoints < cpoints {
					fmt.Println("You lost")
					fmt.Println("Your tiles	: ", player)
					fmt.Println("Opponent tiles	: ", comp)
					fmt.Printf("You have %v pips count, while opponents have %v\n", ppoints, cpoints)
					round++
					fmt.Println()
					fmt.Println("Your score is ", score, "|", round)
					fmt.Println()
				} else {
					fmt.Println("It's a draw, current game data will not be submitted to the total score/round")
					fmt.Println()
				}
			}
			generateTilePlayer(&player)
			generateTileComp(&comp)
			dec1, dec2, dec3, dec4, decshuff = false, false, false, false, false
			dealing_check = true
		} else if decision == -1 && !decshuff {
			generateTilePlayer(&player)
			generateTileComp(&comp)
			dec1, dec2, dec3, dec4, decshuff = false, false, false, false, true
		} else if decision == 9 {
		} else {
			fmt.Println("*Action is not available*")
		}
	}
	if round != 0 {
		winrate = (score * 100) / round
	}
	fmt.Printf("Your last score %v | Round %v\n", score, round)
	fmt.Printf("Your winning rate is %v%s\n", winrate, "%")
	fmt.Println()
	P.player[P.n-1].score = score
	P.player[P.n-1].round = round
	P.player[P.n-1].winrate = winrate
	mainmenu(&*P)
}

// Leaderboard
func leaderboard(P arrPlayer) {
	var name string
	var answer int
	sortLeaderboard(&P)
	for i := 0; i < P.n; i++ {
		fmt.Printf("Player: %s 	[Score: %v	| Round: %v	| Winrate:%v%s]\n", P.player[i].name, P.player[i].score, P.player[i].round, P.player[i].winrate, "%")
	}
	fmt.Println()
	fmt.Println("1. Search your name")
	fmt.Println("2. Go back to main menu")
	answer_check := true
	enter_check := true
	enter := ""
	for answer_check {
		fmt.Print("Please select: ")
		fmt.Scan(&answer)
		if answer == 1 {
			fmt.Println()
			fmt.Print("What is your name? ")
			fmt.Scanln(&name)
			displayPlayerGames(P, name)
			fmt.Println()
			for enter_check {
				fmt.Print("Press ENTER to go back")
				fmt.Scanln(&enter)
				enter_check = false
			}
			fmt.Println()
			leaderboard(P)
			answer_check = false
		} else if answer == 2 {
			fmt.Println()
			mainmenu(&P)
			answer_check = false
		} else {
			fmt.Println("*Action not available*")
		}
	}
}

// Leaderboard sort
func sortLeaderboard(P *arrPlayer) {
	for i := 0; i < P.n-1; i++ {
		mindex := i
		// for j := i + 1; j < nA; j++ {
		for j := i + 1; j < P.n; j++ {
			if P.player[j].score > P.player[mindex].score {
				mindex = j
			}

		}
		// }
		P.player[i], P.player[mindex] = P.player[mindex], P.player[i]
	}
}
func displayPlayerGames(P arrPlayer, name string) {
	var arrName arrPlayer
	arrName = P
	sortName(&arrName)
	startex := searchRange(arrName, name)
	if startex >= 0 {
		fmt.Printf("%s's Rating: ", name)
		fmt.Printf("[Score %d    | Round %d    | Win Rate %v%s]\n", P.player[startex].score, P.player[startex].round, P.player[startex].winrate, "%")
	} else {
		fmt.Println("Player not found.")
	}
}

func sortName(P *arrPlayer) {
	for i := 0; i < P.n-1; i++ {
		mindex := i
		for j := i + 1; j < P.n; j++ {
			if P.player[j].name < P.player[mindex].name {
				mindex = j
			}
		}
		P.player[i], P.player[mindex] = P.player[mindex], P.player[i]
	}
}

func searchRange(P arrPlayer, name string) int {
	l1 := 0
	r1 := P.n - 1
	startex := -1
	for l1 <= r1 && startex == -1 {
		m := (l1 + r1) / 2
		if P.player[m].name == name {
			startex = m
		} else if P.player[m].name < name {
			l1 = m + 1
		} else {
			r1 = m - 1
		}
	}
	return startex
}
