package entity

import "fmt"

type Snake struct {
	StartPos int
	EndPos   int
}

func IsValidSnake(start_idx int, end_idx int) bool {
	return start_idx > end_idx
}

func (snake *Snake) HandleObjectInteraction(currentPos int, name string) int {
	if currentPos == snake.StartPos {
		fmt.Printf("Player %s found a ladder at position %d, moving to position %d\n", name, snake.StartPos, snake.EndPos)
		return snake.EndPos
	}
	return currentPos
}
