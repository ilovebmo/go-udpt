package udpt

import "net"

type Message struct {
	msg    []byte
	Length int
	Addr   *net.UDPAddr
	Flags  int
}

func (m Message) Contents() []byte {
	return m.msg[0:m.Length]
}

func (m Message) ConnectionId() []byte {
	return m.Contents()[0:8]
}

func (m Message) Magic() []byte {
	return m.ConnectionId()
}

func (m Message) Action() []byte {
	return m.Contents()[8:12]
}

func (m Message) TransactionId() []byte {
	return m.Contents()[12:16]
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
