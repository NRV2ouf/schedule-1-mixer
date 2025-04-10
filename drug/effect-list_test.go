package drug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEffectListSetHas(t *testing.T) {
	testCases := []struct {
		name    string
		effects []Effect
	}{
		{
			name:    "nominal single effect",
			effects: []Effect{AntiGravity},
		},
		{
			name:    "nominal two effects",
			effects: []Effect{AntiGravity, Athletic},
		},
		{
			name:    "nominal five effects",
			effects: []Effect{AntiGravity, Athletic, Balding, BrightEyed, Calming},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			el := effectList(0)
			for _, e := range tc.effects {
				el.set(e)
			}

			for _, e := range tc.effects {
				assert.True(t, el.has(e))
			}
		})
	}
}

func TestEffectListGet(t *testing.T) {
	testCases := []struct {
		name    string
		effects []Effect
	}{
		{
			name:    "nominal single effect",
			effects: []Effect{AntiGravity},
		},
		{
			name:    "nominal two effects",
			effects: []Effect{AntiGravity, Athletic},
		},
		{
			name:    "nominal five effects",
			effects: []Effect{AntiGravity, Athletic, Balding, BrightEyed, Calming},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			el := effectList(0)
			for _, e := range tc.effects {
				el.set(e)
			}

			for _, e := range el.get() {
				assert.Contains(t, tc.effects, e)
			}
		})
	}
}
