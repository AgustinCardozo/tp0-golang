package main

import (
	"net/http"

	"github.com/AgustinCardozo/tp0-golang/lib/utils"
	"github.com/AgustinCardozo/tp0-golang/server/utils"
)

func main() {
	commons.ConfigurarLogger("../server/logs/server.log", "info")
	mux := http.NewServeMux()

	mux.HandleFunc("/handshake", utils.Handshake)
	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	// panic("no implementado!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
