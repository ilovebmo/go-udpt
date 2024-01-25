package udpt

import (
	"encoding/binary"
	"hash"
	"net"
	"slices"

	"golang.org/x/crypto/blake2b"
)

func ErrorReply(msg Message, errorMessage string) []byte {
	var reply []byte
	reply = append(reply, ERROR...)
	reply = append(reply, msg.TransactionId()...)
	reply = append(reply, []byte(errorMessage)...)
	return reply
}

func ConnectionReply(msg Message, connId uint64) []byte {
	var reply []byte
	reply = append(reply, msg.Action()...)
	reply = append(reply, msg.TransactionId()...)
	reply = binary.BigEndian.AppendUint64(reply, connId)
	return reply
}

func AnnouncementReply(msg Message, torrent_list map[hash.Hash][]Peer) []byte {
	var reply []byte
	reply = append(reply, msg.Action()...)
	reply = append(reply, msg.TransactionId()...)
	reply = binary.BigEndian.AppendUint32(reply, 60)

	var peers []Peer
	var leechers []*net.UDPAddr
	var seeders []*net.UDPAddr
	h, _ := blake2b.New(32, msg.InfoHash())
	for k, v := range torrent_list {
		if slices.Equal(k.Sum([]byte{}), h.Sum([]byte{})) {
			peers = v
		}
	}
	for _, peer := range peers {
		if peer.Left > 0 {
			leechers = append(leechers, peer.Addr)
		} else {
			seeders = append(seeders, peer.Addr)
		}
	}

	reply = binary.BigEndian.AppendUint32(reply, uint32(len(leechers)))
	reply = binary.BigEndian.AppendUint32(reply, uint32(len(seeders)))

	for _, addr := range leechers {
		reply = append(reply, addr.IP...)
		reply = binary.BigEndian.AppendUint16(reply, uint16(addr.Port))
	}
	for _, addr := range seeders {
		reply = append(reply, addr.IP...)
		reply = binary.BigEndian.AppendUint16(reply, uint16(addr.Port))
	}

	return reply
}
