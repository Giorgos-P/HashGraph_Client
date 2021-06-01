#!/bin/bash

N=10
CLIENTS=2
REM=0

go install BFTWithoutSignatures_Client

for (( ID=0; ID<$CLIENTS; ID++ ))
do
	BFTWithoutSignatures_Client $ID $N $REM &
done
