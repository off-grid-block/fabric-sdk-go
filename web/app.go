package web

import (
	"fmt"
	"net/http"

	"github.com/gosdk-example/web/controllers"
)

func Serve(app *controllers.Application) {
	http.HandleFunc("/application", app.UserHandler)
	http.HandleFunc("/query", app.QueryHandler)
	http.HandleFunc("/invoke", app.InvokeHandler)
	http.HandleFunc("/history", app.HistoryHandler)
	http.HandleFunc("/adduser", app.AddUserHandler)
	http.HandleFunc("/updateuser", app.UpdateUserHandler)
	http.HandleFunc("/userquery", app.UserQueryHandler)
	http.HandleFunc("/initVote", app.InitVoteHandler)
	http.HandleFunc("/getVote", app.GetVoteHandler)

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}
