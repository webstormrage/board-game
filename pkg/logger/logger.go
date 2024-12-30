package logger

import (
	"board-game/pkg/unit"
	"fmt"
)

var isEnabled = false

func Enable(){
	isEnabled = true
}

func Disable(){
	isEnabled = false
}

func LogAttack(initiator *unit.Squad, target *unit.Squad, damage int, moralDebugg int){
	if !isEnabled {
		return
	}
	fmt.Printf("%d %s [%d] attack %d %s [%d] - ",
		initiator.Strength,
		 initiator.UnitType.Name,
		 initiator.Spirit,
		 target.Strength,
		 target.UnitType.Name,
		 target.Spirit,
	)
	fmt.Printf("%d damage ", damage)
	fmt.Printf("(%d moral debuff)\n", moralDebugg)
}

func LogRoll(action string, value int){
	if !isEnabled {
		return
	}
	fmt.Printf("[%s] roll %d\n", action, value)
}

func LogEvent(event string){
	if !isEnabled {
		return
	}
	fmt.Printf("\n%s\n", event)
}

func LogRound(round int){
	fmt.Printf("\nROUND %d\n", round)
}