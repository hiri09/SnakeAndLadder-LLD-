package service

import (
	// Add this import for fmt.Println

	"fmt"
	"snake-and-ladder/entity"
	"snake-and-ladder/utils"
)

type BoardObjects interface {
	HandleObjectInteraction(currentPos int, name string) int
}

type Game struct {
	BoardSize        int
	Players          utils.Queue
	GameBoardObjects []BoardObjects
	PlayerPos        map[int]int
	NoOfDice         int
}

func (game *Game) putPlayerInQueue(players map[int]string) {

	for id, name := range players {
		player := &entity.Player{
			Name: name,
			Id:   id,
		}

		game.Players.AddPlayerInQueue(player)
		game.PlayerPos[id] = 0
	}
}

func getBoardObject(start int, end int) BoardObjects {
	if entity.IsValidLadder(start, end) {
		return &entity.Ladder{
			StartPos: start,
			EndPos:   end,
		}
	}

	return &entity.Snake{
		StartPos: start,
		EndPos:   end,
	}
}

func (game *Game) FillBoardObject(ladderPositions map[int]int, snakePositions map[int]int) {

	for start, end := range ladderPositions {
		gameBoardObject := getBoardObject(start, end)
		game.GameBoardObjects = append(game.GameBoardObjects, gameBoardObject)
	}

	for start, end := range snakePositions {
		gameBoardObject := getBoardObject(start, end)
		game.GameBoardObjects = append(game.GameBoardObjects, gameBoardObject)
	}
}

func CreateNewGame(boardSize int, players map[int]string, ladderPositions map[int]int, snakePositions map[int]int, noOfDice int) *Game {
	game := &Game{
		BoardSize: boardSize,
		PlayerPos: make(map[int]int),
		NoOfDice:  noOfDice,
	}

	game.putPlayerInQueue(players)

	game.FillBoardObject(ladderPositions, snakePositions)

	return game
}

func (game *Game) StartGame() {
	totalPlayer := game.Players.NoOfPlayers()
	for game.Players.NoOfPlayers() > 1 {
		currentPlayer := game.Players.RemovePlayerFromQueue()

		diceRollValue := utils.CalculateRandomNo(game.NoOfDice)

		currentPlayerPos := game.PlayerPos[currentPlayer.Id]
		nextPosition := diceRollValue + currentPlayerPos

		fmt.Printf("Player %s Rolls dice value is %d\n", currentPlayer.Name, nextPosition)
		if nextPosition > game.BoardSize {
			game.Players.AddPlayerInQueue(currentPlayer)
			continue
		} else {

			for _, object := range game.GameBoardObjects {
				nextPosition = object.HandleObjectInteraction(nextPosition, currentPlayer.Name)
			}

			if nextPosition == game.BoardSize {
				fmt.Printf("Player %s won comes %d!\n", currentPlayer.Name, totalPlayer-game.Players.NoOfPlayers())
			} else {
				game.Players.AddPlayerInQueue(currentPlayer)
				game.PlayerPos[currentPlayer.Id] = nextPosition
			}
		}
	}
	fmt.Println("Game Over!")
}
