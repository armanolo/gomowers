package main

import (
	"flag"
	"fmt"
	"os"

	"mowers/internal/domain"
	"mowers/internal/infrastructure"
)

var it string
var ct string

func main() {
	fmt.Println("CLI: Start Mowers App")
	flag.StringVar(&it, "i", "text", "input type: text | file")
	flag.StringVar(&ct, "c", "", "coordaintes composition: check example.coord")
	flag.Parse()

	if it == "" || ct == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var p = new(domain.Plateau)
	var ml []domain.Mower = []domain.Mower{}
	var content string = getContent(it, ct)

	/*
		var content string

		if it == "file" {
			file, err := os.Open(ct) // For read access.
			if err != nil {
				fmt.Printf("not file found in: %s \n", ct)
				os.Exit(0)
			}
			data := make([]byte, 512)
			_, err = file.Read(data)
			if err != nil {
				fmt.Println("Error reading file")
				os.Exit(0)
			}
			content = string(data)
		} else {
			content = string(ct)
		}
	*/

	err := infrastructure.ValidateInput(content, p, ml)

	if err != nil {
		fmt.Println("Error")
		os.Exit(0)
	}

	fmt.Println("\n\nCLI: End Mowers App")
}

func getContent(it string, ct string) string {

	if it == "file" {
		file, err := os.Open(ct) // For read access.
		if err != nil {
			fmt.Printf("not file found in: %s \n", ct)
			os.Exit(0)
		}
		data := make([]byte, 512)
		_, err = file.Read(data)
		if err != nil {
			fmt.Println("Error reading file")
			os.Exit(0)
		}
		return string(data)
	} else {
		return string(ct)
	}
}
