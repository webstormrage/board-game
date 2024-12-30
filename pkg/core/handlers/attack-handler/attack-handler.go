package attackHandler

import (
	"board-game/pkg/unit"
	"board-game/pkg/dice"
)

func Handle(initiator *unit.Squad, target *unit.Squad)int{
	dices := initiator.UnitType.MeeleDices
	damage := 0
	for i := 0; i < dices; i++ {
		roll := dice.Roll("attack")
		if roll >= target.UnitType.DefenceClass {
			damage++
		}
	}
	return damage
}