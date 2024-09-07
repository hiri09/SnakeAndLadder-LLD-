package main

import (
	"fmt"
	"snake-and-ladder/service"
)

func main() {
	players := map[int]string{
		1: "Alice",
		2: "Bob",
		3: "joy",
		4: "raman",
	}

	// Dummy data for ladders
	ladders := map[int]int{
		3:  22, // Ladder from position 3 to 22
		5:  8,  // Ladder from position 5 to 8
		11: 26, // Ladder from position 11 to 26
	}

	// Dummy data for snakes
	snakes := map[int]int{
		17: 4, // Snake from position 17 to 4
		19: 7, // Snake from position 19 to 7
		21: 9, // Snake from position 21 to 9
	}
	game := service.CreateNewGame(100, players, ladders, snakes, 1)
	fmt.Printf("Board Size: %d\n", game.BoardSize)
	fmt.Printf("Players: %d\n", game.Players.NoOfPlayers())
	// fmt.Printf("Ladders: %v\n", game.Ladders)
	// fmt.Printf("Snakes: %v\n", game.Snakes)
	fmt.Printf("Player Positions: %+v\n", game.PlayerPos)

	game.StartGame() // Call StartGame as a method of Game
}
