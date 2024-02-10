package infrastructure

import (
	"errors"
	"mowers/internal/domain"
	"strings"
)

func ValidateInput(coordinates string, p *domain.Plateau, ml []domain.Mower) error {
	if coordinates == "" {
		return errors.New("input is empty")
	}

	listCoordinates := strings.Split(coordinates, "\n")

	if len(listCoordinates) < 3 || len(listCoordinates)%2 == 0 {
		return errors.New("bad format input data")
	}

	err := ValidatePlateu(listCoordinates[0], p)
	if err != nil {
		return err
	}

	for i := 1; i < len(listCoordinates); i++ {
		if i%2 == 1 {
			mp, err := domain.MowerPositionFromString(listCoordinates[i])
			if err != nil {
				return err
			}
			path, err := domain.PathFromString(listCoordinates[i+1])
			if err != nil {
				return err
			}

			nm, err := domain.CreateMower(p, mp, path)
			if err != nil {
				return err
			}
			ml = append(ml, nm)
		}
	}

	return nil
}

func ValidatePlateu(coordinates string, p *domain.Plateau) error {

	if c, err := domain.CoordinatesFromString(coordinates); err != nil {
		return err
	} else {
		p.Coordinates = domain.Coordinates{X: c.X, Y: c.Y}
		return nil
	}

}
