package msg_test

import (
	"encoding/xml"
	"testing"
	"github.com/tomcyr/nolfix/msg"
)

func TestTradingSessionStatusRequestRoundtrip(t *testing.T) {
	orig := msg.TradingSessionStatusRequest{
		ReqID:     "1",
		SubReqTyp: string(msg.SessSubReqTypOnline),
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.TradingSessionStatusRequest
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ReqID != orig.ReqID || got.SubReqTyp != orig.SubReqTyp {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestTradingSessionStatusRoundtrip(t *testing.T) {
	orig := msg.TradingSessionStatus{
		ReqID:  "1",
		SesSub: "COCO",
		Instrmts: []msg.Instrument{
			{Sym: "TPSA", ID: "PLTLKPL00017", Src: intPtr(4)},
		},
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.TradingSessionStatus
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ReqID != orig.ReqID || len(got.Instrmts) != 1 {
		t.Errorf("roundtrip mismatch: got %+v", got)
	}
}
