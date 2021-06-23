package domain

// D2s109Character represents a Diablo II character.
type D2s109Character struct {
	CharacterName   string
	CharacterStatus characterStatus
	CharacterClass  string
	CharacterLevel  int16
	LastPlayed      string
	Attributes      attributes
	Items           []item
	//Skills          map[string]interface{} // Rest of the fields should go here.
}

type characterStatus struct {
	Expansion bool
	Dead      bool
	HardCore  bool
}

type attributes struct {
	Strength            int16
	Energy              int16
	Dexterity           int16
	Vitality            int16
	RemainingStatPoints int16
	CurrentLife         int16
	BaseLife            int16
	CurrentMana         int16
	BaseMana            int16
	CurrentStamina      int16
	BaseStamina         int16
	Level               int16
	Experience          int16
	InventoryGold       int16
	StashGold           int16
}
type item struct {
	Identified           bool
	Socketed             bool
	Ethereal             bool
	Personalized         bool
	Runeword             bool
	PositionX            int16
	PositionY            int16
	ItemName             string
	ItemLocation         string
	ItemPosition         string
	ItemLevel            int16
	ItemCode             string
	Vibrant              bool
	TotalNumberOfSockets int16
}
