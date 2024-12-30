package spiritSaveHandler

import (
	"board-game/pkg/unit"
	"board-game/pkg/dice"
)

func Handle(target *unit.Squad, damage int)int{
	if damage < 1 {
		return 0
	}
	for i := 0; i < (target.Strength + 1 - damage)/2; i++ {
		roll := dice.Roll("spirit save")
		if roll > 4 {
			return 0
		}
	}
	return 1
}