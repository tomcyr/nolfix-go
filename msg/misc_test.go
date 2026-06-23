package msg_test

import (
	"encoding/xml"
	"testing"
	"github.com/tomcyr/nolfix/msg"
)

func TestHeartbeatRoundtrip(t *testing.T) {
	orig := msg.Heartbeat{}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.Heartbeat
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	_ = got
}

func TestNewsRoundtrip(t *testing.T) {
	orig := msg.News{OrigTm: "20080910-10:12:23", Headline: "Test", Message: "Hello world"}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.News
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.Headline != orig.Headline || got.Message != orig.Message {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestApplMsgRptRoundtrip(t *testing.T) {
	orig := msg.ApplMsgRpt{ApplRepID: "1", Txt: "22"}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.ApplMsgRpt
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ApplRepID != orig.ApplRepID || got.Txt != orig.Txt {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestStatementRoundtrip(t *testing.T) {
	orig := msg.Statement{
		Acct:     "00-55-006638",
		AcctType: "M",
		Ike:      "N",
		Funds: []msg.Fund{{Name: "Cash", Value: "100.00"}},
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.Statement
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.Acct != orig.Acct || len(got.Funds) != 1 || got.Funds[0].Name != "Cash" {
		t.Errorf("roundtrip mismatch: got %+v", got)
	}
}
