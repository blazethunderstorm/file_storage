package p2p

import "net"



type Transport interface{
    Addr

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}