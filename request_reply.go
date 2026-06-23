// /home/tomcyr/code/nolfix/request_reply.go
package nolfix

import "github.com/tomcyr/nolfix/msg"

// NolRequestReplyClient sends a message and returns the response in a single call.
type NolRequestReplyClient struct {
	host string
	port int
}

// NewNolRequestReplyClient creates a sync request-reply client for the given address.
func NewNolRequestReplyClient(host string, port int) *NolRequestReplyClient {
	return &NolRequestReplyClient{host: host, port: port}
}

// Send connects, sends f, receives the response, disconnects, and returns the response.
// Returns NolConnectionError on connection/IO failure.
// Returns BusinessMessageRejectError or MarketDataRequestRejectError on rejection responses.
func (c *NolRequestReplyClient) Send(f *msg.Fixml) (*msg.Fixml, error) {
	client := NewNolClient(c.host, c.port)
	if err := client.Connect(); err != nil {
		return nil, err
	}
	defer client.Disconnect() //nolint:errcheck

	if err := client.Send(f); err != nil {
		return nil, err
	}
	resp, err := client.Receive()
	if err != nil {
		return nil, err
	}
	if err := validateResponse(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func validateResponse(f *msg.Fixml) error {
	if len(f.BizMsgRejs) > 0 {
		return &BusinessMessageRejectError{Message: f.BizMsgRejs[0]}
	}
	if len(f.MktDataReqRejs) > 0 {
		return &MarketDataRequestRejectError{Message: f.MktDataReqRejs[0]}
	}
	return nil
}
