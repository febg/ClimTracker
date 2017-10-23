package main

import (
	"log"
	"net/http"

	"github.com/febg/Climbtracker/Go/api"
	//"./api"
)

func main() {

	c, err := api.NewControl(api.ControlConfig{
		LocalMySQL: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	router := api.StandardRouter(c)
	c.StartUpdateTimer(1)
	log.Println("[LOG] Listening on http://localhost:8080")
	log.Println("----------------------------------------")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
