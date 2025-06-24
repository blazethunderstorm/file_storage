package p2p

import "net"



type Transport interface{
    netAddr

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}