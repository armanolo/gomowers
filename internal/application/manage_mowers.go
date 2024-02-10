package application

import (
	"mowers/internal/domain"
	"strings"
)

func ManageMowers(mowers *[]domain.Mower) (string, error) {

	var b = new(strings.Builder)

	for _, mower := range *mowers {

		lp, err := mower.MovementProcess()

		if err != nil {
			return "", nil
		}

		b.WriteString(lp)
		b.WriteString("\n")
	}

	return b.String(), nil
}
