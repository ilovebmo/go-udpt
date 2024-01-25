package udpt

import (
	"encoding/binary"
	"net"

	"encoding/hex"
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

func AnnouncementReply(msg Message, torrent_list map[string][]Peer) []byte {
	var reply []byte
	reply = append(reply, msg.Action()...)
	reply = append(reply, msg.TransactionId()...)
	reply = binary.BigEndian.AppendUint32(reply, 5)

	var leechers []*net.UDPAddr
	var seeders []*net.UDPAddr

	for _, peer := range torrent_list[hex.Dump(msg.InfoHash())] {
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

func ScrapingReply(msg Message, torrent_list map[string][]Peer) []byte {
	var reply []byte
	reply = append(reply, msg.Action()...)
	reply = append(reply, msg.TransactionId()...)
	for _, torrent := range msg.AllInfoHash() {
		var seeds uint32 = 0
		var downd uint32 = 0
		var incom uint32 = 0
		for _, peer := range torrent_list[hex.Dump(torrent.InfoHash)] {
			if peer.Left > 0 {
				incom++
				continue
			}
			downd++
			seeds++
		}
		reply = binary.BigEndian.AppendUint32(reply, seeds)
		reply = binary.BigEndian.AppendUint32(reply, downd)
		reply = binary.BigEndian.AppendUint32(reply, incom)
	}

	return reply
}
