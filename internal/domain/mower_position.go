package domain

import (
	"errors"
)

type MowerPosition struct {
	Coordinates Coordinates
	Cardinal    Cardinal
}

func MowerPositionFromString(position string) (MowerPosition, error) {

	if len(position) != 3 {
		return MowerPosition{}, errors.New("bad mower position format")
	}

	cardinal, err := CreateCardinal(position[2:])

	if err != nil {
		return MowerPosition{}, err
	}

	coord, err := CoordinatesFromString(position[:2])

	if err != nil {
		return MowerPosition{}, err
	}

	return CreateMowerPosition(coord, cardinal)
}

func CreateMowerPosition(coordinates Coordinates, cardinal Cardinal) (MowerPosition, error) {

	return MowerPosition{Coordinates: coordinates, Cardinal: cardinal}, nil
}

func (c MowerPosition) String() string {
	return c.Coordinates.String() + c.Cardinal.String()
}
