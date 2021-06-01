package app

import (
	"HashGraph_Client/logger"
	"HashGraph_Client/messenger"
	"HashGraph_Client/types"
	"HashGraph_Client/variables"
	"fmt"
	"math/rand"
	"time"
)

var (
	runes = []rune("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")

	replies  = make(map[int]map[int]bool) // id, from
	accepted = make(map[int]bool)         // if this id is accepted

	// Client metrics regarding the experiment evaluation
	sentTime  = make(map[int]time.Time)
	OpLatency = time.Duration(0)
	Rounds    = 0
	Empty     = 0
	stop      = false
)

///**********************************************
func Client() {

	go ClientTransactionCreation()

	go IncomingMessages()

}

var firstSend time.Time
var NumOfTransactions int
var EmptyTransaction int

func IncomingMessages() {

	totalAck := 0

	for message := range messenger.ResponseChannel {
		if _, in := replies[message.Value][message.From]; in {
			continue // Only one value can be received from each server
		}
		if replies[message.Value] == nil {
			replies[message.Value] = make(map[int]bool)
		}
		replies[message.Value][message.From] = true

		// If more than f+1 with the same value, accept the array.
		if len(replies[message.Value]) >= (variables.F+1) && !accepted[message.Value] {
			accepted[message.Value] = true
			OpLatency += time.Since(sentTime[message.Value])
			logger.OutLogger.Print("RECEIVED ACK for ", message.Value, " [ ",
				time.Since(sentTime[message.Value]).Seconds(), " ] -", "\n")
			// currentTime := time.Now()
			// currentHour := fmt.Sprintf("%v:%v:%v", currentTime.Hour(), currentTime.Minute(), currentTime.Second())
			totalAck += 1
			logger.OutLogger.Printf("Time(%d): %v", totalAck, time.Since(firstSend).Seconds())

			if totalAck == NumOfTransactions {
				logger.OutLogger.Printf("Total for %d: %v", NumOfTransactions, time.Since(firstSend).Seconds())
				stop = true
				fmt.Printf("\tDONE - %d\n", variables.ID)

			}
		}
	}
}

func ClientTransactionCreation() {
	rand.Seed(time.Now().UnixNano())

	//firstWait := 0
	firstWait := variables.ID / 4
	time.Sleep(time.Duration(firstWait) * time.Second)

	//frequency := variables.ID/2 + 1
	//frequency := variables.ID/5 + 1
	frequency := 3

	// for 20 clients
	// NumOfTransactions := 10
	// EmptyTransaction := 5

	//for sleep scenario
	//example set  clients=4 / nodes=4
	NumOfTransactions = 20
	EmptyTransaction = 0

	// NumOfTransactions = 1
	// EmptyTransaction = 0

	// if variables.ID > 0 {
	// 	return
	// }

	character := rune('a' + 0)

	time.Sleep(time.Duration(frequency) * time.Second)

	firstSend = time.Now()
	num := 0
	for i := 1; i <= NumOfTransactions; i += 1 {
		num = i
		transaction := fmt.Sprintf("%d%c", num, character)

		sendTransaction(num, transaction)

		//fmt.Println(variables.ID, ": ", transaction, " ", num, " to: ", to)

		time.Sleep(time.Duration(frequency) * time.Second)

	}

	betweenWait := 3

	time.Sleep(time.Duration(betweenWait) * time.Second)
	//time.Sleep(time.Duration(NumOfTransactions*2+frequency) * time.Second)
	frequency = 2
	for i := NumOfTransactions + 1; i <= NumOfTransactions+EmptyTransaction; i += 1 {
		//for i := NumOfTransactions + 1; !stop; i += 1 {

		num = i
		transaction := fmt.Sprintf("empty%d", num)

		// sentTime[num] = time.Now()

		// message := types.NewClientMsg(transaction, num)

		// to := rand.Intn(variables.N)
		// messenger.SendRequestMsg(message, to)

		sendTransaction(num, transaction)

		//fmt.Println(variables.ID, ": ", transaction, " ", num, " to: ", to)

		time.Sleep(time.Duration(frequency) * time.Second)

	}

	fmt.Printf("\n(%d)Sent - %d", num, variables.ID)

}

func sendTransaction(num int, transaction string) {
	sentTime[num] = time.Now()

	message := types.NewClientMsg(transaction, num)

	owner := make([]bool, variables.N, variables.N) // count the owners we visited

	count := variables.F + 1
	// count := 2
	// to := -2
	for count > 0 {

		to := rand.Intn(variables.N)
		//to += 2
		if !owner[to] {
			owner[to] = true
			messenger.SendRequestMsg(message, to)
			count--
			//	fmt.Println(variables.ID, ": ", transaction, " ", num, " to: ", to)

		}
	}

}
