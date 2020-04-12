package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	var username string = r.Form.Get("username")
	var userID string = r.Form.Get("userID")
	var userAddress string = r.Form.Get("userAddress")
	var userPhotoLocation string = r.Form.Get("userPhotoLocation")
	var createdBy string = r.Form.Get("createdBy")
	var createdDate string = r.Form.Get("createdDate")
	var updatedBy string = r.Form.Get("updatedBy")
	var updatedDate string = r.Form.Get("updatedDate")

	txid, err := app.Fabric.Adduser(username,userID, userAddress, userPhotoLocation, createdBy, createdDate, updatedBy, updatedDate)
	if err != nil {
		fmt.Println("adduser error", err)
		http.Error(w, "Unable to invoke hello in the blockchain1", 500)
	}
	data.TransactionId = txid
	data.Success = true
	data.Response = true
}
