// /home/tomcyr/code/nolfix/client_test.go
package nolfix_test

import (
	"fmt"
	"io"
	"net"
	"testing"

	nolfix "github.com/tomcyr/nolfix-go"
	"github.com/tomcyr/nolfix-go/msg"
)

// simulateServerReceive reads one message from conn as the NOL3 server would:
// 4-byte binary header (bytes 0-2 = little-endian uint24 payload length), then payload.
func simulateServerReceive(conn net.Conn) (string, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(conn, header); err != nil {
		return "", err
	}
	size := int(header[0]) + int(header[1])*256 + int(header[2])*65536
	buf := make([]byte, size)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return "", err
	}
	if len(buf) > 0 && buf[len(buf)-1] == 0 {
		buf = buf[:len(buf)-1]
	}
	return string(buf), nil
}

// simulateServerSend sends one message from server using binary header (NOL3 format).
func simulateServerSend(conn net.Conn, xmlStr string) error {
	payload := append([]byte(xmlStr), 0)
	size := len(payload)
	header := []byte{
		byte(size & 0xff),
		byte((size >> 8) & 0xff),
		byte((size >> 16) & 0xff),
		0,
	}
	if _, err := conn.Write(header); err != nil {
		return err
	}
	_, err := conn.Write(payload)
	return err
}

func TestNolClientSendAndReceive(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer serverConn.Close()
	defer clientConn.Close()

	// Build a simple FIXML message to send
	f := msg.NewFixml()
	n := 1
	f.UserRequests = []msg.UserRequest{{UserReqID: "1", UserReqTyp: &n, Username: "BOS", Password: "BOS"}}

	xmlStr, _ := msg.Serialize(f)
	expectedResponse := `<FIXML v="5.0" r="20080317" s="20080314"><UserRsp UserReqID="1" Username="BOS" UserStat="1"/></FIXML>`

	done := make(chan error, 1)
	go func() {
		// server: receive from client, send response
		got, err := simulateServerReceive(serverConn)
		if err != nil {
			done <- err
			return
		}
		if got != xmlStr {
			done <- fmt.Errorf("server got %q, want %q", got, xmlStr)
			return
		}
		done <- simulateServerSend(serverConn, expectedResponse)
	}()

	client := nolfix.NewNolClientFromConn(clientConn)
	if err := client.Send(f); err != nil {
		t.Fatalf("Send: %v", err)
	}
	resp, err := client.Receive()
	if err != nil {
		t.Fatalf("Receive: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatalf("server error: %v", err)
	}
	inner, err := resp.Unpack()
	if err != nil {
		t.Fatal(err)
	}
	if inner.MsgName() != "UserRsp" {
		t.Errorf("expected UserRsp, got %q", inner.MsgName())
	}
}
