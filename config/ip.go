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
// 	"192.36.94.2",
// 	"141.22.213.35",
// 	"139.30.241.191",
// 	"132.227.123.14",
// 	"129.242.19.196",
// 	"141.24.249.131",
// 	"130.192.157.138",
// 	"141.22.213.34",
// 	"192.33.193.18",
// 	"192.33.193.16",
// 	"131.246.19.201",
// 	"155.185.54.249",
// 	"128.232.103.202",
// 	"195.251.248.180",
// 	"194.42.17.164",
// 	"128.232.103.201",
// 	"193.1.201.27",
// 	"193.226.19.30",
// 	"132.65.240.103",
// 	"193.1.201.26",
// 	"129.16.20.70",
// 	"129.16.20.71",
// 	"195.113.161.13",
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
