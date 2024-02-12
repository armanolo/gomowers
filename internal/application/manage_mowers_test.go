package application

import (
	"fmt"
	"mowers/internal/domain"
	"mowers/internal/infrastructure"
	"testing"
)

func TestManageMowers(t *testing.T) {

	p := new(domain.Plateau)
	var ml []*domain.Mower

	input := "55\n12N\nLM"

	ml, err := infrastructure.ValidateInput(input, p)

	if err != nil {
		t.Fatal(err)
	}

	out, err := ManageMowers(ml)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}
