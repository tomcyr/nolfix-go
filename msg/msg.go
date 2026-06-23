// /home/tomcyr/code/nolfix/msg/msg.go
package msg

// Placeholder types for errors.go — these are defined in Task 9
type BusinessMessageReject struct {
	RefMsgTyp string
	BizRejRsn interface{}
	Txt       string
}

type MarketDataRequestReject struct {
	ReqID      string
	ReqRejResn string
}
