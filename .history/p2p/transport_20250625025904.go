package p2p

import "net"



type Transport interface{
    Li

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}