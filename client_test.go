// /home/tomcyr/code/nolfix/client_test.go
package nolfix_test

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"testing"

	nolfix "github.com/tomcyr/nolfix-go"
	"github.com/tomcyr/nolfix-go/msg"
)

// simulateServerReceive reads one message from conn as the NOL3 server would:
// read ASCII size digits (until non-digit byte from message start '<'), then read payload.
func simulateServerReceive(conn net.Conn) (string, error) {
	// read digits until we see '<'
	var sizeStr []byte
	buf := make([]byte, 1)
	for {
		if _, err := conn.Read(buf); err != nil {
			return "", err
		}
		if buf[0] == '<' {
			break
		}
		sizeStr = append(sizeStr, buf[0])
	}
	size, _ := strconv.Atoi(string(sizeStr))
	// we already consumed '<', so read size-1 more bytes (size includes null byte; '<' is 1 byte already read)
	rest := make([]byte, size-1)
	if _, err := io.ReadFull(conn, rest); err != nil {
		return "", err
	}
	// full xml = '<' + rest (strip trailing null)
	full := append([]byte{'<'}, rest...)
	if full[len(full)-1] == 0 {
		full = full[:len(full)-1]
	}
	return string(full), nil
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
