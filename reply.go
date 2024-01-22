package udpt

import "encoding/binary"

func ConnectionReply(msg Message) []byte {
	var reply []byte
	reply = append(reply, msg.Action()...)
	reply = append(reply, msg.TransactionId()...)
	reply = binary.BigEndian.AppendUint64(reply, 6)
	return reply
}
