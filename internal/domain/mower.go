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

func (m *Mower) GetPosition() Coordinates {
	return m.MowerPosition.Coordinates
}

func (m *Mower) MovementProcess() (string, error) {

	for _, nm := range m.Path.Movement {
		if err := m.move(nm); err != nil {
			return "", err
		}
	}

	return m.MowerPosition.String(), nil
}

func (m *Mower) move(movement string) error {

	coord := m.MowerPosition.Coordinates
	card := m.MowerPosition.Cardinal
	var err error

	if movement == PATH_M {
		x := coord.X
		y := coord.Y

		switch nm := m.MowerPosition.Cardinal.Letter; nm {
		case CARDINAL_E:
			x++
		case CARDINAL_W:
			x--
		case CARDINAL_N:
			y++
		case CARDINAL_S:
			y--
		default:
		}

		coord, err = CreateCoordinates(x, y)
		if err != nil {
			return err
		}
	} else {
		card, err = m.MowerPosition.Cardinal.TurnTo(movement)
		if err != nil {
			return err
		}
	}

	nmp, err := CreateMowerPosition(coord, card)
	if err != nil {
		return nil
	}

	m.MowerPosition = nmp

	return nil
}
