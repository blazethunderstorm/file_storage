package p2p

import "io"

type Decoder interface{
	Decode(io.Reader,*RPC) error 
}


type GOBDecoder struct{}


func(d *GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	
}

type DefaultDecoder struct{}