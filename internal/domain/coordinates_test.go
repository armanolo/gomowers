package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinatesFromString(t *testing.T) {
	var tests = []struct {
		c  string
		xe int
		ye int
		e  string
	}{
		{"55", 5, 5, ""},
		{"-1-2", 0, 0, "bad coordinate format"},
		{"1", 0, 0, "bad coordinate format"},
		{"133", 0, 0, "bad coordinate format"},
		{"A3", 0, 0, "bad coodinate x"},
		{"3A", 0, 0, "bad coodinate y"},
	}

	for n, test := range tests {
		c, err := CoordinatesFromString(test.c)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.xe, c.X,
				fmt.Sprintf("test %d: got position x: %d and %d was expected", n, test.xe, c.X))

			assert.Equal(t, test.xe, c.Y,
				fmt.Sprintf("test %d: got position y: %d and %d was expected", n, test.xe, c.Y))
		}
	}
}

func TestCreateCoordinates(t *testing.T) {
	var tests = []struct {
		xe int
		ye int
		e  string
	}{
		{5, 5, ""},
		{-1, 0, "bad coodinate x"},
		{1, 10, "bad coodinate y"},
	}

	for n, test := range tests {
		c, err := CreateCoordinates(test.xe, test.ye)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.xe, c.X,
				fmt.Sprintf("test %d: got position x: %d and %d was expected", n, test.xe, c.X))

			assert.Equal(t, test.xe, c.Y,
				fmt.Sprintf("test %d: got position y: %d and %d was expected", n, test.xe, c.Y))
		}
	}
}
