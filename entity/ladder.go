package entity

type Ladder struct {
	StartPos int
	EndPos   int
}

func IsValidLadder(start_idx int, end_idx int) bool {
	return start_idx < end_idx
}
