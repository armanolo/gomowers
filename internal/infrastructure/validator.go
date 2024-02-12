package infrastructure

import (
	"errors"
	"mowers/internal/domain"
	"strings"
)

func ValidateInput(coordinates string, p *domain.Plateau) ([]*domain.Mower, error) {
	ml := []*domain.Mower{}

	if coordinates == "" {
		return []*domain.Mower{}, errors.New("input is empty")
	}

	listCoordinates := strings.Split(coordinates, "\n")

	if len(listCoordinates) < 3 || len(listCoordinates)%2 == 0 {
		return []*domain.Mower{}, errors.New("bad format input data")
	}

	coord, err := domain.CoordinatesFromString(listCoordinates[0])
	if err != nil {
		return []*domain.Mower{}, err
	}

	*p, err = domain.CreatePlateau(coord)
	if err != nil {
		return []*domain.Mower{}, err
	}

	for i := 1; i < len(listCoordinates); i++ {
		if i%2 == 1 {
			mp, err := domain.MowerPositionFromString(listCoordinates[i])
			if err != nil {
				return []*domain.Mower{}, err
			}
			path, err := domain.PathFromString(listCoordinates[i+1])
			if err != nil {
				return []*domain.Mower{}, err
			}

			nm, err := domain.CreateMower(p, mp, path)
			if err != nil {
				return []*domain.Mower{}, err
			}
			ml = append(ml, nm)
		}
	}

	return ml, nil
}
