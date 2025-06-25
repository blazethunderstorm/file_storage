package p2p

import "io"

type Decoder interface{
	Decode(r io.Reader,*RPC) error 
}