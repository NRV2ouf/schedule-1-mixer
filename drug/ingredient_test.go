package drug

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	allIngredients = []Ingredient{
		Cuke,
		Banana,
		Paracetamol,
		Donut,
		Viagra,
		MouthWash,
		FluMedicine,
		Gasoline,
		EnergyDrink,
		MotorOil,
		MegaBean,
		Chili,
		Battery,
		Iodine,
		Addy,
		HorseSemen,
	}
)

func TestMixCombinaisons(t *testing.T) {
	testCases := []struct {
		name        string
		batchSize   int
		expectedMix [][]Ingredient
	}{
		{
			name:        "nominal batchSize -1",
			batchSize:   -1,
			expectedMix: [][]Ingredient{},
		},
		{
			name:        "nominal batchSize 0",
			batchSize:   0,
			expectedMix: [][]Ingredient{},
		},
		{
			name:      "nominal batchSize 1",
			batchSize: 1,
			expectedMix: func() [][]Ingredient {
				m := make([][]Ingredient, 0)
				for _, ing := range allIngredients {
					m = append(m, []Ingredient{ing})
				}
				return m
			}(),
		},
		{
			name:        "nominal batchSize 2",
			batchSize:   2,
			expectedMix: nil,
		},
		{
			name:        "nominal batchSize 3",
			batchSize:   3,
			expectedMix: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mixIterator := MixCombinaisons(allIngredients, tc.batchSize)
			counter := 0
			mixes := make([][]Ingredient, 0)

			for m := range mixIterator {
				counter++
				if tc.expectedMix != nil {
					mixes = append(mixes, m)
				}
			}

			assert.Equal(t,
				int(math.Pow(float64(len(allIngredients)), float64(tc.batchSize))),
				counter)

			if tc.expectedMix != nil { // Check for slice equivalence
				for _, m := range tc.expectedMix {
					assert.Contains(t, mixes, m)
				}
			} else { // Check for doubles
				m := map[string]bool{}
				var key bytes.Buffer

				for _, mix := range mixes {
					for _, ing := range mix {
						key.WriteString(fmt.Sprint(ing))
					}
					assert.False(t, m[key.String()])
					m[key.String()] = true
					key.Reset()
				}
			}
		})
	}
}
