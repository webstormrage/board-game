package core

import (
	attackHandler "board-game/pkg/core/handlers/attack-handler"
	spiritSaveHandler "board-game/pkg/core/handlers/spirit-save-handler"
	queue "board-game/pkg/core/queue"
	"board-game/pkg/dice"
	EventTypes "board-game/pkg/event-types"
	"board-game/pkg/logger"
	unit "board-game/pkg/unit"
	"container/list"
)

func handleSpirit(target *unit.Squad, damage int) int {
	for i := 0; i < (target.Strength+1-damage)/2; i++ {
		roll := dice.Roll("spirit save")
		if roll > 4 {
			return 0
		}
	}
	return 1
}

func QueueSquad(squad *unit.Squad) {
	queue.QueueSquad(squad)
}

func Initialize() {
	queue.Initialize()
}

func GetTurn() (current *unit.Squad, squads *list.List, moves []int) {
	current, squads, moves = queue.GetTurn()
	return current, squads, moves
}

func GetAvailableTargets(eventType int) *list.List {
	targets := list.New()
	current, squads, _ := queue.GetTurn()
	if eventType == EventTypes.REACTION_MELLEE_COUNTER_ATTACK {
		targets.PushBack(queue.GetCurrent())
		return targets
	}
	if eventType == EventTypes.REACTION_SKIP {
		return targets
	}
	if eventType == EventTypes.ACTION_SKIP {
		return targets
	}
	if current.InBattleWith != nil {
		targets.PushBack(current.InBattleWith)
		return targets
	}
	for e := squads.Front(); e != nil; e = e.Next() {
		itemSquad := e.Value.(*unit.Squad)
		// TODO: добавить разграничение СВОЙ/ЧУЖОЙ
		if itemSquad != current && itemSquad.UnitType.Speed <= current.UnitType.Speed {
			targets.PushBack(itemSquad)
		}
	}
	return targets
}

func Act(initiator *unit.Squad, target *unit.Squad, eventType int) {
	switch eventType {
	case EventTypes.ACTION_MELLEE_ATTACK:
		logger.LogEvent("Mellee attack")
		damage := attackHandler.Handle(initiator, target)
		debuff := spiritSaveHandler.Handle(target, damage)

		logger.LogAttack(initiator, target, damage, debuff)

		target.Strength -= damage
		target.Spirit -= debuff

		// Если инициатор кого-то окружил в прошлом ходу, он больше его не окружает
		if initiator.InBattleWith != nil {
			initiator.InBattleWith.InBattleWith = nil
			initiator.InBattleWith = nil
		}

		// Если цель кого-то окружала, то её жертва освобождается
		if target.InBattleWith != nil && target.Strength <= target.InBattleWith.Strength {
			target.InBattleWith.InBattleWith = nil
			target.InBattleWith = nil
		}

		queue.Intercept(target, []int{
			EventTypes.REACTION_MELLEE_COUNTER_ATTACK,
		})
	case EventTypes.ACTION_SURROUND:
		logger.LogEvent("Surround")
		damage := attackHandler.Handle(initiator, target)
		debuff := spiritSaveHandler.Handle(target, damage)

		logger.LogAttack(initiator, target, damage, debuff)

		target.Strength -= damage
		target.Spirit -= debuff

		// Если инициатор кого-то окружил в прошлом ходу, он больше его не окружает
		if initiator.InBattleWith != nil {
			initiator.InBattleWith.InBattleWith = nil
			initiator.InBattleWith = nil
		}

		// Если у инициатора больше численность, он окружает свою цель
		if initiator.Strength > target.Strength {
			if target.InBattleWith != nil {
				target.InBattleWith.InBattleWith = nil
				target.InBattleWith = nil
			}

			initiator.InBattleWith = target
			target.InBattleWith = initiator
		}

		queue.Intercept(target, []int{
			EventTypes.REACTION_MELLEE_COUNTER_ATTACK,
		})
	case EventTypes.REACTION_MELLEE_COUNTER_ATTACK:
		logger.LogEvent("Melee counterattack")
		damage := attackHandler.Handle(initiator, target)
		debuff := spiritSaveHandler.Handle(target, damage)

		logger.LogAttack(initiator, target, damage, debuff)

		target.Strength -= damage
		target.Spirit -= debuff

		// Если численность инициатора сравнялась с численностью цели, то инициатор выходит из окружения
		if initiator.Strength >= target.Strength && initiator.InBattleWith == target {
			initiator.InBattleWith = nil
			target.InBattleWith = nil
		}
		queue.ResolveInterception(true)
		queue.NexTurn()
	case EventTypes.ACTION_SKIP:
		logger.LogEvent("Action skip")
		queue.NexTurn()
	case EventTypes.REACTION_SKIP:
		logger.LogEvent("Reaction skip")
		queue.ResolveInterception(false)
		queue.NexTurn()
	}

}
