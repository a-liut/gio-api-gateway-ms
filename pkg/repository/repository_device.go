/*
 * API Gateway Service
 *
 * API Gateway Microservice for the Giò system.
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gio-api-gateway/pkg/config"
	"gio-api-gateway/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type DeviceRepository struct {
	devicesServiceUrl string
}

func (r *DeviceRepository) Get(roomId string, id string) (*model.Device, error) {
	u := fmt.Sprintf("%s/rooms/%sdevices/%s", r.devicesServiceUrl, roomId, id)

	resp, err := http.Get(u)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("device %s not found", id)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error while getting data for device %s", id)
	}

	var d model.Device
	err = json.NewDecoder(resp.Body).Decode(&d)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DeviceRepository) GetReadings(roomId string, id string) ([]*model.Reading, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices/%s/readings", r.devicesServiceUrl, roomId, id)

	resp, err := http.Get(u)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("device %s not found", id)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error while getting data for device %s", id)
	}

	var rs []*model.Reading
	err = json.NewDecoder(resp.Body).Decode(&rs)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (r *DeviceRepository) GetAll(roomId string) ([]*model.Device, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices", r.devicesServiceUrl, roomId)

	resp, err := http.Get(u)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error while getting data for devices")
	}

	var d []*model.Device
	err = json.NewDecoder(resp.Body).Decode(&d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (r *DeviceRepository) Insert(roomId string, device *model.Device) (*model.Device, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices", r.devicesServiceUrl, roomId)

	b, err := json.Marshal(device)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(u, "application/json", bytes.NewBuffer(b))
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error while performing the operation: %d - %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	var d model.Device
	err = json.NewDecoder(resp.Body).Decode(&d)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DeviceRepository) InsertReading(roomId string, deviceId string, readingData *model.Reading) (*model.Reading, error) {
	u := fmt.Sprintf("%s/rooms/%s/devices/%s/readings", r.devicesServiceUrl, roomId, deviceId)

	b, err := json.Marshal(readingData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(u, "application/json", bytes.NewBuffer(b))
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error while performing the operation: %d - %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	var rea model.Reading
	err = json.NewDecoder(resp.Body).Decode(&rea)

	if err != nil {
		return nil, err
	}

	return &rea, nil
}

var devicesRepository *DeviceRepository

func NewDeviceRepository(serviceConfig *config.DeviceServiceConfig) (*DeviceRepository, error) {
	if devicesRepository == nil {
		u := fmt.Sprintf("http://%s:%d", serviceConfig.Host, serviceConfig.Port)
		log.Printf("DeviceService URL: %s\n", u)

		serviceUrl, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		devicesRepository = &DeviceRepository{serviceUrl.String()}
	}

	return devicesRepository, nil
}
