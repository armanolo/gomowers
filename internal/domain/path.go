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

	var n int = len(movements)

	m := make([]string, n)

	lu := strings.ToUpper(movements)

	for p, a := range strings.Split(lu, "") {

		if strings.Contains(MOVEMENTS, a) {
			m[p] = a
		} else {
			return Path{}, errors.New("bad path value")
		}
	}

	return Path{Movement: m}, nil
}
