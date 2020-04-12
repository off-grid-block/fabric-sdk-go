package controllers

import (
	"net/http"
)

func (app *Application) InitVoteHandler(w http.ResponseWriter, r *http.Request) {
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
	var PollID string = r.Form.Get("PollID")
	var VoterID string = r.Form.Get("VoterID")
	var VoterSex string = r.Form.Get("VoterSex")
	var VoterAge string = r.Form.Get("VoterAge")
	var Salt string = r.Form.Get("Salt")
	var VoteHash string = r.Form.Get("VoteHash")

	txid, err := app.Fabric.InitVoteHandler(username,PollID, VoterID, VoterSex, VoterAge, Salt, VoteHash)
	if err != nil {
		http.Error(w, "Unable to invoke hello in the blockchain1", 500)
	}
	data.TransactionId = txid
	data.Success = true
	data.Response = true
}
