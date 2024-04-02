package position

// Position represents a Row and Col position on a grid
type Position struct {
	R, C int
}

func (p *Position) Up() {
	if p.R > 0 {
		p.R -= 1
	}
}

func (p *Position) Down() {
	if p.R < 2 {
		p.R += 1
	}
}

func (p *Position) Left() {
	if p.C > 0 {
		p.C -= 1
	}
}

func (p *Position) Right() {
	if p.C < 2 {
		p.C += 1
	}
}
