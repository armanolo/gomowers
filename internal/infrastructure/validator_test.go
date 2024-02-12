package infrastructure

import (
	"fmt"
	"mowers/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInputCheckObjects(t *testing.T) {

	p := new(domain.Plateau)
	var ml []*domain.Mower
	input := "55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM"

	ml, err := ValidateInput(input, p)

	if err != nil {
		t.Fatalf("Error: %q", err)
	}

	nmp := len(p.Mowers)
	ne := 2

	assert.Equal(t, ne, len(ml),
		fmt.Sprintf("%d mowers expected to be created but we got %d", ne, len(ml)))
	assert.Equal(t, ne, nmp,
		fmt.Sprintf("%d mowers in plateau expected but we got %d", ne, nmp))

}

func TestValidateInput(t *testing.T) {

	var tests = []struct {
		p         *domain.Plateau
		mowersexp int
		c         string
		e         string
	}{
		{new(domain.Plateau), 0, "", "input is empty"},
		{new(domain.Plateau), 2, `55
12N
LMLMLMLMM
33E
MMRMMRMRRM`, ""},
		{new(domain.Plateau), 2, "55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM", ""},
		{new(domain.Plateau), 0, "55\n12N", "bad format input data"},
		{new(domain.Plateau), 0, "55\n12N\nLMLMLMLMM\n33E", "bad format input data"},
		{new(domain.Plateau), 0, "555\n12N\nLML", "bad coordinate format"},
		{new(domain.Plateau), 0, "55\n12A\nLML", "bad cardinal value"},
		{new(domain.Plateau), 0, "55\n12N\nA", "bad path value"},
		{new(domain.Plateau), 0, "55\n12N\nL\n12N\nL", "there is another mower in the position: 1:2"},
	}

	for n, test := range tests {

		ml, err := ValidateInput(test.c, test.p)

		if test.e != "" {
			if err == nil || err.Error() != test.e {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}
		}

		if len(ml) != test.mowersexp {
			t.Errorf("Test %d: error: %s", n, err)
		}
	}
}
