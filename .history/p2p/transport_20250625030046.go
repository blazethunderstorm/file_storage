package p2p

import "net"



type Transport interface{
    Addr() string 
	Dial(string) error
	ListenAndServe(string) error
	Close() 
}


type Peer interface {
	net.Conn

	closeStream()
}