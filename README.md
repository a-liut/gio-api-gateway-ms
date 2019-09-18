# gio-api-gateway-ms
The API Gateway component is a microservice that provides an entry point for the whole Giò Plants platform.
It exposes all the functionalities of the system to clients and the Frontend microservice.

## Run

### Docker

The simplest way to execute the server is to use Docker.
Follow these steps to run the container:

```bash
docker build -t gio-api-gateway-ms:latest .
docker run -it -p 8080:8080 gio-api-gateway-ms:latest
```

### Standalone

gio-api-gateway-ms is developed as a Go module.

```bash
go build -o apigateway cmd/apigateway/main.go

./apigateway
```

## Entities

### Device

A Device represents a physical device registered in the system. Each device is identified by an ID.
A Device must be related to an existing room.

A Device has the followings fields:

- id: *string* - The UUID of the object
- name: *string* - The name of the device.
- mac: *string* -  The MAC address of the device. Must be unique.
- room *string* - The UUID of the room in which the device is placed

Example:

```json
{
  "id": "20335c53-6929-4b9d-900e-5548d99fd4e1",
  "name": "device1",
  "mac": "f6:f1:bb:06:31:71",
  "room": "5c8fa383-f9b7-4d80-9572-3276da1ab67d"
}
```

### Room

A Room is a (possibly empty) collection of devices.

A Room has the following fields:

- id: *string* - The UUID of the object
- name: *string* - The name of the room.

Example:

```json
{
  "id": "5c8fa383-f9b7-4d80-9572-3276da1ab67d",
  "name": "Room1"
}
```

## Endpoints

- ### /rooms

    **GET**: return all registered rooms.

    **POST**: register a new room.
    
    Example body:
```json
{
  "name": "Room1"
}
```

- ### /rooms/{roomId}

    **GET**: return the specified room.
    
- ### /rooms/{roomId}/devices

    **GET**: return all registered devices belonging a specific room.

    **POST**: register a new device in a specific room.
    
    Example body:
```json
{
  "mac": "f6:f1:bb:06:31:71",
  "name": "device1"
}
```

- ### /rooms/{roomId}/devices/{deviceId}

    **GET**: return the specified device of a specific room.

- ### /rooms/{roomId}/devices/{deviceId}/readings

    **GET**: return all readings of the specified devices in a specific room.
    
    Optional query parameters
    
    - limit(n): limit the results obtained to the last n entries stored
    
    - name(s): filter the entries to those who have `s` as name
    
    **POST**: register a new reading for the specified device.
    
    Example body:
```json
{
    "name": "temperature",
    "value": "22",
    "unit": "°C",
    "creation_timestamp": "yyyyyyy"
}
```

- ### /rooms/{roomId}/devices/{deviceId}/actions/{actionName}

    **POST**: Triggers the specified action on a specific device. It allows sending a value with the request.
    
    Example body:
```json
{
    "value": "22"
}
```

    WARNING: always specify a body when requesting 'watering' action with a positive number as value.
