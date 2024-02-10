package domain

import (
	"fmt"
)

type Plateau struct {
	Coordinates    Coordinates
	MowerPositions []Coordinates
}

func (p *Plateau) ValidateMowerPosition(coordinates Coordinates) error {
	for _, mc := range p.MowerPositions {
		if mc.X == coordinates.X && mc.Y == coordinates.Y {
			return fmt.Errorf("there is another mower in the position: %d:%d", mc.X, mc.Y)
		}
	}
	return nil
}

func (p *Plateau) AddMowerPosition(coordinates Coordinates) error {

	err := p.ValidateMowerPosition(coordinates)
	if err != nil {
		return err
	}

	p.MowerPositions = append(p.MowerPositions, coordinates)
	return nil
}
