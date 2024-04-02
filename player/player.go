package player

// Symbol represents a players icon on the board
type Symbol byte

func (s Symbol) String() string {
	return string(s)
}

// Player holds a player's information
type Player struct {
	Name   string
	Symbol Symbol
}

func (p *Player) AssignInfo(name string, symbol Symbol) {
	p.Name = name
	p.Symbol = symbol
}
