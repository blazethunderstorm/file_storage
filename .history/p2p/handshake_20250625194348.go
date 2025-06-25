package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface{
	Decode(io.Reader,*RPC) error 
}


type GOBDecoder struct{}


func(dec *GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct{}


func (dec *DefaultDecoder) Decode(r io.Reader, rpc *RPC) error {
	firstbuff:= make([]byte, 1)
	_, err := r.Read(firstbuff)
	if err != nil {
		return nil
	}
}