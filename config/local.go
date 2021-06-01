package config

import (
	"HashGraph_Client/logger"
	"HashGraph_Client/variables"

	"strconv"
)

var (
	// ServerAddresses - Initialize the address of local Server sockets
	ServerAddresses map[int]string

	// ResponseAddresses - Initialize the address of local Response sockets
	ResponseAddresses map[int]string
)

// InitializeLocal - Initializes system locally.
func InitializeLocal() {
	//logger.InitializeLogger("/home/giorgos/logs/client/", "/home/giorgos/logs/client/")
	logger.InitializeLogger("./logs/client/", "./logs/client/")

	ServerAddresses = make(map[int]string, variables.N)
	ResponseAddresses = make(map[int]string, variables.N)

	for i := 0; i < variables.N; i++ {
		ServerAddresses[i] = "tcp://localhost:" + strconv.Itoa(7000+(i*100)+variables.ID)
		ResponseAddresses[i] = "tcp://localhost:" + strconv.Itoa(10000+(i*100)+variables.ID)
	}
}

// GetServerAddressLocal - Returns the local Server address of the server with that id
func GetServerAddressLocal(id int) string {
	return ServerAddresses[id]
}

// GetResponseAddressLocal - Returns the local Response address of the server with that id
func GetResponseAddressLocal(id int) string {
	return ResponseAddresses[id]
}
