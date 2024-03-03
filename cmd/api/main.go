package main

import (
	"flag"
	"fmt"
	"net/http"
)

type Application struct {
	Port string
}

func main() {
	app := Application{}
	flag.StringVar(&app.Port, "p", "3333", "input type: text | file")
	flag.Parse()
	http.ListenAndServe(fmt.Sprintf(":%s", app.Port), app.routes())
}
