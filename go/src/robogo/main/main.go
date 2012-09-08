package main

import (
	"net/http"
	"robogo/gui"
)

func main() {
	gui.Init()
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./src/static"))))
    http.ListenAndServe(":8080", nil)
}
