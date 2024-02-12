package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMower(t *testing.T) {

	p, _ := CreatePlateau(Coordinates{X: 5, Y: 5})

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

		m, err := CreateMower(&p, mp, path)

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

func TestMovementProcess(t *testing.T) {

	p, _ := CreatePlateau(Coordinates{X: 5, Y: 5})

	tests := []struct {
		mp   string
		path string
		ep   string
		e    string
	}{
		{"11N", "M", "12N", ""},
		{"22N", "ML", "23W", ""},
		{"33N", "R", "33E", ""},
		{"10W", "M", "00W", ""},
		{"02S", "M", "01S", ""},
		{"03E", "M", "13E", ""},
		{"31S", "M", "30S", ""},
	}

	for n, test := range tests {

		mp, err := MowerPositionFromString(test.mp)
		if err != nil {
			t.Fatalf("test %d: wrong mower: %q", n, err.Error())
		}
		path, err := PathFromString(test.path)
		if err != nil {
			t.Fatalf("test %d: wrong mower: %q", n, err.Error())
		}
		m, err := CreateMower(&p, mp, path)

		if err != nil {
			t.Fatalf("test %d: wrong mower: %q", n, err.Error())
		}

		m.MovementProcess()

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}
		}

		assert.Equal(t, test.ep, m.MowerPosition.String(),
			fmt.Sprintf("test %d: got position y: %q and %q was expected", n, m.MowerPosition.String(), test.ep))

	}

}
