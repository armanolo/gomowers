package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMower(t *testing.T) {

	p := new(Plateau)

	tests := []struct {
		mpx  int
		mpy  int
		path []string
		e    string
	}{
		{5, 5, []string{"M", "L", "R"}, ""},
		{1, 1, []string{}, ""},
		{1, 1, []string{}, "there is another mower in the position: 1:1"},
	}

	for n, test := range tests {
		coord, _ := CreateCoordinates(test.mpx, test.mpy)
		card, _ := CreateCardinal("N")
		mp, _ := CreateMowerPosition(coord, card)
		path, _ := CreatePath(test.path)

		m, err := CreateMower(p, mp, path)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.mpx, m.MowerPosition.Coordinates.X,
				fmt.Sprintf("test %d: got position x: %d and %d was expected", n, test.mpx, m.MowerPosition.Coordinates.X))

			assert.Equal(t, test.mpy, m.MowerPosition.Coordinates.Y,
				fmt.Sprintf("test %d: got position y: %d and %d was expected", n, test.mpx, m.MowerPosition.Coordinates.Y))

			assert.Equal(t, test.path, m.Path.Movement,
				fmt.Sprintf("test %d: got position y: %q and %q was expected", n, test.path, m.Path.Movement))

		}

	}

}
