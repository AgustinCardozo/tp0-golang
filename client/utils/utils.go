package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/AgustinCardozo/tp0-golang/client/globals"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func LeerConsola() Paquete {
	// Leer de la consola
	paquete := Paquete{}
	for {
		reader := bufio.NewReader(os.Stdin)
		log.Println("Ingrese un mensaje (o escriba 'fin' para terminar):")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text) // Eliminamos \n y espacios
		log.Print(text)

		if text == "fin" {
			break
		}

		paquete.Valores = append(paquete.Valores, text)
	}
	return paquete
}

func GenerarYEnviarPaquete() {
	// Leemos y cargamos el paquete
	paquete := LeerConsola()

	slog.Info(fmt.Sprintf("paquete a enviar: %+v", paquete))
	// Enviamos el paquete
	EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		slog.Error(fmt.Sprintf("error codificando mensaje: %s", err.Error()))
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		slog.Error(fmt.Sprintf("error enviando mensaje a ip:%s puerto:%d", ip, puerto))
	}

	slog.Info(fmt.Sprintf("respuesta del servidor: %s", resp.Status))
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		slog.Info(fmt.Sprintf("error codificando mensajes: %s", err.Error()))
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		slog.Error(fmt.Sprintf("error enviando mensajes a ip:%s puerto:%d", ip, puerto))
	}

	slog.Info(fmt.Sprintf("respuesta del servidor: %s", resp.Status))
}

func ValidarServidor(ip string, puerto int) {
	url := fmt.Sprintf("http://%s:%d/handshake", ip, puerto)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error(fmt.Sprintf("error conectando a ip:%s puerto:%d", ip, puerto))
	}

	slog.Info(fmt.Sprintf("respuesta del servidor: %s", resp.Status))
}
