package servers

import (
	"errors"
	"log"
	"net"
	"time"
)

var (
	// Indicates that a channel has been exhausted.
	ChannelExhausted = errors.New("channel_exhausted")
	// Indicates that the client received and unexpected reply.
	UnexpectedReply = errors.New("Client received an unexpected reply")
	// Indicates that the server replied, but less than what is expected.
	NotEnoughBytes = errors.New("Client received too few bytes from the server")
)

// Connect to the given address (in <ip>:<port> form) and return the UDP connection or an error.
func connect(addr, laddr string) (connection *net.UDPConn, error error) {

	l, err := net.ResolveUDPAddr("udp", laddr)
	if err != nil {
		log.Fatal("Nope1", err)
		error = err
		return
	}

	udpAddr, error := net.ResolveUDPAddr("udp", addr)
	if error != nil {
		log.Fatal("Nope2", error)
		return
	}

	connection, error = net.DialUDP("udp", l, udpAddr)
	if error != nil {
		log.Fatal("Nope3", error)
	}
	return
}

// Start listening on the given UDP address and return the UDP connection or an error.
func listen(addr *net.UDPAddr) (connection *net.UDPConn, error error) {
	connection, error = net.ListenUDP("udp", addr)
	log.Println(addr, error)
	return
}

// Sets the read deadline on the given connection using the given timeout (which should be a duration). The timeout is added to time.Now().
func setReadDeadline(connection *net.UDPConn, timeout string) error {
	duration, error := time.ParseDuration(timeout)
	if error != nil {
		return error
	}

	connection.SetReadDeadline(time.Now().Add(duration))
	return nil
}
