package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Application) GetVoteHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.URL.Query()["username"]
	pid, ok := r.URL.Query()["pollID"]
	vid, ok := r.URL.Query()["voterID"]

	fmt.Println("Username",username,pid)

	if !ok || len(username[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		fmt.Fprint(w, "Url Param 'key' is missing")
		return
	}

	value, err := app.Fabric.GetVoteSDK(username[0],pid[0],vid[0])
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	fmt.Fprint(w, value)
}
