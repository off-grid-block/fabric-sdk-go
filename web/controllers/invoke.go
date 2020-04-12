package controllers

import (
	"net/http"
)

func (app *Application) InvokeHandler(w http.ResponseWriter, r *http.Request) {
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
	var fcn string = r.Form.Get("fcn")
	var key string = r.Form.Get("key")
	var value string = r.Form.Get("value")

	txid, err := app.Fabric.Invoke(fcn, key, value)
	if err != nil {
		http.Error(w, "Unable to invoke hello in the blockchain1", 500)
	}
	data.TransactionId = txid
	data.Success = true
	data.Response = true
}
