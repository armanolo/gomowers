package application

import (
	"fmt"
	"mowers/internal/domain"
	"mowers/internal/infrastructure"
	"testing"
)

func TestManageMowers(t *testing.T) {

	p := new(domain.Plateau)
	ml := new([]domain.Mower)

	input := "55\n12N\nLM"

	err := infrastructure.ValidateInput(input, p, ml)

	if err != nil {
		t.Fatal(err)
	}

	out, err := ManageMowers(ml)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}
