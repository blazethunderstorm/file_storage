package p2p

import "net"

type TCPPeer struct{
	conn net.conn
}