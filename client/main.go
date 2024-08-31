package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatal("Error al cargar la configuracion")
	}

	// loggeamos el valor de la config
	log.Printf("IP: %s, Puerto: %d\n", globals.ClientConfig.Ip, globals.ClientConfig.Puerto)
	log.Printf("Mensaje: %s\n", globals.ClientConfig.Mensaje)

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
