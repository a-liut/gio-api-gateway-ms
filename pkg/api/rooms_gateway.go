/*
 * API Gateway Service
 *
 * API Gateway Microservice for the Giò system.
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"encoding/json"
	"gio-api-gateway/pkg/model"
	"gio-api-gateway/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoomById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId := vars["roomId"]

	repo, _ := repository.NewRoomRepository(nil)
	room, err := repo.Get(roomId)

	if err != nil {
		errorHandler(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(room)
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	repo, _ := repository.NewRoomRepository(nil)
	rooms, err := repo.GetAll()

	if rooms == nil {
		errorHandler(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(rooms)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var roomData model.Room
	err := json.NewDecoder(r.Body).Decode(&roomData)
	if err != nil {
		errorHandler(w, http.StatusBadRequest, "Invalid data")
		return
	}

	repo, _ := repository.NewRoomRepository(nil)
	room, err := repo.Insert(&roomData)

	if err != nil {
		errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(room)
}
