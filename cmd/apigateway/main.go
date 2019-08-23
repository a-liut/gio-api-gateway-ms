/*
 * API Gateway Service
 *
 * API Gateway Microservice for the Giò system.
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gio-api-gateway/pkg/api"
	"gio-api-gateway/pkg/config"
	"gio-api-gateway/pkg/repository"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "port to be used")
	configPath := flag.String("config", "config.json", "Configuration file")

	flag.Parse()

	if err := loadConfig(*configPath); err != nil {
		panic(err)
	}

	log.Printf("Server started on port %d", *port)

	router := api.NewRouter()

	p := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(p, router))
}

func loadConfig(path string) error {
	file, _ := os.Open(path)
	defer file.Close()

	var conf config.Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return err
	}

	if _, err := repository.NewDeviceRepository(conf.DeviceServiceConfig); err != nil {
		return err
	}

	return nil
}
