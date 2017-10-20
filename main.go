package main

import (
	"log"
	"net/http"

	"github.com/febg/Climbtracker/api"
)

func main() {

	c, err := api.NewControl(api.ControlConfig{
		LocalMySQL: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	router := api.StandardRouter(c)
	log.Println("[INFO] Listening on http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
