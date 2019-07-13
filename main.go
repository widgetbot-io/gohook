package main

import (
	"fmt"
	"git.deploys.io/disweb/gohook/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	embed := utils.NewEmbed()
	embed.SetTitle("Title")
	embed.SetDescription("Description")

	router := mux.NewRouter()

	setupRoutes(router)
	_ = http.ListenAndServe(":8443", nil)
}

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/api/hook/{ID}/{Secret}/{Provider}/{Options}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		fmt.Print(params)

	})

	http.Handle("/", router)
}
