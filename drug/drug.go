package drug

type Drug struct {
	effects     effectList
	ingredients []Ingredient
	base        string
	basePrice   int
}

// Add adds an Ingredient to a drug. It updates its Effects.
// It returns a boolean indicating if the operation did something.
func (d *Drug) Add(i Ingredient) {
	for _, r := range mixMap[i] {
		d.effects = r.react(d.effects)
	}

	d.effects.set(ingredientBaseEffects[i])
}

func (d Drug) GetEffects() []Effect {
	return d.effects.get()
}

func (d Drug) GetIngredients() []Ingredient {
	return d.ingredients
}

func (d Drug) GetBase() string {
	return d.base
}

func (d Drug) GetBasePrice() int {
	return d.basePrice
}

func (d Drug) GetPrice() int {
	m := 1.
	for _, e := range d.GetEffects() {
		m *= effectsMultipliers[e]
	}
	return int(m * float64(d.basePrice))
}

func (d Drug) GetMixPrice() int {
	p := 0
	for _, ing := range d.GetIngredients() {
		p += ingredientPrices[ing]
	}
	return p
}

// Grass
func GetOgKush() Drug {
	el := effectList(0)
	el.set(Calming)
	return Drug{
		effects:     el,
		ingredients: []Ingredient{},
		base:        "OgKush",
		basePrice:   38,
	}
}

func GetSourDiesel() Drug {
	el := effectList(0)
	el.set(Refreshing)
	return Drug{
		effects:     el,
		ingredients: []Ingredient{},
		base:        "SourDiesel",
		basePrice:   40,
	}
}

func GetGreenCrack() Drug {
	el := effectList(0)
	el.set(Energizing)
	return Drug{
		effects:     el,
		ingredients: []Ingredient{},
		base:        "GreenCrack",
		basePrice:   43,
	}
}

func GetGrandaddyPurple() Drug {
	el := effectList(0)
	el.set(Sedating)
	return Drug{
		effects:     el,
		ingredients: []Ingredient{},
		base:        "GrandaddyPurple",
		basePrice:   44,
	}
}

func GetMethamphetamine() Drug {
	return Drug{
		effects:     effectList(0),
		ingredients: []Ingredient{},
		base:        "Methamphetamine",
		basePrice:   70,
	}
}

// To keep in mind:
// For cocaine quality adds an aditional multiplier
//
//	Standard: 1
//	Premium: 1.5
//	Hevenly: 2
func GetCocaine() Drug {
	return Drug{
		effects:     effectList(0),
		ingredients: []Ingredient{},
		base:        "Cocaine",
		basePrice:   90,
	}
}
