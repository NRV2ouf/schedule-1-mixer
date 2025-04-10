package main

import (
	"fmt"
	"schedule-1/drug"
)

const (
	pipeSize = 5
)

func main() {
	allIngredients := []drug.Ingredient{
		drug.Cuke,
		drug.Banana,
		drug.Paracetamol,
		drug.Donut,
		drug.Viagra,
		drug.MouthWash,
		drug.FluMedicine,
		drug.Gasoline,
		drug.EnergyDrink,
		drug.MotorOil,
		drug.MegaBean,
		drug.Chili,
		drug.Battery,
		drug.Iodine,
		drug.Addy,
		drug.HorseSemen,
	}

	best := struct {
		effects         []drug.Effect
		ingredients     []drug.Ingredient
		price           int
		productionPrice int
	}{
		price:           -1,
		productionPrice: -1,
	}

	mixes := drug.MixCombinaisons(allIngredients, pipeSize)
	for mix := range mixes {
		d := drug.GetMethamphetamine()
		for _, ing := range mix {
			d.Add(ing)
		}
		p := d.GetPrice()
		mixP := d.GetMixPrice()
		if p > best.price ||
			(p == best.price && mixP < best.productionPrice) {
			best.effects = make([]drug.Effect, len(d.GetEffects()))
			copy(best.effects, d.GetEffects())
			best.ingredients = make([]drug.Ingredient, len(mix))
			copy(best.ingredients, mix)
			best.price = p
			best.productionPrice = mixP
		}

	}

	var effectList []string
	for _, e := range best.effects {
		effectList = append(effectList, e.GetName())
	}

	var ingredientList []string
	for _, ing := range best.ingredients {
		ingredientList = append(ingredientList, ing.GetName())
	}

	fmt.Printf("Best combo effects: %v\n", effectList)
	fmt.Printf("Ingredients: %v\n", ingredientList)
	fmt.Printf("Final price: %d\n", best.price)
}
