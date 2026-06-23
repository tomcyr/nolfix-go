package msg_test

import (
	"encoding/xml"
	"testing"

	"github.com/tomcyr/nolfix-go/msg"
)

func TestInstrumentRoundtrip(t *testing.T) {
	orig := msg.Instrument{Sym: "COMARCH", ID: "PLCOMAR00012", Src: intPtr(4)}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.Instrument
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.Sym != orig.Sym || got.ID != orig.ID || *got.Src != *orig.Src {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestBusinessMessageRejectRoundtrip(t *testing.T) {
	n := 5
	orig := msg.BusinessMessageReject{RefMsgTyp: "BE", BizRejRsn: &n, Txt: "xml error"}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.BusinessMessageReject
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.RefMsgTyp != orig.RefMsgTyp || *got.BizRejRsn != *orig.BizRejRsn {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func intPtr(v int) *int { return &v }
