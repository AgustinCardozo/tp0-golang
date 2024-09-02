package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func RecibirPaquetes(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var paquete Paquete
	err := decoder.Decode(&paquete)
	if err != nil {
		slog.Error(fmt.Sprintf("error al decodificar mensaje: %s\n", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error al decodificar mensaje"))
		return
	}

	slog.Info("me llego un paquete de un cliente")
	slog.Info(fmt.Sprintf("%+v\n", paquete))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func RecibirMensaje(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var mensaje Mensaje
	err := decoder.Decode(&mensaje)
	if err != nil {
		slog.Error(fmt.Sprintf("Error al decodificar mensaje: %s\n", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al decodificar mensaje"))
		return
	}

	slog.Info("Me llego un mensaje de un cliente")
	slog.Info(fmt.Sprintf("%+v\n", mensaje))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Handshake(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
