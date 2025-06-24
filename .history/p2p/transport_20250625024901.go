package p2p

import "net"



type Transport interface{
	
}


type Peer interface {
	net.Conn
	clos
}