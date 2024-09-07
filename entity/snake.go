package entity

type Snake struct {
	StartPos int
	EndPos   int
}

func IsValidSnake(start_idx int, end_idx int) bool {
	return start_idx > end_idx
}
