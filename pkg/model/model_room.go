/*
 * API Gateway Service
 *
 * API Gateway Microservice for the Giò system.
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

import (
	"fmt"
)

// A Room is a virtual place that may contains devices
type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (i *Room) Validate() (bool, error) {
	if i.Name == "" {
		return false, fmt.Errorf("invalid name")
	}

	return true, nil
}
