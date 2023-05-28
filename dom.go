package main

import "fmt"

type tileNum struct {
	d1, d2 int
}
type arrTile [4]tileNum

func header() {
	fmt.Println("Welcome to the game of Ceme-4tile")
	fmt.Println("---------------------------------")
}
func main() {
	header()
}
