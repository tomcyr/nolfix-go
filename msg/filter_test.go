package msg_test

import (
	"encoding/xml"
	"testing"

	"github.com/tomcyr/nolfix-go/msg"
)

func TestMarketDataRequestRoundtrip(t *testing.T) {
	depth := 1
	orig := msg.MarketDataRequest{
		ReqID:     "1",
		SubReqTyp: string(msg.SubReqTypAddToFilter),
		MktDepth:  &depth,
		Reqs: []msg.MdReqGrp{
			{Typ: string(msg.EntryTypeOffer)},
			{Typ: string(msg.EntryTypeLastTrade)},
		},
		InstReq: &msg.InstReq{
			Instrmts: []msg.Instrument{
				{Sym: "COMARCH", ID: "PLCOMAR00012", Src: intPtr(4)},
			},
		},
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.MarketDataRequest
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ReqID != orig.ReqID || len(got.Reqs) != 2 {
		t.Errorf("roundtrip mismatch: got %+v", got)
	}
	if got.InstReq == nil || len(got.InstReq.Instrmts) != 1 {
		t.Errorf("InstReq mismatch: got %+v", got.InstReq)
	}
}

func TestMarketDataIncrementalRefreshRoundtrip(t *testing.T) {
	sz := float32(10.0)
	orig := msg.MarketDataIncrementalRefresh{
		ReqID: "20",
		Incs: []msg.MdIncGrp{
			{
				Typ:      string(msg.EntryTypeOffer),
				Sz:       &sz,
				Instrmts: []msg.Instrument{{Sym: "COMARCH", ID: "PLCOMAR00012", Src: intPtr(4)}},
			},
		},
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.MarketDataIncrementalRefresh
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ReqID != orig.ReqID || len(got.Incs) != 1 {
		t.Errorf("roundtrip mismatch: got %+v", got)
	}
}
