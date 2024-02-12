package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardinalFromString(t *testing.T) {
	tests := []struct {
		c  string
		ce string
		ee string
	}{
		{"", "", "bad cardinat format"},
		{"AA", "", "bad cardinat format"},
		{"A", "", "bad cardinal value"},
		{"N", "N", ""},
	}

	for n, test := range tests {

		c, err := CreateCardinal(test.c)

		if test.ee != "" {
			if err == nil || err.Error() != test.ee {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.ce, c.Letter,
				fmt.Sprintf("test %d: got bad letter: %s and %s was expected", n, test.ce, c.Letter))
		}

	}
}

func TestTurnTo(t *testing.T) {
	tests := []struct {
		cc string
		nm string
		ev string
		e  string
	}{
		{"N", "L", "W", ""},
		{"N", "M", "", "bad movement value"},
		{"N", "R", "E", ""},
	}

	for n, test := range tests {

		c, err := CreateCardinal(test.cc)
		if err != nil {
			t.Fatalf("test %d: not error expected", n)
		}

		nc, err := c.TurnTo(test.nm)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.ev, nc.String(),
				fmt.Sprintf("test %d: turned to %s but turned to %s was expected", n, test.ev, c.Letter))
		}

	}

}
