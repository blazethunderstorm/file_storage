package p2p

import "io"

type Decoder interface{
	Decode(io.Reader,*RPC) error 
}


type GOBDecoder struct{}


func(d *GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	// Implement GOB decoding logic here
	return nil
}

type DefaultDecoder struct{}