// /home/tomcyr/code/nolfix/factory.go
package nolfix

import "github.com/tomcyr/nolfix/msg"

// RequestFactory creates typed FIXML request objects with auto-generated IDs.
type RequestFactory struct {
	idGen IdGenerator
}

// NewRequestFactory creates a RequestFactory using the given ID generator.
func NewRequestFactory(idGen IdGenerator) *RequestFactory {
	return &RequestFactory{idGen: idGen}
}

func (f *RequestFactory) CreateUserRequest() *msg.UserRequest {
	return &msg.UserRequest{UserReqID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateSecurityListRequest() *msg.SecurityListRequest {
	return &msg.SecurityListRequest{ReqID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateTradingSessionStatusRequest() *msg.TradingSessionStatusRequest {
	return &msg.TradingSessionStatusRequest{ReqID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateMarketDataRequest() *msg.MarketDataRequest {
	return &msg.MarketDataRequest{ReqID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateNewOrderSingle() *msg.NewOrderSingle {
	return &msg.NewOrderSingle{ID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateOrderCancelRequest() *msg.OrderCancelRequest {
	return &msg.OrderCancelRequest{ID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateOrderCancelReplaceRequest() *msg.OrderCancelReplaceRequest {
	return &msg.OrderCancelReplaceRequest{ID: f.idGen.NextID()}
}

func (f *RequestFactory) CreateOrderStatusRequest() *msg.OrderStatusRequest {
	return &msg.OrderStatusRequest{StatReqID: f.idGen.NextID()}
}
