package domain

import (
	"errors"
	"strings"
)

const (
	PATH_L = "L"
	PATH_R = "R"
	PATH_M = "M"
)

var MOVEMENTS string = PATH_L + PATH_R + PATH_M

type Path struct {
	Movement []string
}

func PathFromString(movements string) (Path, error) {

	if len(movements) < 1 {
		return Path{}, errors.New("bad path format")
	}
	lu := strings.ToUpper(movements)
	arr := strings.Split(lu, "")

	return CreatePath(arr)
}

func CreatePath(movements []string) (Path, error) {
	var n int = len(movements)
	m := make([]string, n)

	for p, a := range movements {
		if IsMovementValue(a) {
			m[p] = a
		} else {
			return Path{}, errors.New("bad path value")
		}
	}
	return Path{Movement: m}, nil
}

func IsMovementValue(movement string) bool {
	return strings.Contains(MOVEMENTS, movement)
}
