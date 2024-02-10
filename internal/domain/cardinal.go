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
	Letter string
}

func CreateCardinal(character string) (Cardinal, error) {
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

	return Cardinal{Letter: l}, nil
}

func (c Cardinal) TurnTo(movement string) (Cardinal, error) {
	var offset int
	if movement == PATH_R {
		offset = 1
	} else {
		offset = -1
	}

	cardinals := []string{CARDINAL_N, CARDINAL_E, CARDINAL_S, CARDINAL_W}
	currentOrdinal := getIndex(c.Letter, cardinals)
	newOrdinal := (currentOrdinal + offset + len(cardinals)) % len(cardinals)

	return CreateCardinal(cardinals[newOrdinal])
}

func getIndex(cardinal string, cardinals []string) int {
	for i, c := range cardinals {
		if c == cardinal {
			return i
		}
	}
	return -1
}

func (c Cardinal) String() string {
	return c.Letter
}
