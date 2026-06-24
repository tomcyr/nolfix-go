// /home/tomcyr/code/nolfix/factory.go
package nolfix

import (
	"strconv"

	"github.com/tomcyr/nolfix-go/msg"
)

// RequestFactory creates typed FIXML request objects with auto-generated IDs.
type RequestFactory struct {
	idGen IdGenerator
}

// NewRequestFactory creates a RequestFactory using the given ID generator.
func NewRequestFactory(idGen IdGenerator) *RequestFactory {
	return &RequestFactory{idGen: idGen}
}

func (f *RequestFactory) CreateUserRequest() *msg.UserRequest {
	return &msg.UserRequest{UserReqID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateSecurityListRequest() *msg.SecurityListRequest {
	return &msg.SecurityListRequest{ReqID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateTradingSessionStatusRequest() *msg.TradingSessionStatusRequest {
	return &msg.TradingSessionStatusRequest{ReqID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateMarketDataRequest() *msg.MarketDataRequest {
	return &msg.MarketDataRequest{ReqID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateNewOrderSingle() *msg.NewOrderSingle {
	return &msg.NewOrderSingle{ID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateOrderCancelRequest() *msg.OrderCancelRequest {
	return &msg.OrderCancelRequest{ID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateOrderCancelReplaceRequest() *msg.OrderCancelReplaceRequest {
	return &msg.OrderCancelReplaceRequest{ID: strconv.Itoa(f.idGen.NextID())}
}

func (f *RequestFactory) CreateOrderStatusRequest() *msg.OrderStatusRequest {
	return &msg.OrderStatusRequest{StatReqID: strconv.Itoa(f.idGen.NextID())}
}
