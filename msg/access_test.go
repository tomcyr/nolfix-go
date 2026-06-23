package msg_test

import (
	"encoding/xml"
	"testing"
	"github.com/tomcyr/nolfix/msg"
)

func TestUserRequestRoundtrip(t *testing.T) {
	n := 1
	orig := msg.UserRequest{
		UserReqID:  "22",
		UserReqTyp: &n,
		Username:   "Comarch",
		Password:   "Comarch",
	}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.UserRequest
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.UserReqID != orig.UserReqID || *got.UserReqTyp != *orig.UserReqTyp ||
		got.Username != orig.Username || got.Password != orig.Password {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}

func TestUserResponseRoundtrip(t *testing.T) {
	stat := 1
	orig := msg.UserResponse{UserReqID: "22", Username: "Comarch", UserStat: &stat}
	b, err := xml.Marshal(&orig)
	if err != nil {
		t.Fatal(err)
	}
	var got msg.UserResponse
	if err := xml.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.Username != orig.Username || *got.UserStat != *orig.UserStat {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", got, orig)
	}
}
