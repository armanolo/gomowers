package domain

import (
	"errors"
	"strings"
)

const (
	CARDINAL_N = "N"
	CARDINAL_S = "S"
	CARDINAL_W = "W"
	CARDINAL_E = "E"
)

var CARDINALS = [4]string{CARDINAL_N, CARDINAL_E, CARDINAL_W, CARDINAL_S}

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
