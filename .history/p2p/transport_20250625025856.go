package p2p

import "net"



type Transport interface{


	Clo
}


type Peer interface {
	net.Conn

	closeStream()
}