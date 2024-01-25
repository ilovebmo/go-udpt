package udpt

import (
	"encoding/binary"
	"net"
	"slices"
)

type Message struct {
	msg    []byte
	Length int
	Addr   *net.UDPAddr
	Flags  int
}

func (m Message) Contents() []byte {
	return m.msg[0:m.Length]
}

func (m Message) ConnectionId() uint64 {
	return binary.BigEndian.Uint64(m.Contents()[0:8])
}

func (m Message) Magic() []byte {
	return m.Contents()[0:8]
}

func (m Message) Action() []byte {
	return m.Contents()[8:12]
}

func (m Message) TransactionId() []byte {
	return m.Contents()[12:16]
}

func (m Message) InfoHash() []byte {
	return m.Contents()[16:36]
}

func (m Message) PeerId() []byte {
	return m.Contents()[36:56]
}

func (m Message) Downloaded() []byte {
	return m.Contents()[56:64]
}

func (m Message) Left() []byte {
	return m.Contents()[64:72]
}

func (m Message) Uploaded() []byte {
	return m.Contents()[72:80]
}

func (m Message) Event() []byte {
	return m.Contents()[80:84]
}

func (m Message) IPAddr() ([]byte, error) {
	if slices.Equal(m.Contents()[84:88], NONE) {
		return m.Contents()[84:88], NoIPAddrError{}
	}
	return m.Contents()[84:88], nil
}

func (m Message) Key() []byte {
	return m.Contents()[88:92]
}

func (m Message) NumWant() []byte {
	return m.Contents()[92:96]
}

func (m Message) Port() []byte {
	return m.Contents()[96:98]
}

type NoIPAddrError struct{}

func (e NoIPAddrError) Error() string {
	return "IPAddr set to 0, user requested to use the sender IP"
}

func GetMessage(connection *net.UDPConn, messageChannel chan<- Message, errorChannel chan<- error) {
	for {
		msg := make([]byte, 150) // TODO: Extensions; 150 should hold everything for now
		lengthMsg, _, flags, addr, err := connection.ReadMsgUDP(msg, []byte{})

		if lengthMsg >= 16 {
			messageChannel <- Message{msg, lengthMsg, addr, flags}
		}

		if err != nil {
			errorChannel <- err
		}
	}
}
