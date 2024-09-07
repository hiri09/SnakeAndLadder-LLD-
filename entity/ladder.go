package entity

import "fmt"

type Ladder struct {
	StartPos int
	EndPos   int
}

func IsValidLadder(start_idx int, end_idx int) bool {
	return start_idx < end_idx
}

func (ladder *Ladder) HandleObjectInteraction(currentPos int, name string) int {
	if currentPos == ladder.StartPos {
		fmt.Printf("Player %s found a ladder at position %d, moving to position %d\n", name, ladder.StartPos, ladder.EndPos)
		return ladder.EndPos
	}
	return currentPos
}
