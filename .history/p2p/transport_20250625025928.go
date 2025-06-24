package p2p

import "net"



type Transport interface{
    neAddr

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}