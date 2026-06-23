package msg_test

import (
	"encoding/xml"
	"testing"

	"github.com/tomcyr/nolfix/msg"
)

func TestNewOrderSingleRoundtrip(t *testing.T) {
	stopPx := float32(199.0)
	orig := msg.NewOrderSingle{
		ID:        "1",
		TrdDt:     "20130417",
		Acct:      "00-55-012345",
		Side:      string(msg.SideTypeBuy),
		TxnTm:     "20130417",
		OrdTyp:    string(msg.OrderTypePEGL),
		StopPx:    &stopPx,
		Ccy:       "PLN",
		TmInForce: "0",
		Instrmt:   &msg.Instrument{ID: "PLCOMAR00012", Src: intPtr(4)},
		OrdQty:    &msg.OrderQtyData{Qty: 10.0},
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.NewOrderSingle
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ID != orig.ID || got.Acct != orig.Acct || got.Side != orig.Side {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
	if got.Instrmt == nil || got.Instrmt.ID != orig.Instrmt.ID {
		t.Errorf("Instrmt mismatch: got %+v", got.Instrmt)
	}
}

func TestOrderCancelRequestRoundtrip(t *testing.T) {
	orig := msg.OrderCancelRequest{
		ID:    "1",
		OrdID: "168193517",
		Acct:  "00-55-006638",
		Side:  string(msg.SideTypeBuy),
		TxnTm: "20100628-08:01:23",
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.OrderCancelRequest
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ID != orig.ID || got.OrdID != orig.OrdID {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestExecutionReportRoundtrip(t *testing.T) {
	leavesQty := float32(1.0)
	cumQty := float32(0.0)
	orig := msg.ExecutionReport{
		ID:        "5",
		OrdID:     "178803909",
		Stat:      "0",
		Acct:      "00-55-012345",
		Side:      "1",
		LeavesQty: &leavesQty,
		CumQty:    &cumQty,
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.ExecutionReport
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ID != orig.ID || got.OrdID != orig.OrdID || got.Stat != orig.Stat {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}
