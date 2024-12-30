package dice

import (
	"math/rand"
	"time"
	logger "board-game/pkg/logger"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Roll(action string)int{
	result := r.Intn(6) + 1
	logger.LogRoll(action, result)
	return result
}