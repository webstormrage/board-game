package main

import (
	core "board-game/pkg/core"
	EventTypes "board-game/pkg/event-types"
	"board-game/pkg/logger"
	unit "board-game/pkg/unit"
	"container/list"
	"slices"
)

func findFirstTarget(current *unit.Squad, squads *list.List)*unit.Squad{
	e := squads.Front()
	if e == nil {
		return nil
	}

	return e.Value.(*unit.Squad)
}

func chooseAction(actions []int)int{
	if(slices.Contains(actions, EventTypes.ACTION_MELLEE_ATTACK)){
		return EventTypes.ACTION_MELLEE_ATTACK
	}
	if(slices.Contains(actions, EventTypes.REACTION_MELLEE_COUNTER_ATTACK)){
		return EventTypes.REACTION_MELLEE_COUNTER_ATTACK
	}
	if(slices.Contains(actions, EventTypes.ACTION_SKIP)){
		return EventTypes.ACTION_SKIP
	}
	if(slices.Contains(actions, EventTypes.REACTION_SKIP)){
		return EventTypes.REACTION_SKIP
	}
	return -1
}

func chooseSkip(actions []int)int {
	if(slices.Contains(actions, EventTypes.ACTION_SKIP)){
		return EventTypes.ACTION_SKIP
	}
	if(slices.Contains(actions, EventTypes.REACTION_SKIP)){
		return EventTypes.REACTION_SKIP
	}
	return -1
}

func countDead(squads *list.List)int{
	dead := 0
	for e := squads.Front(); e != nil; e = e.Next() {
        itemSquad := e.Value.(*unit.Squad)
        if itemSquad.Strength <= 0 || itemSquad.Spirit <= 0 {
			dead++
		}
    }
	return dead
}

func main() {
	core.QueueSquad(&unit.Squad{
		UnitType: unit.Militia,
		Strength: 4,
		Spirit: 2,
	})
	core.QueueSquad(&unit.Squad{
		UnitType: unit.Raider,
		Strength: 2,
		Spirit: 2,
	})
	logger.Enable()
	core.Initialize()
	for true {
		current, squads, actions := core.GetTurn()
		if(countDead(squads) > 0){
			return
		}
		action := chooseAction(actions)
	    target := findFirstTarget(current, core.GetAvailableTargets(action))
		if target == nil {
			actions := chooseSkip(actions)
			core.Act(nil, nil, actions)
		} else {
			core.Act(current, target, action)
		}
	}
	
}