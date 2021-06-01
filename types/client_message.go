package types

import (
	"HashGraph_Client/logger"
	"bytes"
	"encoding/gob"
)

// ClientMessage - Client message struct
type ClientMessage struct {
	Cid   int
	Num   int
	Value rune
}

type ClientMsg struct {
	ClientTransaction string
	TransactionNumber int
}

// NewClientMessage - Creates a new Client message
func NewClientMsg(transaction string, num int) ClientMsg {
	return ClientMsg{ClientTransaction: transaction, TransactionNumber: num}
}

// NewClientMessage - Creates a new Client message
func NewClientMessage(id int, num int, value rune) ClientMessage {
	return ClientMessage{Cid: id, Num: num, Value: value}
}

// GobEncode - Client message encoder
func (cm ClientMessage) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(cm.Cid)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	err = encoder.Encode(cm.Num)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	err = encoder.Encode(cm.Value)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	return w.Bytes(), nil
}

// GobDecode - Client message decoder
func (cm *ClientMessage) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&cm.Cid)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	err = decoder.Decode(&cm.Num)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	err = decoder.Decode(&cm.Value)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	return nil
}

// GobEncode - Client message encoder
func (cm *ClientMsg) GobEncodeMsg() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(cm.ClientTransaction)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	err = encoder.Encode(cm.TransactionNumber)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	return w.Bytes(), nil
}

// GobDecode - Client message decoder
func (cm *ClientMsg) GobDecodeMsg(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&cm.ClientTransaction)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	err = decoder.Decode(&cm.TransactionNumber)
	if err != nil {
		logger.ErrLogger.Fatal(err)
	}
	return nil
}
