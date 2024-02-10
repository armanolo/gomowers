package domain

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPathFromString(t *testing.T) {
	tests := []struct {
		c  string
		ce []string
		ee string
	}{
		{"", []string{}, "bad path format"},
		{"LALRLMLMM", []string{}, "bad path value"},
		{"LMLMLMLMM", []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"}, ""},
	}

	for n, test := range tests {

		p, err := PathFromString(test.c)

		if test.ee != "" {
			if err == nil || err.Error() != errors.New(test.ee).Error() {
				t.Errorf("Test %d: error: %s", n, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test %d: error: %q", n, err)
			}

			assert.Equal(t, test.ce, p.Movement,
				fmt.Sprintf("test %d: got bad letter: %q and %q was expected", n, test.ce, p.Movement))
		}

	}

}
