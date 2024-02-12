package domain

import (
	"fmt"
)

type Plateau struct {
	Coordinates Coordinates
	Mowers      []*Mower
}

const MINIMUM_COORDINATE_VALUE int = 1

func CreatePlateau(coordinates Coordinates) (Plateau, error) {
	if coordinates.X < MINIMUM_COORDINATE_VALUE || coordinates.Y < MINIMUM_COORDINATE_VALUE {
		return Plateau{}, fmt.Errorf("coordinates with values under %d for plateau", MINIMUM_COORDINATE_VALUE)
	}

	return Plateau{Coordinates: coordinates}, nil
}

func (p *Plateau) ValidateMowerPosition(m *Mower) error {

	if m.GetPosition().X > p.Coordinates.X || m.GetPosition().Y > p.Coordinates.Y {
		return fmt.Errorf("mower with coordiantes %s is out of plateau area %s",
			m.GetPosition().String(), p.Coordinates.String())
	}

	for _, mc := range p.Mowers {
		if m == mc {
			continue
		}
		if mc.GetPosition().X == m.GetPosition().X && mc.GetPosition().Y == m.GetPosition().Y {
			return fmt.Errorf("there is another mower in the position: %d:%d", mc.GetPosition().X, mc.GetPosition().Y)
		}
	}
	return nil
}

func (p *Plateau) AddMower(mower *Mower) error {

	err := p.ValidateMowerPosition(mower)
	if err != nil {
		return err
	}
	p.Mowers = append(p.Mowers, mower)

	return nil
}
