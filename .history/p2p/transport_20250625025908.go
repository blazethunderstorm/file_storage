package p2p

import "net"



type Transport interface{
    

	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}