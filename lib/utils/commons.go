package commons

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
)

type Cliente struct {
	Ip      string `json:"ip"`
	Puerto  int    `json:"puerto"`
	Mensaje string `json:"mensaje"`
}

func Decode(filePath string, configJson interface{}) error {
	// Abre el archivo
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// Cierra el archivo una vez que la función termina (ejecuta el return)
	defer configFile.Close()

	// Decodifica la info del json en la variable config
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(configJson)
	if err != nil {
		return err
	}

	return nil
}

func Iniciar(filePath string, configJson interface{}) {
	err := Decode(filePath, &configJson)

	if err != nil {
		fmt.Println("Error al iniciar configuración: ", err)
	}
}

func ConfigurarLogger(logPath string, logLevel string) {
	file, err := os.Create(logPath)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	switch logLevel {
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case "warn":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "error":
		slog.SetLogLoggerLevel(slog.LevelError)
	default:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}
}
