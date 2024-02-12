package domain

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlateau(t *testing.T) {
	var tests = []struct {
		c  string
		xe int
		ye int
		e  string
	}{
		{"55", 5, 5, ""},
		{"00", 5, 5, "coordinates with values under 1 for plateau"},
	}

	for n, test := range tests {

		coord, err := CoordinatesFromString(test.c)
		if err != nil {
			t.Error("wrong coordinates")
		}

		p, err := CreatePlateau(coord)

		if test.e != "" {
			if err == nil || err.Error() != errors.New(test.e).Error() {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.xe, p.Coordinates.X,
				fmt.Sprintf("test %d: got position x: %d and %d was expected", n, p.Coordinates.X, test.xe))

			assert.Equal(t, test.ye, p.Coordinates.Y,
				fmt.Sprintf("test %d: got position y: %d and %d was expected", n, p.Coordinates.Y, test.ye))
		}
	}

}
