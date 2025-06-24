package p2p

import "net"



type Transport interface{
    net.AddrAddr

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}