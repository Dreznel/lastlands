package character

type CharacterInventory struct {
	OpenSlots int
	Capacity  int
	Items     []Item
}

// Make methods to add and remove items

type Item struct {
	Name       string
	Slots      int
	Durability int
	Stacksize  int // amount that can stack in the given number of slots. Will probably complicate later
	Details    string
}
