package domain

type Mower struct {
	Plateu        *Plateau
	MowerPosition MowerPosition
	Path          Path
}

func CreateMower(p *Plateau, mp MowerPosition, path Path) (Mower, error) {
	m := Mower{Plateu: p, MowerPosition: mp, Path: path}
	err := p.AddMowerPosition(m.GetPosition())

	if err != nil {
		return Mower{}, err
	}
	return m, nil
}

func (m Mower) GetPosition() Coordinates {
	return m.MowerPosition.Coordinates
}

func (m Mower) Move() {

}
