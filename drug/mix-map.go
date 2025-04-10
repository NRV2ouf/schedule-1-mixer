package drug

import "fmt"

type reaction struct {
	// The effects required for a reaction to happen
	wanted []Effect
	// The effects preventing a reaction
	unwanted []Effect
	// The effects to add in case of success
	products []Effect
}

var (
	// Set within the init function
	mixMap = map[Ingredient][]reaction{}
)

func (r reaction) with(e Effect) reaction {
	r.wanted = append(r.wanted, e)
	return r
}

func (r reaction) without(e Effect) reaction {
	r.unwanted = append(r.unwanted, e)
	return r
}

func (r reaction) produces(e Effect) reaction {
	r.products = append(r.products, e)
	return r
}

func (r reaction) react(effects effectList) effectList {
	newEffects := effects

	for _, e := range r.wanted {
		if !effects.has(e) {
			return effects
		}
		newEffects.remove(e)
	}

	for _, e := range r.unwanted {
		if effects.has(e) {
			return effects
		}
	}

	for _, e := range r.products {
		newEffects.set(e)
	}

	return newEffects
}

func validatereactions() {
	for ing, reactions := range mixMap {
		seen := make(map[string]bool)
		for _, r := range reactions {
			key := fmt.Sprintf("%v-%v-%v", r.wanted, r.unwanted, r.products)
			if seen[key] {
				panic(fmt.Sprintf("duplicate reaction for %v: %v", ing, key))
			}
			seen[key] = true
		}
	}
}

func init() {
	mixMap[Addy] = []reaction{
		reaction{}.with(LongFaced).produces(Electrifying),
		reaction{}.with(Foggy).produces(Energizing),
		reaction{}.with(Explosive).produces(Euphoric),
		reaction{}.with(Sedating).produces(Gingeritis),
		reaction{}.with(Glowing).produces(Refreshing),
	}
	mixMap[Banana] = []reaction{
		reaction{}.with(Smelly).produces(AntiGravity),
		reaction{}.with(Disorienting).produces(Focused),
		reaction{}.with(Paranoia).produces(Jennerising),
		reaction{}.with(LongFaced).produces(Refreshing),
		reaction{}.with(Focused).produces(SeizureInducing),
		reaction{}.with(Toxic).produces(Smelly),
		reaction{}.with(Calming).produces(Sneaky),
		reaction{}.with(ThoughtProvoking).produces(Cyclopean),
		reaction{}.with(ThoughtProvoking).without(Energizing).produces(Energizing),
	}
	mixMap[Battery] = []reaction{
		reaction{}.with(Laxative).produces(CalorieDense),
		reaction{}.with(Electrifying).without(Zombifying).produces(Euphoric),
		reaction{}.with(Cyclopean).produces(Glowing),
		reaction{}.with(Shrinking).produces(Munchies),
		reaction{}.with(Munchies).produces(TropicThunder),
		reaction{}.with(Euphoric).without(Electrifying).produces(Zombifying),
	}
	mixMap[Chili] = []reaction{
		reaction{}.with(Sneaky).produces(BrightEyed),
		reaction{}.with(Athletic).produces(Euphoric),
		reaction{}.with(Laxative).produces(LongFaced),
		reaction{}.with(Shrinking).produces(Refreshing),
		reaction{}.with(Munchies).produces(Toxic),
		reaction{}.with(AntiGravity).produces(TropicThunder),
	}
	mixMap[Cuke] = []reaction{
		reaction{}.with(Munchies).produces(Athletic),
		reaction{}.with(Slippery).without(Munchies).produces(Athletic),
		reaction{}.with(Foggy).produces(Cyclopean),
		reaction{}.with(Toxic).produces(Euphoric),
		reaction{}.with(Euphoric).produces(Laxative),
		reaction{}.with(Slippery).produces(Munchies),
		reaction{}.with(Sneaky).produces(Paranoia),
		reaction{}.with(Gingeritis).produces(ThoughtProvoking),
	}
	mixMap[Donut] = []reaction{
		reaction{}.with(Shrinking).produces(Energizing),
		reaction{}.with(Focused).produces(Euphoric),
		reaction{}.with(CalorieDense).without(Explosive).produces(Explosive),
		reaction{}.with(Jennerising).produces(Gingeritis),
		reaction{}.with(AntiGravity).produces(Slippery),
		reaction{}.with(Balding).produces(Sneaky),
	}
	mixMap[EnergyDrink] = []reaction{
		reaction{}.with(Schizophrenia).produces(Balding),
		reaction{}.with(Glowing).produces(Disorienting),
		reaction{}.with(Disorienting).produces(Electrifying),
		reaction{}.with(Euphoric).produces(Energizing),
		reaction{}.with(Spicy).produces(Euphoric),
		reaction{}.with(Foggy).produces(Laxative),
		reaction{}.with(Sedating).produces(Munchies),
		reaction{}.with(Focused).produces(Shrinking),
		reaction{}.with(TropicThunder).produces(Sneaky),
	}
	mixMap[FluMedicine] = []reaction{
		reaction{}.with(Calming).produces(BrightEyed),
		reaction{}.with(Focused).produces(Calming),
		reaction{}.with(Laxative).produces(Euphoric),
		reaction{}.with(Cyclopean).produces(Foggy),
		reaction{}.with(ThoughtProvoking).produces(Gingeritis),
		reaction{}.with(Athletic).produces(Munchies),
		reaction{}.with(Shrinking).produces(Paranoia),
		reaction{}.with(Electrifying).produces(Refreshing),
		reaction{}.with(Munchies).produces(Slippery),
		reaction{}.with(Euphoric).produces(Toxic),
	}
	mixMap[Gasoline] = []reaction{
		reaction{}.with(Paranoia).produces(Calming),
		reaction{}.with(Electrifying).produces(Disorienting),
		reaction{}.with(Energizing).produces(Euphoric),
		reaction{}.with(Shrinking).produces(Focused),
		reaction{}.with(Laxative).produces(Foggy),
		reaction{}.with(Disorienting).produces(Glowing),
		reaction{}.with(Munchies).produces(Sedating),
		reaction{}.with(Gingeritis).produces(Smelly),
		reaction{}.with(Jennerising).produces(Sneaky),
		reaction{}.with(Energizing).produces(Spicy),
		reaction{}.with(Euphoric).without(Energizing).produces(Spicy),
		reaction{}.with(Sneaky).produces(TropicThunder),
	}
	mixMap[HorseSemen] = []reaction{
		reaction{}.with(AntiGravity).produces(Calming),
		reaction{}.with(ThoughtProvoking).produces(Electrifying),
		reaction{}.with(Gingeritis).produces(Refreshing),
	}
	mixMap[Iodine] = []reaction{
		reaction{}.with(CalorieDense).produces(Gingeritis),
		reaction{}.with(Foggy).produces(Paranoia),
		reaction{}.with(Calming).produces(Sedating),
		reaction{}.with(Euphoric).produces(SeizureInducing),
		reaction{}.with(Toxic).produces(Sneaky),
		reaction{}.with(Refreshing).produces(ThoughtProvoking),
	}
	mixMap[MegaBean] = []reaction{
		reaction{}.with(Sneaky).produces(Calming),
		reaction{}.with(ThoughtProvoking).produces(Cyclopean),
		reaction{}.with(Energizing).without(ThoughtProvoking).produces(Cyclopean),
		reaction{}.with(Focused).produces(Disorienting),
		reaction{}.with(Shrinking).produces(Electrifying),
		reaction{}.with(ThoughtProvoking).produces(Energizing),
		reaction{}.with(SeizureInducing).produces(Focused),
		reaction{}.with(Calming).produces(Glowing),
		reaction{}.with(Sneaky).produces(Glowing),
		reaction{}.with(Athletic).produces(Laxative),
		reaction{}.with(Jennerising).produces(Paranoia),
		reaction{}.with(Slippery).produces(Toxic),
	}
	mixMap[MotorOil] = []reaction{
		reaction{}.with(Paranoia).produces(AntiGravity),
		reaction{}.with(Energizing).produces(Munchies),
		reaction{}.with(Energizing).produces(Schizophrenia),
		reaction{}.with(Munchies).without(Energizing).produces(Schizophrenia),
		reaction{}.with(Euphoric).produces(Sedating),
		reaction{}.with(Foggy).produces(Toxic),
	}
	mixMap[MouthWash] = []reaction{
		reaction{}.with(Calming).produces(AntiGravity),
		reaction{}.with(Focused).produces(Jennerising),
		reaction{}.with(Explosive).produces(Sedating),
		reaction{}.with(CalorieDense).produces(Sneaky),
	}
	mixMap[Paracetamol] = []reaction{
		reaction{}.with(Munchies).produces(AntiGravity),
		reaction{}.with(Electrifying).produces(Athletic),
		reaction{}.with(Paranoia).produces(Balding),
		reaction{}.with(Energizing).without(Paranoia).produces(Balding),
		reaction{}.with(Spicy).produces(BrightEyed),
		reaction{}.with(Foggy).produces(Calming),
		reaction{}.with(Focused).produces(Gingeritis),
		reaction{}.with(Energizing).without(Munchies).produces(Paranoia),
		reaction{}.with(Calming).produces(Slippery),
		reaction{}.with(Glowing).produces(Toxic),
		reaction{}.with(Toxic).produces(TropicThunder),
	}
	mixMap[Viagra] = []reaction{
		reaction{}.with(Euphoric).produces(BrightEyed),
		reaction{}.with(Laxative).produces(Calming),
		reaction{}.with(Athletic).produces(Sneaky),
		reaction{}.with(Disorienting).produces(Toxic),
	}

	validatereactions()
}
