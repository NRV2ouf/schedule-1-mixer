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
		AntiGravity:      0.54,
		Athletic:         0.32,
		Balding:          0.30,
		BrightEyed:       0.40,
		Calming:          0.10,
		CalorieDense:     0.28,
		Cyclopean:        0.56,
		Disorienting:     0.00,
		Electrifying:     0.50,
		Energizing:       0.22,
		Euphoric:         0.18,
		Explosive:        0.00,
		Focused:          0.16,
		Foggy:            0.36,
		Gingeritis:       0.20,
		Glowing:          0.48,
		Jennerising:      0.42,
		Laxative:         0.00,
		Lethal:           0.00,
		LongFaced:        0.52,
		Munchies:         0.12,
		Paranoia:         0.00,
		Refreshing:       0.14,
		Schizophrenia:    0.00,
		Sedating:         0.26,
		SeizureInducing:  0.00,
		Shrinking:        0.60,
		Slippery:         0.34,
		Smelly:           0.00,
		Sneaky:           0.24,
		Spicy:            0.38,
		ThoughtProvoking: 0.44,
		Toxic:            0.00,
		TropicThunder:    0.46,
		Zombifying:       0.58,
	}
)

func (e Effect) GetName() string {
	return effectsNames[e]
}

func (e Effect) GetMultiplier() float64 {
	return effectsMultipliers[e]
}
