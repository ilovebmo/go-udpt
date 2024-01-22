package udpt

import (
	"encoding/binary"
)

func ConnectionReply(msg Message, connId uint64) []byte {
	var reply []byte
	reply = append(reply, msg.Action()...)
	reply = append(reply, msg.TransactionId()...)
	reply = binary.BigEndian.AppendUint64(reply, connId)
	return reply
}
