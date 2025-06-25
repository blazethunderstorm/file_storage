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
		wg: &sync.WaitGroup{},
	}
}

func (p *TCPPeer) closeStream() {
	p.wg.Done()
}

func (p *TCPPeer) Send(b []byte) error {
	p.wg.Add(1)
	defer p.wg.Done()
	_, err := p.Conn.Write(b)
	if err != nil {
		return err
	}
	return nil

}