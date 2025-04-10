package drug

type Effect int

const (
	AntiGravity = iota
	Athletic
	Balding
	BrightEyed
	Calming
	CalorieDense
	Cyclopean
	Disorienting
	Electrifying
	Energizing
	Euphoric
	Explosive
	Focused
	Foggy
	Gingeritis
	Glowing
	Jennerising
	Laxative
	Lethal
	LongFaced
	Munchies
	Paranoia
	Refreshing
	Schizophrenia
	Sedating
	SeizureInducing
	Shrinking
	Slippery
	Smelly
	Sneaky
	Spicy
	ThoughtProvoking
	Toxic
	TropicThunder
	Zombifying
)

var (
	effectsNames = map[Effect]string{
		AntiGravity:      "Anti-Gravity",
		Athletic:         "Athletic",
		Balding:          "Balding",
		BrightEyed:       "Bright-Eyed",
		Calming:          "Calming",
		CalorieDense:     "Calorie-Dense",
		Cyclopean:        "Cyclopean",
		Disorienting:     "Disorienting",
		Electrifying:     "Electrifying",
		Energizing:       "Energizing",
		Euphoric:         "Euphoric",
		Explosive:        "Explosive",
		Focused:          "Focused",
		Foggy:            "Foggy",
		Gingeritis:       "Gingeritis",
		Glowing:          "Glowing",
		Jennerising:      "Jennerising",
		Laxative:         "Laxative",
		Lethal:           "Lethal",
		LongFaced:        "Long Faced",
		Munchies:         "Munchies",
		Paranoia:         "Paranoia",
		Refreshing:       "Refreshing",
		Schizophrenia:    "Schizophrenia",
		Sedating:         "Sedating",
		SeizureInducing:  "Seizure-Inducing",
		Shrinking:        "Shrinking",
		Slippery:         "Slippery",
		Smelly:           "Smelly",
		Sneaky:           "Sneaky",
		Spicy:            "Spicy",
		ThoughtProvoking: "Thought-Provoking",
		Toxic:            "Toxic",
		TropicThunder:    "Tropic Thunder",
		Zombifying:       "Zombifying",
	}

	// Map for enum to multiplier
	effectsMultipliers = map[Effect]float64{
		AntiGravity:      1.54,
		Athletic:         1.32,
		Balding:          1.30,
		BrightEyed:       1.40,
		Calming:          1.10,
		CalorieDense:     1.28,
		Cyclopean:        1.56,
		Disorienting:     1.00,
		Electrifying:     1.50,
		Energizing:       1.22,
		Euphoric:         1.18,
		Explosive:        1.00,
		Focused:          1.16,
		Foggy:            1.36,
		Gingeritis:       1.20,
		Glowing:          1.48,
		Jennerising:      1.42,
		Laxative:         1.00,
		Lethal:           1.00,
		LongFaced:        1.52,
		Munchies:         1.12,
		Paranoia:         1.00,
		Refreshing:       1.14,
		Schizophrenia:    1.00,
		Sedating:         1.26,
		SeizureInducing:  1.00,
		Shrinking:        1.60,
		Slippery:         1.34,
		Smelly:           1.00,
		Sneaky:           1.24,
		Spicy:            1.38,
		ThoughtProvoking: 1.44,
		Toxic:            1.00,
		TropicThunder:    1.46,
		Zombifying:       1.58,
	}
)

func (e Effect) GetName() string {
	return effectsNames[e]
}

func (e Effect) GetMultiplier() float64 {
	return effectsMultipliers[e]
}
