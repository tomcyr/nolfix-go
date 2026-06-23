// /home/tomcyr/code/nolfix/request_reply_test.go
package nolfix_test

import (
	"errors"
	"net"
	"testing"

	nolfix "github.com/tomcyr/nolfix"
	"github.com/tomcyr/nolfix/msg"
)

// startEchoServer runs a goroutine that reads one message via simulateServerReceive
// (defined in client_test.go) and replies with xmlResponse.
func startEchoServer(t *testing.T, serverConn net.Conn, xmlResponse string) chan error {
	t.Helper()
	done := make(chan error, 1)
	go func() {
		defer serverConn.Close()
		_, err := simulateServerReceive(serverConn)
		if err != nil {
			done <- err
			return
		}
		done <- simulateServerSend(serverConn, xmlResponse)
	}()
	return done
}

func TestNolRequestReplyClient_Send_ReturnsUserResponse(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer clientConn.Close()

	responseXML := `<FIXML v="5.0" r="20080317" s="20080314"><UserRsp UserReqID="1" Username="BOS" UserStat="1"/></FIXML>`
	done := startEchoServer(t, serverConn, responseXML)

	// Build a request
	f := msg.NewFixml()
	n := 1
	f.UserRequests = []msg.UserRequest{{UserReqID: "1", UserReqTyp: &n, Username: "BOS", Password: "BOS"}}

	// Use NolRequestReplyClient but inject the pre-wired conn via a loopback trick:
	// We need to connect to a real address, so use NewNolClientFromConn approach.
	// Since NolRequestReplyClient uses NewNolClient internally and calls Connect(),
	// we test it through a thin adapter using net.Pipe trick via a listener.
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer l.Close()
	addr := l.Addr().(*net.TCPAddr)

	// Accept one connection in the background and proxy to our pipe
	accepted := make(chan error, 1)
	go func() {
		conn, err := l.Accept()
		if err != nil {
			accepted <- err
			return
		}
		defer conn.Close()
		// Bridge: relay between conn (from NolRequestReplyClient) and clientConn (our pipe)
		// Actually simpler: just re-implement the echo server on the accepted conn directly
		_, err = simulateServerReceive(conn)
		if err != nil {
			accepted <- err
			return
		}
		accepted <- simulateServerSend(conn, responseXML)
	}()
	// Close the unused pipe side so we don't block
	serverConn.Close()
	clientConn.Close()
	_ = done // original pipe done channel unused now

	client := nolfix.NewNolRequestReplyClient("127.0.0.1", addr.Port)
	resp, err := client.Send(f)
	if err != nil {
		t.Fatalf("Send: %v", err)
	}
	if err := <-accepted; err != nil {
		t.Fatalf("server error: %v", err)
	}

	inner, err := resp.Unpack()
	if err != nil {
		t.Fatalf("Unpack: %v", err)
	}
	if inner.MsgName() != "UserRsp" {
		t.Errorf("expected UserRsp, got %q", inner.MsgName())
	}
}

func TestNolRequestReplyClient_Send_ReturnsBizMsgRejectError(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer l.Close()
	addr := l.Addr().(*net.TCPAddr)

	rejectXML := `<FIXML v="5.0" r="20080317" s="20080314"><BizMsgRej RefMsgTyp="UserReq" BizRejRsn="1" Txt="Rejected"/></FIXML>`
	go func() {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		_, _ = simulateServerReceive(conn)
		_ = simulateServerSend(conn, rejectXML)
	}()

	f := msg.NewFixml()
	n := 1
	f.UserRequests = []msg.UserRequest{{UserReqID: "1", UserReqTyp: &n}}

	client := nolfix.NewNolRequestReplyClient("127.0.0.1", addr.Port)
	_, err = client.Send(f)
	if err == nil {
		t.Fatal("expected BusinessMessageRejectError, got nil")
	}
	var bmErr *nolfix.BusinessMessageRejectError
	if !errors.As(err, &bmErr) {
		t.Errorf("expected BusinessMessageRejectError, got %T: %v", err, err)
	}
}

func TestNolRequestReplyClient_Send_ReturnsMarketDataRejectError(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer l.Close()
	addr := l.Addr().(*net.TCPAddr)

	rejectXML := `<FIXML v="5.0" r="20080317" s="20080314"><MktDataReqRej ReqID="mkt-1" ReqRejResn="1"/></FIXML>`
	go func() {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		_, _ = simulateServerReceive(conn)
		_ = simulateServerSend(conn, rejectXML)
	}()

	f := msg.NewFixml()
	f.MktDataReqs = []msg.MarketDataRequest{{ReqID: "mkt-1"}}

	client := nolfix.NewNolRequestReplyClient("127.0.0.1", addr.Port)
	_, err = client.Send(f)
	if err == nil {
		t.Fatal("expected MarketDataRequestRejectError, got nil")
	}
	var mdErr *nolfix.MarketDataRequestRejectError
	if !errors.As(err, &mdErr) {
		t.Errorf("expected MarketDataRequestRejectError, got %T: %v", err, err)
	}
}

func TestNolRequestReplyClient_Send_ConnectionError(t *testing.T) {
	// Port with nothing listening
	client := nolfix.NewNolRequestReplyClient("127.0.0.1", 1)
	f := msg.NewFixml()
	_, err := client.Send(f)
	if err == nil {
		t.Fatal("expected error when connecting to port 1, got nil")
	}
	var connErr *nolfix.NolConnectionError
	if !errors.As(err, &connErr) {
		t.Errorf("expected NolConnectionError, got %T: %v", err, err)
	}
}
