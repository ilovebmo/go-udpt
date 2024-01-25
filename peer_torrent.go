package udpt

import (
	"math/rand"
	"net"
)

type Peer struct {
	Addr   *net.UDPAddr
	ConnId uint64
	Left   uint64
}

func (p *Peer) Initialize(addr *net.UDPAddr) {
	p.Addr = addr
	p.ConnId = rand.Uint64()
}

type Torrent struct {
	InfoHash []byte
}
