package p2p

import "net"



type Transport interface{
    Addr() string 
	Dial(string) error
	Listen
	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}