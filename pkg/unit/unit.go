package unit

import (
	EventTypes "board-game/pkg/event-types"
)

type Unit struct {
	Name string
	Actions []int
	MeeleDices int
	DefenceClass int
	Speed int
	Initiative int
	Morality int
	Tier int
	Reactions []int
}


type Squad struct {
	UnitType Unit
	Strength int
	Spirit int
	HasUsedReaction bool
	InBattleWith *Squad
}

var Militia = Unit{
	Name: "Militia",
	Actions: []int{
		EventTypes.ACTION_MELLEE_ATTACK,
	},
	MeeleDices: 1,
	DefenceClass: 3,
	Reactions: []int{
		// EventTypes.REACTION_MELLEE_COUNTER_ATTACK,
	},
	Speed: 1,
	Initiative: 2,
	Morality: 2,
	Tier: 1,
}

var Raider = Unit{
	Name: "Raider",
	Actions: []int{
		EventTypes.ACTION_MELLEE_ATTACK,
	},
	MeeleDices: 2,
	DefenceClass: 3,
	Reactions: []int{
		// EventTypes.REACTION_MELLEE_COUNTER_ATTACK,
	},
	Speed: 3,
	Initiative: 5,
	Morality: 2,
	Tier: 2,
}
