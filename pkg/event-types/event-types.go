package EventTypes

const (
	ACTION_MELLEE_ATTACK int = iota
	/*
	    попытка окружить цель (невозможно окружить более быструю цель)
		окруженную цель может бить только тот кто окружил
		окруженная цель может бить только того кто окружил
		окруженная цель не может использовать дальнобойные атаки
		скорость окруженной и окружающей цели считается равной 1
	*/
	ACTION_SURROUND 

	ACITON_RANGED_ATTACK

	ACTION_SKIP

	REACTION_SKIP

	REACTION_MELLEE_COUNTER_ATTACK

	// REACTION_MEELLE_SECOND_WILL_COUNTER_ATTACK
)