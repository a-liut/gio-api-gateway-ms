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
	"flag"
	"fmt"
	"gio-api-gateway/pkg/api"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "port to be used")

	flag.Parse()

	checkVariables()

	log.Printf("Server started on port %d", *port)

	router := api.NewRouter()

	p := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(p, *router))
}

func checkVariables() {
	if deviceServiceHost := os.Getenv("DEVICE_SERVICE_HOST"); deviceServiceHost == "" {
		panic("DEVICE_SERVICE_HOST not set.")
	}
	if deviceServicePort := os.Getenv("DEVICE_SERVICE_PORT"); deviceServicePort == "" {
		panic("DEVICE_SERVICE_PORT not set.")
	}
}
