package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface{
	Decode(io.Reader,*RPC) error 
}


type GOBDecoder struct{}


func(d *GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct{}


func (d *DefaultDecoder) Decode(r io.Reader, rpc *RPC) error {
	
}