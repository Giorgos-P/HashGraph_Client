package main

import (
	"HashGraph_Client/app"
	"HashGraph_Client/config"
	"HashGraph_Client/logger"
	"HashGraph_Client/messenger"
	"HashGraph_Client/variables"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// Initializer - Method that initializes all required processes
func initializer(id int, n int, clients int, rem int) {
	variables.Initialize(id, n, clients, rem)

	if variables.Remote {
		config.InitializeIP()
	} else {
		config.InitializeLocal()
	}

	logger.OutLogger.Print(
		"ID:", variables.ID, " | N:", variables.N, " | F:", variables.F,
		" | Remote:", variables.Remote, "\n\n",
	)

	messenger.InitializeMessenger()
	messenger.Subscribe()
	messenger.TransmitRequests()

	app.Client()
}

func cleanup() {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for range terminate {
			for i := 0; i < variables.N; i++ {
				messenger.ServerSockets[i].Close()
				messenger.ResponseSockets[i].Close()
			}
			os.Exit(0)
		}
	}()
}

func main() {
	args := os.Args[1:]
	if len(args) == 4 {
		id, _ := strconv.Atoi(args[0])
		n, _ := strconv.Atoi(args[1])
		clients, _ := strconv.Atoi(args[2])
		remote, _ := strconv.Atoi(args[3])

		initializer(id, n, clients, remote)
		cleanup()

		done := make(chan interface{}) // To keep the client running
		<-done

	} else {
		log.Fatal("Arguments should be '<id> <n> <remote>'")
	}
}
