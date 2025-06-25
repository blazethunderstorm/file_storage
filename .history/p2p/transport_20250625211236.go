package p2p

import "net"



type Transport interface{
    Addr() string 
	Dial(string) error
	ListenAndAccept(string) error
	
	Close() 
}


type Peer interface {
	net.Conn
	Send([]byte) error
	closeStream()
}