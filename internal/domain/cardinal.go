package domain

import (
	"errors"
	"strings"
)

var CARDINALS = [4]string{"N", "S", "W", "E"}

type Cardinal struct {
	Lettler string
}

func CardinalFromString(character string) (Cardinal, error) {
	var l string = ""

	if len(character) != 1 {
		return Cardinal{}, errors.New("bad cardinat format")
	}

	lu := strings.ToUpper(character)

	for _, a := range CARDINALS {
		if a == lu {
			l = a
		}
	}

	if l == "" {
		return Cardinal{}, errors.New("bad cardinal value")
	}

	return Cardinal{Lettler: l}, nil
}
