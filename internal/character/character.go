package character

import (
	"github.com/dreznel/lastlands/v2/internal/roller"
)

type Seeker struct {
	Name              string
	AbilityScores     BasicStats
	ArmorDefense      AbilityScore
	Resolve           int
	Fade              int
	HitPointsMax      int
	HitPointsCurrent  int
	CoreAttributes    CoreAttributes
	CollectedMemories []CollectedMemory
	Inventory         CharacterInventory
	Traits            TraitManager
}

// Top-level CreateCharacter function.
func CreateCharacter() Seeker {
	baseStats := GenerateBasicStats()
	hp := roller.RollDice(1, 8)[0]

	return Seeker{
		Name:             "Dreznel",
		AbilityScores:    baseStats,
		ArmorDefense:     CreateAbilityScore(1),
		Resolve:          baseStats.Charisma.Bonus,
		HitPointsMax:     hp,
		HitPointsCurrent: hp,
		CoreAttributes: CoreAttributes{
			First: &CoreMemory{
				IsMemoryLost: false,
				Description:  "I was once invincible",
			},
			Second: &CoreMemory{
				IsMemoryLost: false,
				Description:  "The people were fooled by my enemy.",
			},
			Third: &CoreMemory{
				IsMemoryLost: false,
				Description:  "My best friend was a scholar",
			},
		},
		CollectedMemories: []CollectedMemory{},
		Inventory: CharacterInventory{
			Capacity:  baseStats.Constitution.Defense,
			Items:     []Item{},
			OpenSlots: baseStats.Constitution.Defense,
		},
		Traits: TraitManager{
			Traits: map[string]Trait{
				"Weapon Proficiency: Heavy Slashing": Trait{
					Name:        "Weapon Proficiency: Heavy Slashing",
					Rank:        1,
					Description: "Add your Strength bonus to ATTACK ROLLS ONLY with Heavy Slashing weapons.",
				},
			},
		},
	}
}
