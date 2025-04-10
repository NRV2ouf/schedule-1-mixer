package drug

type Ingredient int

const (
	Cuke = iota
	Banana
	Paracetamol
	Donut
	Viagra
	MouthWash
	FluMedicine
	Gasoline
	EnergyDrink
	MotorOil
	MegaBean
	Chili
	Battery
	Iodine
	Addy
	HorseSemen
)

var (
	ingredientNames = map[Ingredient]string{
		Cuke:        "Cuke",
		Banana:      "Banana",
		Paracetamol: "Paracetamol",
		Donut:       "Donut",
		Viagra:      "Viagra",
		MouthWash:   "Mouth Wash",
		FluMedicine: "Flu Medicine",
		Gasoline:    "Gasoline",
		EnergyDrink: "Energy Drink",
		MotorOil:    "Motor Oil",
		MegaBean:    "Mega Bean",
		Chili:       "Chili",
		Battery:     "Battery",
		Iodine:      "Iodine",
		Addy:        "Addy",
		HorseSemen:  "Horse Semen",
	}

	ingredientBaseEffects = map[Ingredient]Effect{
		Cuke:        Energizing,
		Banana:      Gingeritis,
		Paracetamol: Sneaky,
		Donut:       CalorieDense,
		Viagra:      TropicThunder,
		MouthWash:   Balding,
		FluMedicine: Sedating,
		Gasoline:    Toxic,
		EnergyDrink: Athletic,
		MotorOil:    Slippery,
		MegaBean:    Foggy,
		Chili:       Spicy,
		Battery:     BrightEyed,
		Iodine:      Jennerising,
		Addy:        ThoughtProvoking,
		HorseSemen:  LongFaced,
	}

	ingredientPrices = map[Ingredient]int{
		Cuke:        2,
		Banana:      2,
		Paracetamol: 3,
		Donut:       3,
		Viagra:      4,
		MouthWash:   4,
		FluMedicine: 5,
		Gasoline:    5,
		EnergyDrink: 6,
		MotorOil:    6,
		MegaBean:    7,
		Chili:       7,
		Battery:     8,
		Iodine:      8,
		Addy:        9,
		HorseSemen:  9,
	}
)

func (i Ingredient) GetName() string {
	return ingredientNames[i]
}

// MixCombinaisons returns a channel that produces all possible combinations of size n
// from the given ingredients. The channel will be closed when all combinations are exhausted.
func MixCombinaisons(ingredients []Ingredient, n int) <-chan []Ingredient {
	ch := make(chan []Ingredient)

	go func() {
		if n < 0 {
			close(ch)
			return
		}

		defer close(ch)
		current := make([]Ingredient, n)

		var generate func(int, []Ingredient)
		generate = func(size int, current []Ingredient) {
			if size <= 0 {
				combination := make([]Ingredient, len(current))
				copy(combination, current)
				ch <- combination
				return
			}

			for _, ing := range ingredients {
				current[size-1] = ingredients[ing]
				generate(size-1, current)
			}
		}

		generate(n, current)
	}()

	return ch
}
