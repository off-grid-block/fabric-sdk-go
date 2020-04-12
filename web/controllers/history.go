package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Application) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.URL.Query()["username"]
	query, ok := r.URL.Query()["query"]
	if !ok || len(query[0]) < 1 {
		log.Println("Url Param 'query' is missing")
		fmt.Fprint(w, "Url Param 'query' is missing")
		return
	}
	if query[0] == "txn" {
		txnid, ok := r.URL.Query()["txnid"]
		fmt.Println("here")
		if !ok || len(txnid[0]) < 1 {
			log.Println("Url Param 'txnid' is missing")
			fmt.Fprint(w, "Url Param 'txnid' is missing")
			return
		}
		value, err := app.Fabric.GettxbyID(username[0],txnid[0])
		if err != nil {
			http.Error(w, "Unable to get history", 500)
		}

		fmt.Fprint(w, value)
	}
	if query[0] == "blockbytxnid" {
		txnid, ok := r.URL.Query()["txnid"]
		fmt.Println("here2")

		if !ok || len(txnid[0]) < 1 {
			log.Println("Url Param 'txnid' is missing")
			fmt.Fprint(w, "Url Param 'txnid' is missing")
			return
		}
		value, err := app.Fabric.GetBlockbytxID(txnid[0])
		if err != nil {
			http.Error(w, "Unable to get history", 500)
		}

		fmt.Fprint(w, value)
	}
	if query[0] == "blockbyhash" {
		hash, ok := r.URL.Query()["hash"]
		fmt.Println("here2")

		if !ok || len(hash[0]) < 1 {
			log.Println("Url Param 'hash' is missing")
			fmt.Fprint(w, "Url Param 'hash' is missing")
			return
		}
		value, err := app.Fabric.GetBlockbyHash(hash[0])
		if err != nil {
			http.Error(w, "Unable to get history", 500)
		}

		fmt.Fprint(w, value)
	}
	if query[0] == "blockbyid" {
		blockid, ok := r.URL.Query()["blockid"]
		fmt.Println("here2")

		if !ok || len(blockid[0]) < 1 {
			log.Println("Url Param 'blockid' is missing")
			fmt.Fprint(w, "Url Param 'blockid' is missing")
			return
		}
		value, err := app.Fabric.GetBlockbyID(blockid[0])
		if err != nil {
			http.Error(w, "Unable to get history", 500)
		}

		fmt.Fprint(w, value)
	}
}
