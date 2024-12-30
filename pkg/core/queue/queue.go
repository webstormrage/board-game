package queue

import (
	EventTypes "board-game/pkg/event-types"
	"board-game/pkg/logger"
	unit "board-game/pkg/unit"
	"container/list"
	"slices"
)

var round = 0
var squadOrder = list.New()
var currentActionSquad = squadOrder.Front()
var currentReactionSquad *unit.Squad = nil
var currentReactions []int = []int{}


func QueueSquad(squad *unit.Squad){
	for e := squadOrder.Front(); e != nil; e = e.Next() {
        itemSquad := e.Value.(*unit.Squad)
        if itemSquad.UnitType.Initiative < squad.UnitType.Initiative {
			squadOrder.InsertBefore(squad, e)
			return
		}
    }
	squadOrder.PushBack(squad)
}

func Initialize(){
	currentActionSquad = squadOrder.Front()
	for e := squadOrder.Front(); e != nil; e = e.Next() {
		itemSquad := e.Value.(*unit.Squad)
		itemSquad.HasUsedReaction = false
	}
	round += 1
	logger.LogRound(round)
}

func GetCurrent()*unit.Squad{
	return currentActionSquad.Value.(*unit.Squad)
}

func GetTurn()(current *unit.Squad, squads *list.List, moves []int){
	if currentReactionSquad != nil {
		return currentReactionSquad, squadOrder, currentReactions
	}
	squad := currentActionSquad.Value.(*unit.Squad)
	// Берем доступные действия юнита
	actions := squad.UnitType.Actions
	// Добавляем пустое действие
	actions = append(actions, EventTypes.ACTION_SKIP)
	return squad, squadOrder, slices.DeleteFunc(actions, func(a int) bool {
		// Если юнит окружен то он не может использовать дальнобойные атаки
		if squad.InBattleWith != nil && a == EventTypes.ACITON_RANGED_ATTACK {
			return true
		}
		return false
	})
}

func Intercept(current *unit.Squad, reactions []int) {
	if current.HasUsedReaction  {
		return
	}
	currentReactionSquad = current
	currentReactions = slices.DeleteFunc(reactions, func(r int) bool {
		return slices.Contains(current.UnitType.Reactions, r)
	})
}

func ResolveInterception(fullfiled bool){
	if fullfiled {
		currentReactionSquad.HasUsedReaction = true
	}
	currentReactionSquad = nil
	currentReactions = []int{}
}

func NexTurn(){
	if currentActionSquad.Next() == nil {
		Initialize()
	} else {
		currentActionSquad = currentActionSquad.Next()
		squad := currentActionSquad.Value.(*unit.Squad)
		if squad.Strength < 1 || squad.Spirit < 1 {
			NexTurn()
		}
	}
}