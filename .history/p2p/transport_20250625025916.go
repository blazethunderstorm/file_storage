package p2p

import "net"



type Transport interface{
    Add

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}