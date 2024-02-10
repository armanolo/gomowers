package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMowerPositionFromString(t *testing.T) {
	tests := []struct {
		s  string
		cx int
		cy int
		c  string
		e  string
	}{
		{"11N", 1, 1, "N", ""},
		{"11", 1, 1, "N", "bad mower position format"},
		{"11A", 1, 1, "A", "bad cardinal value"},
		{"A1N", 1, 1, "N", "bad coodinate x"},
		{"1AN", 1, 1, "N", "bad coodinate y"},
	}

	for n, test := range tests {

		mp, err := MowerPositionFromString(test.s)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.cx, mp.Coordinates.X,
				fmt.Sprintf("test %d: got position x: %d and %d was expected", n, test.cx, mp.Coordinates.X))

			assert.Equal(t, test.cy, mp.Coordinates.Y,
				fmt.Sprintf("test %d: got position y: %d and %d was expected", n, test.cx, mp.Coordinates.Y))

			assert.Equal(t, test.c, mp.Cardinal.Letter,
				fmt.Sprintf("test %d: got position y: %s and %s was expected", n, test.c, mp.Cardinal.Letter))
		}

	}

}
