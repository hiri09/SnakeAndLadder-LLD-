package utils

import "math/rand" 

func CalculateRandomNo (noOfDice int) int {
	return rand.Intn(6*noOfDice) + 1 ;
}