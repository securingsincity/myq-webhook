package main

import (
	"log"
	"net/http"

	"github.com/joeshaw/myq"
)

func main() {
	log.Println("LET's Do this")
	&myq.Session{}
	http.ListenAndServe(":8080", nil)
}
