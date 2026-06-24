// /home/tomcyr/code/nolfix/factory_test.go
package nolfix_test

import (
	"testing"

	nolfix "github.com/tomcyr/nolfix-go"
)

// staticIDGenerator returns a fixed ID for testing.
type staticIDGenerator struct{ id int }

func (s staticIDGenerator) NextID() int { return s.id }

func TestCreateUserRequest_SetsUserReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 1})
	req := f.CreateUserRequest()
	if req.UserReqID == "" {
		t.Fatal("expected non-empty UserReqID")
	}
	if req.UserReqID != "1" {
		t.Errorf("expected UserReqID=%q, got %q", "1", req.UserReqID)
	}
}

func TestCreateSecurityListRequest_SetsReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 2})
	req := f.CreateSecurityListRequest()
	if req.ReqID == "" {
		t.Fatal("expected non-empty ReqID")
	}
	if req.ReqID != "2" {
		t.Errorf("expected ReqID=%q, got %q", "2", req.ReqID)
	}
}

func TestCreateTradingSessionStatusRequest_SetsReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 3})
	req := f.CreateTradingSessionStatusRequest()
	if req.ReqID == "" {
		t.Fatal("expected non-empty ReqID")
	}
	if req.ReqID != "3" {
		t.Errorf("expected ReqID=%q, got %q", "3", req.ReqID)
	}
}

func TestCreateMarketDataRequest_SetsReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 4})
	req := f.CreateMarketDataRequest()
	if req.ReqID == "" {
		t.Fatal("expected non-empty ReqID")
	}
	if req.ReqID != "4" {
		t.Errorf("expected ReqID=%q, got %q", "4", req.ReqID)
	}
}

func TestCreateNewOrderSingle_SetsID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 5})
	req := f.CreateNewOrderSingle()
	if req.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if req.ID != "5" {
		t.Errorf("expected ID=%q, got %q", "5", req.ID)
	}
}

func TestCreateOrderCancelRequest_SetsID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 6})
	req := f.CreateOrderCancelRequest()
	if req.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if req.ID != "6" {
		t.Errorf("expected ID=%q, got %q", "6", req.ID)
	}
}

func TestCreateOrderCancelReplaceRequest_SetsID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 7})
	req := f.CreateOrderCancelReplaceRequest()
	if req.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	if req.ID != "7" {
		t.Errorf("expected ID=%q, got %q", "7", req.ID)
	}
}

func TestCreateOrderStatusRequest_SetsStatReqID(t *testing.T) {
	f := nolfix.NewRequestFactory(staticIDGenerator{id: 8})
	req := f.CreateOrderStatusRequest()
	if req.StatReqID == "" {
		t.Fatal("expected non-empty StatReqID")
	}
	if req.StatReqID != "8" {
		t.Errorf("expected StatReqID=%q, got %q", "8", req.StatReqID)
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

func (c *counterIDGenerator) NextID() int {
	*c.count++
	return *c.count
}
