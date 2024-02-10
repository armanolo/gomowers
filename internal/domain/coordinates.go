package domain

import (
	"errors"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func CoordinatesFromString(coordinates string) (Coordinates, error) {
	c := strings.Split(coordinates, "")
	if len(c) != 2 {
		return Coordinates{}, errors.New("bad coordinate format")
	}

	cx, err := strconv.Atoi(c[0])
	if err != nil {
		return Coordinates{}, errors.New("bad coodinate x")
	}
	cy, err := strconv.Atoi(c[1])
	if err != nil {
		return Coordinates{}, errors.New("bad coodinate y")
	}

	return Coordinates{X: cx, Y: cy}, nil
}
