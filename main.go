package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joeshaw/myq"
)

func main() {
	s := &myq.Session{}
	s.Username = os.Getenv("MYQ_EMAIL")
	s.Password = os.Getenv("MYQ_PASSWORD")
	deviceID := os.Getenv("MYQ_DEVICEID")
	port := os.Getenv("MYQ_WEBHOOK_PORT")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("We're running")
	})
	r.Post("/{id}/open", func(w http.ResponseWriter, r *http.Request) {
		e := s.Login()
		if e != nil {
			log.Printf("Error logging in: %s", e)
			return
		}

		s.SetDoorState(deviceID, myq.StateOpen)
		w.Write([]byte(fmt.Sprintf("opening:%s", deviceID)))
		log.Printf("opening device id %s", deviceID)
	})
	r.Post("/{id}/close", func(w http.ResponseWriter, r *http.Request) {
		e := s.Login()
		if e != nil {
			log.Printf("Error logging in: %s", e)
			return
		}
		s.SetDoorState(deviceID, myq.ActionClose)
		w.Write([]byte(fmt.Sprintf("closing:%s", deviceID)))
		log.Printf("closing device id %s", deviceID)
	})
	log.Printf("Bound to port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
