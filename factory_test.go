// /home/tomcyr/code/nolfix/factory_test.go
package nolfix_test

import (
	"testing"

	nolfix "github.com/tomcyr/nolfix-go"
)

// staticIDGenerator returns a fixed ID for testing.
type staticIDGenerator struct{ id string }

func (s staticIDGenerator) NextID() string { return s.id }

func TestCreateUserRequest_SetsUserReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "test-id-1"})
	req := f.CreateUserRequest()
	if req.UserReqID == "" {
		t.Fatal("expected non-empty UserReqID")
	}
	if req.UserReqID != "test-id-1" {
		t.Errorf("expected UserReqID=%q, got %q", "test-id-1", req.UserReqID)
	}
}

func TestCreateSecurityListRequest_SetsReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "sec-id"})
	req := f.CreateSecurityListRequest()
	if req.ReqID == "" {
		t.Fatal("expected non-empty ReqID")
	}
	if req.ReqID != "sec-id" {
		t.Errorf("expected ReqID=%q, got %q", "sec-id", req.ReqID)
	}
}

func TestCreateTradingSessionStatusRequest_SetsReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "sess-id"})
	req := f.CreateTradingSessionStatusRequest()
	if req.ReqID == "" {
		t.Fatal("expected non-empty ReqID")
	}
	if req.ReqID != "sess-id" {
		t.Errorf("expected ReqID=%q, got %q", "sess-id", req.ReqID)
	}
}

func TestCreateMarketDataRequest_SetsReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "mkt-id"})
	req := f.CreateMarketDataRequest()
	if req.ReqID == "" {
		t.Fatal("expected non-empty ReqID")
	}
	if req.ReqID != "mkt-id" {
		t.Errorf("expected ReqID=%q, got %q", "mkt-id", req.ReqID)
	}
}

func TestCreateNewOrderSingle_SetsID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "order-id"})
	req := f.CreateNewOrderSingle()
	if req.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if req.ID != "order-id" {
		t.Errorf("expected ID=%q, got %q", "order-id", req.ID)
	}
}

func TestCreateOrderCancelRequest_SetsID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "cancel-id"})
	req := f.CreateOrderCancelRequest()
	if req.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if req.ID != "cancel-id" {
		t.Errorf("expected ID=%q, got %q", "cancel-id", req.ID)
	}
}

func TestCreateOrderCancelReplaceRequest_SetsID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "replace-id"})
	req := f.CreateOrderCancelReplaceRequest()
	if req.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if req.ID != "replace-id" {
		t.Errorf("expected ID=%q, got %q", "replace-id", req.ID)
	}
}

func TestCreateOrderStatusRequest_SetsStatReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: "stat-id"})
	req := f.CreateOrderStatusRequest()
	if req.StatReqID == "" {
		t.Fatal("expected non-empty StatReqID")
	}
	if req.StatReqID != "stat-id" {
		t.Errorf("expected StatReqID=%q, got %q", "stat-id", req.StatReqID)
	}
}

func TestRequestFactory_EachCallGetsUniqueID(t *testing.T) {
	counter := 0
	gen := &counterIDGenerator{count: &counter}
	f := nolfix.NewRequestFactory(gen)
	_ = f.CreateUserRequest()
	_ = f.CreateNewOrderSingle()
	if counter != 2 {
		t.Errorf("expected 2 calls to NextID, got %d", counter)
	}
}

type counterIDGenerator struct{ count *int }

func (c *counterIDGenerator) NextID() string {
	*c.count++
	return "id"
}
