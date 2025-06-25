package p2p

import "net"

type TCPPeer struct{
	net.Conn
	
}