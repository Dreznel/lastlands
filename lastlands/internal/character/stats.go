package character

import "github.com/dreznel/lastlands/v2/lastlands/internal/roller"

type BasicStats struct {
	Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma AbilityScore
}

func GenerateBasicStats() BasicStats {
	var statValues map[string]int
	statValues["Strength"] = roller.RollAbilityScore()
	statValues["Dexterity"] = roller.RollAbilityScore()
	statValues["Constitution"] = roller.RollAbilityScore()
	statValues["Intelligence"] = roller.RollAbilityScore()
	statValues["Wisdom"] = roller.RollAbilityScore()
	statValues["Charisma"] = roller.RollAbilityScore()

	// TODO: Make the Swap values come from user input.
	swap1 := "Strength"
	swap2 := "Charisma"

	tmp := statValues[swap1]
	statValues[swap1] = statValues[swap2]
	statValues[swap2] = tmp

	return BasicStats{
		Strength:     CreateAbilityScore(statValues["Strength"]),
		Dexterity:    CreateAbilityScore(statValues["Dexterity"]),
		Constitution: CreateAbilityScore(statValues["Constitution"]),
		Intelligence: CreateAbilityScore(statValues["Intelligence"]),
		Wisdom:       CreateAbilityScore(statValues["Wisdom"]),
		Charisma:     CreateAbilityScore(statValues["Charisma"]),
	}
}

type AbilityScore struct {
	Bonus   int
	Defense int
}

func CreateAbilityScore(bonusValue int) AbilityScore {
	return AbilityScore{
		Bonus:   bonusValue,
		Defense: 10 + bonusValue,
	}
}
