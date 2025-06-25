package p2p

import "io"

type Decoder interface{
	Decode(io.Reader,*RPC) error 
}


type GOBDecoder struct{}



type DeafaultDecoder struct{}