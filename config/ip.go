package config

import (
	"HashGraph_Client/logger"
	"HashGraph_Client/variables"

	"strconv"
)

var address0 = []string{}

var address1 = []string{}

var (
	// ServerAddressesIP - Initialize the address of IP Server sockets
	ServerAddressesIP map[int]string

	// ResponseAddressesIP - Initialize the address of IP Response sockets
	ResponseAddressesIP map[int]string
)

// InitializeIP - Initializes system with ips.
func InitializeIP() {
	logger.InitializeLogger("./logs/client/", "./logs/client/")
	clients := variables.Clients
	ServerAddressesIP = make(map[int]string, variables.N)
	ResponseAddressesIP = make(map[int]string, variables.N)

	address := address1
	if variables.N == 4 {
		address = address0
	}

	for i := 0; i < variables.N; i++ {
		ad := i
		if i >= len(address) {
			ad = (i % 2) + 2
		}

		ServerAddressesIP[i] = "tcp://" + address[ad] + ":" + strconv.Itoa(27250+variables.ID+(i*clients))
		ResponseAddressesIP[i] = "tcp://" + address[ad] + ":" + strconv.Itoa(27625+variables.ID+(i*clients))
	}
}

// GetServerAddress - Returns the IP Server address of the server with that id
func GetServerAddress(id int) string {
	return ServerAddressesIP[id]
}

// GetResponseAddress - Returns the IP Response address of the server with that id
func GetResponseAddress(id int) string {
	return ResponseAddressesIP[id]
}

// var addresses = []string{
// }

// var (
// 	// ServerAddressesIP - Initialize the address of IP Server sockets
// 	ServerAddressesIP map[int]string

// 	// ResponseAddressesIP - Initialize the address of IP Response sockets
// 	ResponseAddressesIP map[int]string
// )

// // InitializeIP - Initializes system with ips.
// func InitializeIP() {
// 	ServerAddressesIP = make(map[int]string, variables.N)
// 	ResponseAddressesIP = make(map[int]string, variables.N)

// 	for i := 0; i < variables.N; i++ {
// 		ServerAddressesIP[i] = "tcp://" + addresses[i] + ":" + strconv.Itoa(7000+(i*100)+variables.ID)
// 		ResponseAddressesIP[i] = "tcp://" + addresses[i] + ":" + strconv.Itoa(10000+(i*100)+variables.ID)
// 	}
// }

// // GetServerAddress - Returns the IP Server address of the server with that id
// func GetServerAddress(id int) string {
// 	return ServerAddressesIP[id]
// }

// // GetResponseAddress - Returns the IP Response address of the server with that id
// func GetResponseAddress(id int) string {
// 	return ResponseAddressesIP[id]
// }
