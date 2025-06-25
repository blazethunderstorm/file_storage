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

func (p *TCPPeer) NewPeer