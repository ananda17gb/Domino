# Solitaire Game: Western Double-Six Domino
Group project with I Nyoman Manutama Surya Jagadhita for Algorithm Programming course in Telkom University

---------------------------------------------------------------------------------------------------------------

a. The program design must be modulars, i.e. the use of subprograms must be employed.

b. Each group can design how to break down the program into several subprograms. Please make sure that each subprogram is complete with parameters and specifications.

c. Use arrays and structures when appropriate. Define your own types suitable for the problem.

d. The program must implement at least the sequential and binary search algorithms, deleting and modifying data stored in the array and/or sturcture.

e. The program must implement Selection Sort and Insertion Sort algorithms, each for different data or field. They can be used to sort ascending or descending order.

f. The program must be structured, only one entry and one exit points. e.g. Must not have return statement, break statement, or continue statement in the middle of any loop.

g. Global definition is forbidden for variabel declaration. You can use global definition only for newly defined data types and/or constant definition.


---------------------------------------------------------------------------------------------------------------

Description: Western double-six dominoes have 28 tiles as illustrated below. Each tile is divided into two sections, each containing 0, 1, 2, upto 6 dots. Double tiles are tile that have the same number of dots on both sections. The value of a tile is the sum of dots on the tile; 0..12.

Ceme-4tile game can be played as follows: The domino set is scrambled with all tiles down (forming a boneyard). A player should take and open (turn upside) four tiles from the boneyard. The opponent also takes four tiles, but keeps the tiles face down. A player can change his/her tiles at most twice by taking new tiles from the boneyard and throwing tiles he/she doesn’t want to keep. 

Winner will be determine as follows:
a. Having two double tiles and the opponent has at most only one double tile.

b. If the opponent also have double tiles, the winner is the one with higher double tile values.

c. Otherwise, the one with higher sum of tile values is the winner.

Specification:
a. The app should be able to scramble the 28 tiles (not the number on the tiles). Use random() function from package math/rand.

b. The app should be able to pick tiles and change the tile.

c. The player should be able to restart the game, get four tiles (and provide 4 tiles to the opponent), select and replace tiles, and finish to compute the winner.

d. The app should be able to compute score, keep track how many games have been played, update score, and determine the winner.

e. The app can be stopped and show the game statistics.

f. To start a game, the player must first enter his/her initial/name.

g. Highest score for player with same name will be recorded and can be displayed.

h. The app can display player list, ranked by highest scores.

---------------------------------------------------------------------------------------------------------------
