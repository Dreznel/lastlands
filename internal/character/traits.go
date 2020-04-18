package character

type TraitManager struct {
	Traits map[string]Trait
}

type Trait struct {
	Name        string
	Rank        int
	Description string
}
