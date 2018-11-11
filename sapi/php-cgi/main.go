package main

import (
	"log"
	"net/http/cgi"

	"github.com/MagicalTux/gophp/core"
	_ "github.com/MagicalTux/gophp/ext/standard"
)

func main() {
	p := core.NewProcess()
	err := cgi.Serve(p.Handler("."))
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
