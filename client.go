// /home/tomcyr/code/nolfix/client.go
package nolfix

import (
	"io"
	"net"
	"strconv"

	"github.com/tomcyr/nolfix/msg"
)

// Sender can connect, send FIXML messages, and disconnect.
type Sender interface {
	Connect() error
	Disconnect() error
	Send(f *msg.Fixml) error
	SendRaw(rawMsg string) error
}

// Receiver can connect, receive FIXML messages, and disconnect.
type Receiver interface {
	Connect() error
	Disconnect() error
	Receive() (*msg.Fixml, error)
}

// NolClient is a TCP client for the Nol3 broker API. It implements both Sender and Receiver.
type NolClient struct {
	host string
	port int
	conn net.Conn
}

// NewNolClient creates a NolClient that will connect to the given host and port.
func NewNolClient(host string, port int) *NolClient {
	return &NolClient{host: host, port: port}
}

// NewNolClientFromConn creates a NolClient using an existing net.Conn (for testing).
func NewNolClientFromConn(conn net.Conn) *NolClient {
	return &NolClient{conn: conn}
}

// Connect opens the TCP connection to the Nol3 server.
func (c *NolClient) Connect() error {
	conn, err := net.Dial("tcp", c.host+":"+strconv.Itoa(c.port))
	if err != nil {
		return &NolConnectionError{Cause: err}
	}
	c.conn = conn
	return nil
}

// Disconnect closes the TCP connection.
func (c *NolClient) Disconnect() error {
	if err := c.conn.Close(); err != nil {
		return &NolConnectionError{Cause: err}
	}
	return nil
}

// Send marshals f to FIXML XML and writes it to the server.
func (c *NolClient) Send(f *msg.Fixml) error {
	xmlStr, err := msg.Serialize(f)
	if err != nil {
		return &NolConnectionError{Cause: err}
	}
	return c.SendRaw(xmlStr)
}

// SendRaw writes a raw XML string to the server using the Nol3 wire format:
// ASCII decimal length of (payload+null), then the payload bytes, then a null byte.
func (c *NolClient) SendRaw(rawMsg string) error {
	msgBytes := []byte(rawMsg)
	payload := make([]byte, len(msgBytes)+1) // +1 for null terminator
	copy(payload, msgBytes)
	// payload[len(msgBytes)] is already 0 (zero-initialized)

	sizeStr := strconv.Itoa(len(payload))
	if _, err := c.conn.Write([]byte(sizeStr)); err != nil {
		return &NolConnectionError{Cause: err}
	}
	if _, err := c.conn.Write(payload); err != nil {
		return &NolConnectionError{Cause: err}
	}
	return nil
}

// Receive reads one FIXML message from the server.
// The Nol3 server uses a 4-byte binary header where bytes 0-2 encode the message
// size as a little-endian uint24, followed by the message bytes ending in a null byte.
func (c *NolClient) Receive() (*msg.Fixml, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(c.conn, header); err != nil {
		return nil, &NolConnectionError{Cause: err}
	}

	size := int(header[0]) + int(header[1])*256 + int(header[2])*65536
	if size == 0 {
		return &msg.Fixml{}, nil
	}

	buf := make([]byte, size)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return nil, &NolConnectionError{Cause: err}
	}

	// Strip trailing null byte
	if len(buf) > 0 && buf[len(buf)-1] == 0 {
		buf = buf[:len(buf)-1]
	}

	rawXML := string(buf)
	f, err := msg.Deserialize(rawXML)
	if err != nil {
		return nil, &NolConnectionError{Cause: err}
	}
	f.RawMessage = rawXML
	return f, nil
}
