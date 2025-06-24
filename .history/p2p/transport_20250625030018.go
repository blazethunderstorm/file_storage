package p2p

import "net"



type Transport interface{
    Addr()
	Dial()
	
	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}