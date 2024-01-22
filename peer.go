package udpt

import (
	"math/rand"
	"net"
)

type Peer struct {
	Addr     *net.UDPAddr
	ConnId   uint64
	Torrents [][]byte
}

func (p *Peer) Initialize(addr *net.UDPAddr) {
	p.Addr = addr
	p.ConnId = rand.Uint64()
}
