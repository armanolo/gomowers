package domain

import "errors"

type MowerPosition struct {
	Coordinates Coordinates
	Cardinal    Cardinal
}

func MowerPositionFromString(position string) (MowerPosition, error) {

	if len(position) != 3 {
		return MowerPosition{}, errors.New("bad mower position format")
	}

	cardinal, err := CardinalFromString(position[2:])

	_ = cardinal

	coord, err := CoordinatesFromString(position[:2])

	_ = coord

	_ = err

	return MowerPosition{}, nil
}
