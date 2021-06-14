
# Hashgraph Client

## Contents
  - [About](#about)
  - [Transaction](#transaction)
  - [Installation](#installation)
  - [Execution](#execution)
  - [References](#references)

## About
Implementation of Hashgraph in Go programming language with ZeroMQ framework for the algorithm based on: <br>

"The Hashgraph Protocol: Efficient Asynchronous BFT for High-Throughput Distributed Ledgers"[1] <br>
by Leemon Baird and Atul Luykx.

The implementation of the servers can be found [here](https://github.com/Giorgos-P/HashGraphBFT).

## Transaction
Each Client send a number of transaction. A client transaction contains
*  string: Transaction
*  int   : increased integer number
  
## Installation
### Golang
If you have not already installed **Golang** follow the instructions [here](https://golang.org/doc/install).

### Clone Repository
```bash
cd ~/go/src/
git clone https://github.com/Giorgos-P/HashGraph_Client.git
```

## Execution

### Manually
Open different terminal for each client:
```bash
Hashgraph_Client <ID> <N> <Clients> <Remote>
```

ID : is the client id [0 to (N-1)]<br>
N : is the total number of servers<br>
Clients : is the total number of clients<br>
Remote : is [0 or 1] 0=local execution, 1 is remote execution (we have to set the correct ip address of the remote servers in ip.go)

### Script
Adjust the script (Hashgraph_Client/scripts/runClient.sh) and run:
```bash
bash ~/go/src/Hashgraph_Client/scripts/runClient.sh
```

### Kill Processes
When you are done and want to kill the processes run:
```bash
bash ~/go/src/Hashgraph_Client/scripts/killHash.sh
```

## References
- [[1]The Hashgraph Protocol: Efficient Asynchronous BFT for High-Throughput Distributed Ledgers](https://hedera.com/hh-ieee_coins_paper-200516.pdf).


[Go To TOP](#hashgraph-client)