package utils

import "snake-and-ladder/entity"

type Queue struct {
	QueuePlayer []*entity.Player
}

func (q *Queue) AddPlayerInQueue(player *entity.Player) {
	q.QueuePlayer = append(q.QueuePlayer, player)
}

func (q *Queue) RemovePlayerFromQueue() *entity.Player {
	length := len(q.QueuePlayer)

	if length == 0 {
		return nil
	}

	currentPlayerTurn := q.QueuePlayer[0]
	q.QueuePlayer = q.QueuePlayer[1:]

	return currentPlayerTurn
}

func (q *Queue) NoOfPlayers() int {
	return len(q.QueuePlayer)
}
