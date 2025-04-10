package drug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name            string
		base            Drug
		ingredients     []Ingredient
		expectedEffects []Effect
	}{
		{
			name:            "nominal no ingredients",
			base:            GetMethamphetamine(),
			ingredients:     []Ingredient{},
			expectedEffects: nil,
		},
		{
			name:            "nominal test 1",
			base:            GetCocaine(),
			ingredients:     []Ingredient{MotorOil, Cuke, Paracetamol, Gasoline, Cuke, Battery, HorseSemen, MegaBean},
			expectedEffects: []Effect{AntiGravity, Glowing, TropicThunder, Zombifying, Cyclopean, BrightEyed, LongFaced, Foggy},
		},
		{
			name:            "nominal test 2",
			base:            GetOgKush(),
			ingredients:     []Ingredient{Banana, Gasoline, Paracetamol, Cuke, MegaBean, Battery, Banana, Cuke},
			expectedEffects: []Effect{TropicThunder, AntiGravity, Zombifying, Jennerising, Glowing, Cyclopean, BrightEyed, ThoughtProvoking},
		},
		{
			name:            "nominal test 3",
			base:            GetMethamphetamine(),
			ingredients:     []Ingredient{Banana, Cuke, Paracetamol, Gasoline, Cuke, Battery, HorseSemen, MegaBean},
			expectedEffects: []Effect{Electrifying, Glowing, TropicThunder, Zombifying, Cyclopean, BrightEyed, LongFaced, Foggy},
		},
		{
			name:            "nominal Shrinking",
			base:            GetGrandaddyPurple(),
			ingredients:     []Ingredient{Cuke, Banana, Chili, EnergyDrink},
			expectedEffects: []Effect{Gingeritis, Shrinking, Munchies, Euphoric, Athletic},
		},
		{
			name:            "nominal Schizophrenia",
			base:            GetGreenCrack(),
			ingredients:     []Ingredient{MotorOil, MotorOil},
			expectedEffects: []Effect{Slippery, Schizophrenia},
		},
		{
			name:            "nominal Glowing",
			base:            GetSourDiesel(),
			ingredients:     []Ingredient{Cuke, MegaBean, Battery},
			expectedEffects: []Effect{Refreshing, Foggy, Glowing, BrightEyed},
		},
		{
			name:            "nominal Explosive",
			base:            GetMethamphetamine(),
			ingredients:     []Ingredient{Donut, Donut},
			expectedEffects: []Effect{Explosive, CalorieDense},
		},
		{
			name:            "nominal Cyclopean",
			base:            GetMethamphetamine(),
			ingredients:     []Ingredient{Cuke, MegaBean},
			expectedEffects: []Effect{Cyclopean, Foggy},
		},
		{
			name:            "nominal AntiGravity",
			base:            GetMethamphetamine(),
			ingredients:     []Ingredient{Gasoline, Banana, Banana},
			expectedEffects: []Effect{Gingeritis, AntiGravity},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := tc.base
			for _, ing := range tc.ingredients {
				d.Add(ing)
			}

			res := d.GetEffects()
			assert.Equal(t, len(tc.expectedEffects), len(res))
		})
	}
}
