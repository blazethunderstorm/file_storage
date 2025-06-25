package p2p

import (
	"net"
	"sync"
)

type TCPPeer struct{
	net.Conn
	outBound bool
	wg   *sync.WaitGroup
}

func NewTCPPeer(conn net.Conn, outBound bool, wg *sync.WaitGroup) *TCPPeer {
	return &TCPPeer{
		Conn: conn,
		outBound: outBound,
		wg: &S,
	}
}