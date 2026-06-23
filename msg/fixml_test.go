package msg_test

import (
	"encoding/xml"
	"errors"
	"os"
	"testing"

	"github.com/tomcyr/nolfix/msg"
)

func TestSerializeDeserializeRoundtrip(t *testing.T) {
	f := msg.NewFixml()
	f.UserRequests = []msg.UserRequest{{UserReqID: "1", Username: "BOS", Password: "BOS"}}

	xmlStr, err := msg.Serialize(f)
	if err != nil {
		t.Fatal(err)
	}

	got, err := msg.Deserialize(xmlStr)
	if err != nil {
		t.Fatalf("Deserialize(%q): %v", xmlStr, err)
	}
	if len(got.UserRequests) != 1 || got.UserRequests[0].Username != "BOS" {
		t.Errorf("roundtrip mismatch: got %+v", got)
	}
}

func TestDeserializeUserReqFixture(t *testing.T) {
	data, _ := os.ReadFile("testdata/UserReq.xml")
	f, err := msg.Deserialize(string(data))
	if err != nil {
		t.Fatal(err)
	}
	inner, err := f.Unpack()
	if err != nil {
		t.Fatal(err)
	}
	req, ok := inner.(msg.UserRequest)
	if !ok {
		t.Fatalf("expected UserRequest, got %T", inner)
	}
	if req.UserReqID != "22" || req.Username != "Comarch" {
		t.Errorf("unexpected values: %+v", req)
	}
}

func TestDeserializeOrderFixture(t *testing.T) {
	data, _ := os.ReadFile("testdata/Order.xml")
	f, err := msg.Deserialize(string(data))
	if err != nil {
		t.Fatal(err)
	}
	inner, err := f.Unpack()
	if err != nil {
		t.Fatal(err)
	}
	order, ok := inner.(msg.NewOrderSingle)
	if !ok {
		t.Fatalf("expected NewOrderSingle, got %T", inner)
	}
	if order.ID != "1" || order.Acct != "00-55-012345" {
		t.Errorf("unexpected values: %+v", order)
	}
}

func TestDeserializeExecRptFixture(t *testing.T) {
	data, _ := os.ReadFile("testdata/ExecRpt.xml")
	f, err := msg.Deserialize(string(data))
	if err != nil {
		t.Fatal(err)
	}
	inner, err := f.Unpack()
	if err != nil {
		t.Fatal(err)
	}
	rpt, ok := inner.(msg.ExecutionReport)
	if !ok {
		t.Fatalf("expected ExecutionReport, got %T", inner)
	}
	if rpt.OrdID != "178803909" {
		t.Errorf("unexpected OrdID: %q", rpt.OrdID)
	}
}

func TestDeserializeSecurityListFixture(t *testing.T) {
	data, _ := os.ReadFile("testdata/SecurityList.xml")
	f, err := msg.Deserialize(string(data))
	if err != nil {
		t.Fatal(err)
	}
	inner, err := f.Unpack()
	if err != nil {
		t.Fatal(err)
	}
	sl, ok := inner.(msg.SecurityList)
	if !ok {
		t.Fatalf("expected SecurityList, got %T", inner)
	}
	if sl.SecL == nil || len(sl.SecL.Instrmts) != 2 {
		t.Errorf("expected 2 instruments, got: %+v", sl.SecL)
	}
}

func TestUnpackEmptyFixmlReturnsError(t *testing.T) {
	f := msg.NewFixml()
	_, err := f.Unpack()
	if err == nil {
		t.Fatal("expected error from empty Fixml.Unpack()")
	}
	var notFound *msg.FixmlElementNotFoundError
	if !errors.As(err, &notFound) {
		t.Errorf("expected FixmlElementNotFoundError, got %T: %v", err, err)
	}
}

func TestPackUserRequest(t *testing.T) {
	req := msg.UserRequest{UserReqID: "1", UserReqTyp: intPtr(1), Username: "BOS", Password: "BOS"}
	f := req.Pack()
	if f.V != "5.0" || f.R != "20080317" || f.S != "20080314" {
		t.Errorf("unexpected Fixml version fields: %+v", f)
	}
	if len(f.UserRequests) != 1 {
		t.Errorf("expected 1 UserRequest, got %d", len(f.UserRequests))
	}
}

func TestSerializeCompact(t *testing.T) {
	f := msg.NewFixml()
	f.UserRequests = []msg.UserRequest{{UserReqID: "1"}}
	s, err := msg.Serialize(f)
	if err != nil {
		t.Fatal(err)
	}
	// Compact XML must not contain newlines
	for _, c := range s {
		if c == '\n' {
			t.Errorf("Serialize must produce compact XML without newlines, got: %s", s)
			break
		}
	}
}

// Verify that xml.MarshalIndent produces valid XML that xml.Unmarshal can round-trip.
func TestSerializeIndented(t *testing.T) {
	f := msg.NewFixml()
	f.UserRequests = []msg.UserRequest{{UserReqID: "1"}}
	s, err := msg.SerializeIndented(f)
	if err != nil {
		t.Fatal(err)
	}
	var f2 msg.Fixml
	if err := xml.Unmarshal([]byte(s), &f2); err != nil {
		t.Fatalf("SerializeIndented produced invalid XML: %v\n%s", err, s)
	}
}
