package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/AgustinCardozo/tp0-golang/client/globals"
	"github.com/AgustinCardozo/tp0-golang/client/utils"
	"github.com/AgustinCardozo/tp0-golang/lib/utils"
)

func main() {
	commons.ConfigurarLogger("../client/logs/cliente.log", "info")
	// loggear "Hola soy un log" usando la biblioteca log
	slog.Info("Hola soy un log")
	commons.Iniciar("../client/configs/config.json", &globals.ClientConfig)
	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatal("Error al cargar la configuracion")
	}

	// loggeamos el valor de la config
	slog.Info(fmt.Sprintf("IP: %s, Puerto: %d\n", globals.ClientConfig.Ip, globals.ClientConfig.Puerto))
	slog.Info(fmt.Sprintf("Mensaje: %s", globals.ClientConfig.Mensaje))

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él
	utils.ValidarServidor(globals.ClientConfig.Ip, globals.ClientConfig.Puerto)

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	for {
		valor := utils.LeerConsola()
		// generamos un paquete y lo enviamos al servidor
		utils.GenerarYEnviarPaquete(valor)
	}
}
