package main

import (
	"log"
	"os"
	"testing"
)

var app Application

func TestMain(m *testing.M) {
	log.Println("MMM doing stuff in api side BEFORE the tests!")
	log.Println("MMM doing stuff in api side AFTER the tests!")
	os.Exit(m.Run())
}
