package infrastructure

import (
	"errors"
	"fmt"
	"mowers/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInputCheckObjects(t *testing.T) {

	p := new(domain.Plateau)
	ml := new([]domain.Mower)
	input := "55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM"

	err := ValidateInput(input, p, ml)

	if err != nil {
		t.Fatalf("Error: %q", err)
	}

	nmp := len(p.MowerPositions)
	ne := 2

	assert.Equal(t, ne, nmp,
		fmt.Sprintf("%d mowers in plateau expected but we got %d", ne, nmp))

}

func TestValidateInput(t *testing.T) {

	var tests = []struct {
		p  *domain.Plateau
		ml *[]domain.Mower
		c  string
		e  string
	}{
		{new(domain.Plateau), new([]domain.Mower), "", "input is empty"},
		{new(domain.Plateau), new([]domain.Mower), `55
12N
LMLMLMLMM
33E
MMRMMRMRRM`, ""},
		{new(domain.Plateau), new([]domain.Mower), "55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM", ""},
		{new(domain.Plateau), new([]domain.Mower), "55\n12N", "bad format input data"},
		{new(domain.Plateau), new([]domain.Mower), "55\n12N\nLMLMLMLMM\n33E", "bad format input data"},
		{new(domain.Plateau), new([]domain.Mower), "555\n12N\nLML", "bad coordinate format"},
		{new(domain.Plateau), new([]domain.Mower), "55\n12A\nLML", "bad cardinal value"},
		{new(domain.Plateau), new([]domain.Mower), "55\n12N\nA", "bad path value"},
		{new(domain.Plateau), new([]domain.Mower), "55\n12N\nL\n12N\nL", "there is another mower in the position: 1:2"},
	}

	for n, test := range tests {

		err := ValidateInput(test.c, test.p, test.ml)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}
		}
	}
}

func TestValidatePlateu(t *testing.T) {
	var tests = []struct {
		p  *domain.Plateau
		c  string
		xe int
		ye int
		e  string
	}{
		{new(domain.Plateau), "55", 5, 5, ""},
		{new(domain.Plateau), "551", 0, 0, "bad coordinate format"},
	}

	for n, test := range tests {

		err := ValidatePlateu(test.c, test.p)

		if test.e != "" {
			if err == nil || err.Error() != errors.New(test.e).Error() {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

		}
	}

}
