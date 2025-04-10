package drug

type effectList uint64

func (el effectList) has(e Effect) bool {
	return (el & (1 << e)) != 0
}

func (el *effectList) set(e Effect) {
	*el = *el | (1 << e)
}

func (el *effectList) remove(e Effect) {
	*el = *el & ^(1 << e)
}

func (el effectList) get() []Effect {
	var effects []Effect
	for e, _ := range effectsNames {
		if el.has(e) {
			effects = append(effects, e)
		}
	}
	return effects
}
