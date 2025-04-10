package drug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReact(t *testing.T) {
	testCases := []struct {
		name string

		// reaction
		wanted   []Effect
		unwanted []Effect
		products []Effect

		// effectLists
		baseEffects     []Effect
		expectedEffects []Effect
	}{
		{
			name:            "nominal reaction emtpy base empty",
			wanted:          []Effect{},
			unwanted:        []Effect{},
			products:        []Effect{},
			baseEffects:     []Effect{},
			expectedEffects: []Effect{},
		},
		{
			name:            "nominal reaction emtpy base not empty",
			wanted:          []Effect{},
			unwanted:        []Effect{},
			products:        []Effect{},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{AntiGravity, Athletic},
		},
		{
			name:            "nominal reaction should not occur",
			wanted:          []Effect{Electrifying},
			unwanted:        []Effect{},
			products:        []Effect{Calming},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{AntiGravity, Athletic},
		},
		{
			name:            "nominal reaction should not occur missing unwanted",
			wanted:          []Effect{AntiGravity},
			unwanted:        []Effect{Cyclopean},
			products:        []Effect{Calming},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{AntiGravity, Athletic},
		},
		{
			name:            "nominal reaction should not occur missing wanted",
			wanted:          []Effect{Electrifying},
			unwanted:        []Effect{Cyclopean},
			products:        []Effect{Calming},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{AntiGravity, Athletic},
		},
		{
			name:            "nominal reaction occurs",
			wanted:          []Effect{AntiGravity},
			unwanted:        []Effect{},
			products:        []Effect{Calming},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming, Athletic},
		},
		{
			name:            "nominal reaction occurs with unwanted",
			wanted:          []Effect{AntiGravity},
			unwanted:        []Effect{Zombifying},
			products:        []Effect{Calming},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming, Athletic},
		},
		{
			name:            "nominal reaction occurs effect already present",
			wanted:          []Effect{AntiGravity},
			unwanted:        []Effect{},
			products:        []Effect{Athletic},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Athletic},
		},
		{
			name:            "nominal reaction occurs with several wanted",
			wanted:          []Effect{AntiGravity, Athletic},
			unwanted:        []Effect{},
			products:        []Effect{Electrifying},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming},
		},
		{
			name:            "nominal reaction does not occur with several unwanted",
			wanted:          []Effect{AntiGravity},
			unwanted:        []Effect{Zombifying, Athletic},
			products:        []Effect{Electrifying},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming},
		},
		{
			name:            "nominal reaction occurs with several unwanted",
			wanted:          []Effect{AntiGravity},
			unwanted:        []Effect{Zombifying, Balding},
			products:        []Effect{Electrifying},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming, Electrifying},
		},
		{
			name:            "nominal reaction occurs empty wanted",
			wanted:          []Effect{},
			unwanted:        []Effect{},
			products:        []Effect{Electrifying},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming, Athletic, Electrifying},
		},
		{
			name:            "nominal reaction does not occur empty wanted with unwanted",
			wanted:          []Effect{},
			unwanted:        []Effect{AntiGravity},
			products:        []Effect{Electrifying},
			baseEffects:     []Effect{AntiGravity, Athletic},
			expectedEffects: []Effect{Calming, Athletic, Electrifying},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reaction
			r := reaction{}
			for _, e := range tc.wanted {
				r.with(e)
			}
			for _, e := range tc.unwanted {
				r.without(e)
			}
			for _, e := range tc.unwanted {
				r.without(e)
			}

			// Effect Lists
			base := effectList(0)
			for _, e := range tc.baseEffects {
				base.set(e)
			}
			expected := effectList(0)
			for _, e := range tc.baseEffects {
				expected.set(e)
			}

			result := r.react(base)
			assert.Equal(t, expected, result)
		})
	}
}
