package main

import (
	"github.com/coentie/filesync-server/bootstrap"
	"github.com/coentie/filesync-server/router"
	"net/http"
	"os"
)

func main() {
	err := bootstrap.Boostrap()

	if err != nil {
		panic(err)
	}

	router := router.Router()

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, router)
}
