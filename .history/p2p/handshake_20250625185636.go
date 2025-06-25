package p2p

import "io"

type Decoder interface{
	Decode()
}