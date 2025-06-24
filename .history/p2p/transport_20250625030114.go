package p2p

import "net"



type Transport interface{
    Addr() string 
	Dial(string) error
	ListenAndAcce(string) error
	
	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}