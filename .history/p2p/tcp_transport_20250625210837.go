package p2p

import (
	"net"
	"sync"
)

type TCPPeer struct{
	net.Conn
	outBound bool
	wg   sync sync.WaitGroup
}