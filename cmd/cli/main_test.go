package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	tests := []struct {
		args []string
		re   string
	}{
		{
			[]string{"cmd", "-i", "", "-c", ""},
			"",
		},
		{
			[]string{"cmd", "-i", "text", "-c", "example.coord"},
			"error: bad format input data",
		},
		{
			[]string{"cmd", "-i", "text", "-c", "55"},
			"error: bad format input data",
		},
		{
			[]string{"cmd", "-i", "file", "-c", "example.coord"},
			"error: not file found in: example.coord",
		},
		{
			[]string{"cmd", "-i", "text", "-c", "55\n11N\nMMM\n"},
			"14N",
		},
		{
			[]string{"cmd", "-i", "file", "-c", "example_test.coord"},
			"13N",
		},
	}

	for n, test := range tests {
		t.Run(fmt.Sprintf("test %d", n), func(t *testing.T) {
			old := os.Stdout
			os.Args = test.args
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			out, _ := io.ReadAll(r)
			r.Close()
			result := string(out)

			assert.Equal(t, test.re, result,
				fmt.Sprintf("test %d: got %q and %q was expected", n, test.re, result))

			os.Stdout = old

			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		})
	}

}
