package infrastructure

import (
	"errors"
	"mowers/internal/domain"
	"testing"
)

func TestValidateInput(t *testing.T) {
	var tests = []struct {
		p  *domain.Plateau
		ml []domain.Mower
		c  string
		e  string
	}{
		{new(domain.Plateau), []domain.Mower{}, "", "input is empty"},
		/*
			{new(domain.Plateau), []domain.Mower{}, `55
			12N
			LMLMLMLMM
			33E
			MMRMMRMRRM`, ""},
			{new(domain.Plateau), []domain.Mower{}, "55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM", ""},
		*/
		{new(domain.Plateau), []domain.Mower{}, "55\n12N", "bad format input data"},
		{new(domain.Plateau), []domain.Mower{}, "55\n12N\nLMLMLMLMM\n33E", "bad format input data"},
		{new(domain.Plateau), []domain.Mower{}, "555\n12N\nLML", "bad coordinate format"},
	}

	for n, test := range tests {

		err := ValidateInput(test.c, test.p, test.ml)

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
