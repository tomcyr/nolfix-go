// /home/tomcyr/code/nolfix/errors.go
package nolfix

import (
	"fmt"
	"github.com/tomcyr/nolfix/msg"
)

type NolConnectionError struct {
	Cause error
}

func (e *NolConnectionError) Error() string { return "nol connection error: " + e.Cause.Error() }
func (e *NolConnectionError) Unwrap() error  { return e.Cause }

type BusinessMessageRejectError struct {
	Message msg.BusinessMessageReject
}

func (e *BusinessMessageRejectError) Error() string {
	return fmt.Sprintf("business message reject: RefMsgTyp=%s BizRejRsn=%v Txt=%s",
		e.Message.RefMsgTyp, e.Message.BizRejRsn, e.Message.Txt)
}

type MarketDataRequestRejectError struct {
	Message msg.MarketDataRequestReject
}

func (e *MarketDataRequestRejectError) Error() string {
	return fmt.Sprintf("market data request reject: ReqID=%s ReqRejResn=%s",
		e.Message.ReqID, e.Message.ReqRejResn)
}
