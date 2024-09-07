package service

import (
	"fmt" // Add this import for fmt.Println
	"snake-and-ladder/entity"
	"snake-and-ladder/utils"
)

type Game struct {
	BoardSize int
	Players   utils.Queue
	Ladders   []*entity.Ladder
	Snakes    []*entity.Snake
	PlayerPos map[int]int
	NoOfDice  int
}

func putPlayerInQueue(currentPlayers *utils.Queue, players map[int]string) map[int]int {
	playersPos := make(map[int]int)
	for key, value := range players {
		player := &entity.Player{
			Name: value,
			Id:   key,
		}
		playersPos[key] = 0
		currentPlayers.AddPlayerInQueue(player)
	}

	return playersPos
}

func CreateLadders(positions map[int]int) []*entity.Ladder {
	var ladders []*entity.Ladder
	for start, end := range positions {
		if entity.IsValidLadder(start, end) {
			ladder := &entity.Ladder{
				StartPos: start,
				EndPos:   end,
			}
			ladders = append(ladders, ladder)
		}
	}
	return ladders
}

func CreateSnakes(positions map[int]int) []*entity.Snake {
	var snakes []*entity.Snake
	for start, end := range positions {
		if entity.IsValidSnake(start, end) {
			snake := &entity.Snake{
				StartPos: start,
				EndPos:   end,
			}
			snakes = append(snakes, snake)
		}
	}
	return snakes
}

func CreateNewGame(boardSize int, players map[int]string, ladderPositions map[int]int, snakePositions map[int]int, noOfDice int) *Game {
	var currentPlayers utils.Queue
	playersPos := putPlayerInQueue(&currentPlayers, players)

	ladders := CreateLadders(ladderPositions)
	snakes := CreateSnakes(snakePositions)

	return &Game{
		BoardSize: boardSize,
		Players:   currentPlayers,
		Ladders:   ladders,
		Snakes:    snakes,
		PlayerPos: playersPos,
		NoOfDice:  noOfDice,
	}
}

func (game *Game) StartGame() {

	for game.Players.NoOfPlayers() > 1 {
		currentPlayer := game.Players.RemovePlayerFromQueue()

		diceRollValue := utils.CalculateRandomNo(game.NoOfDice)

		currntPlayerPos := game.PlayerPos[currentPlayer.Id]
		nextPos := diceRollValue + currntPlayerPos
		fmt.Printf("Player %s Rolls dice value is %d\n", currentPlayer.Name, nextPos)
		if nextPos > game.BoardSize {
			game.Players.AddPlayerInQueue(currentPlayer)
			continue
		} else {
			//if we find ladder
			for _, ladder := range game.Ladders {
				start := ladder.StartPos
				end := ladder.EndPos
				if nextPos == start {
					fmt.Printf("Player %s found a ladder at position %d, moving to position %d\n", currentPlayer.Name, start, end)
					nextPos = end
					break
				}
			}
			//if we find a snake
			for _, snake := range game.Snakes {
				start := snake.StartPos
				end := snake.EndPos
				if nextPos == start {
					fmt.Printf("Player %s found a Snake at position %d, moving to position %d\n", currentPlayer.Name, start, end)
					nextPos = end
					break
				}
			}

			if nextPos == game.BoardSize {
				fmt.Printf("Player %s won!\n", currentPlayer.Name)
			} else {
				game.Players.AddPlayerInQueue(currentPlayer)
				game.PlayerPos[currentPlayer.Id] = nextPos
			}
		}
	}
	fmt.Println("Game Over!")
}
