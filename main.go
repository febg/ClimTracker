package main

import (
	"log"
	"net/http"
	"os"

	"github.com/febg/Climbtracker/api"
	//"./api"
)

func main() {
	log.Printf("[LOG] Booting up server..")
	c, err := api.NewControl(api.ControlConfig{
		LocalMySQL: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	os.Remove("qr.png")
	router := api.StandardRouter(c)
	log.Println("[LOG] Listening on http://localhost:8080")
	log.Println("----------------------------------------")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
