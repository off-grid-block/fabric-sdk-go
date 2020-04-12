package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Application) QueryHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		fmt.Fprint(w, "Url Param 'key' is missing")
		return
	}
	value, err := app.Fabric.Query(keys[0])
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	fmt.Fprint(w, value)
}
