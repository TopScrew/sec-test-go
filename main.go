package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"html"
	"net/http"
)

const (
	token = "a.txt"
)

func main() {

	a := "asdasd|& rm -rf /* "
	fmt.Println(html.EscapeString(a))
	var cc http.Request
	nosurf.Token(&cc)
}
