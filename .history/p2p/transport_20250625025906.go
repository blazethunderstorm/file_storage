package p2p

import "net"



type Transport interface{
    Listen

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}