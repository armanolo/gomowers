package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"mowers/internal/application"
	"mowers/internal/domain"
	"mowers/internal/infrastructure"
)

var it string
var ct string

func main() {
	flag.StringVar(&it, "i", "text", "input type: text | file")
	flag.StringVar(&ct, "c", "", "coordaintes composition: check example.coord")
	flag.Parse()

	if it == "" || ct == "" {
		flag.PrintDefaults()
		return
	}

	res, err := process()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Print(res)
	}

}

func process() (string, error) {

	var p = new(domain.Plateau)
	var ml = new([]domain.Mower)
	content, err := getContent(it, ct)

	if err != nil {
		return "", err
	}

	err = infrastructure.ValidateInput(content, p, ml)
	if err != nil {
		return "", err
	}

	out, err := application.ManageMowers(ml)
	if err != nil {
		return "", err
	}

	return out, nil

}

func getContent(it string, ct string) (string, error) {

	if it == "file" {
		file, err := os.Open(ct)
		if err != nil {
			return "", fmt.Errorf("not file found in: %s", ct)
		}
		data := make([]byte, 512)
		_, err = file.Read(data)
		data = bytes.Trim(data, "\x00")
		if err != nil {
			return "", errors.New("error reading file")
		}
		return strings.Trim(string(data), "\n "), nil
	} else {
		return strings.Trim(string(ct), "\n "), nil
	}
}
