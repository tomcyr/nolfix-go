package msg_test

import (
	"encoding/xml"
	"testing"

	"github.com/tomcyr/nolfix/msg"
)

func TestSecurityListRequestRoundtrip(t *testing.T) {
	typ := 5
	orig := msg.SecurityListRequest{
		ReqID:      "0",
		ListReqTyp: &typ,
		MktID:      msg.MarketIDNM,
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.SecurityListRequest
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ReqID != orig.ReqID || got.MktID != msg.MarketIDNM {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestSecurityListRoundtrip(t *testing.T) {
	rptID := 0
	orig := msg.SecurityList{
		RptID: &rptID,
		ReqID: "0",
		MktID: "WE",
		SecL: &msg.SecL{
			Instrmts: []msg.Instrument{
				{Sym: "COMARCH", Src: intPtr(4), ID: "PLCOMAR00012", SecGrp: "C7", CFI: "ESXXXX"},
			},
		},
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.SecurityList
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.SecL == nil || len(got.SecL.Instrmts) != 1 {
		t.Errorf("SecL mismatch: got %+v", got.SecL)
	}
	if got.SecL.Instrmts[0].Sym != "COMARCH" {
		t.Errorf("Sym mismatch: got %q", got.SecL.Instrmts[0].Sym)
	}
}
