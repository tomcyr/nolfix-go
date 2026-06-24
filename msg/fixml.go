package msg

import (
	"encoding/xml"
	"fmt"
)

// FixmlMessage is the common interface for all FIXML message types.
type FixmlMessage interface {
	MsgName() string
}

// Packer wraps a request in a Fixml envelope.
type Packer interface {
	Pack() *Fixml
}

// FixmlElementNotFoundError is returned when Fixml.Unpack() finds no inner message.
type FixmlElementNotFoundError struct {
	Msg string
}

func (e *FixmlElementNotFoundError) Error() string { return e.Msg }

// NewFixml returns a Fixml envelope pre-populated with standard version attributes.
func NewFixml() *Fixml {
	return &Fixml{V: "5.0", R: "20080317", S: "20080314"}
}

// Fixml is the root FIXML envelope element (<FIXML>).
type Fixml struct {
	XMLName xml.Name `xml:"FIXML"`
	// RawMessage is excluded from XML — set after receiving from wire.
	RawMessage string `xml:"-"`
	V          string `xml:"v,attr,omitempty"`
	R          string `xml:"r,attr,omitempty"`
	S          string `xml:"s,attr,omitempty"`

	// requests
	UserRequests     []UserRequest                 `xml:"UserReq"`
	MktDataReqs      []MarketDataRequest           `xml:"MktDataReq"`
	TrdgSesStatReqs  []TradingSessionStatusRequest `xml:"TrdgSesStatReq"`
	Orders           []NewOrderSingle              `xml:"Order"`
	OrdCxlRplcReqs   []OrderCancelReplaceRequest   `xml:"OrdCxlRplcReq"`
	OrdCxlReqs       []OrderCancelRequest          `xml:"OrdCxlReq"`
	OrdStatReqs      []OrderStatusRequest          `xml:"OrdStatReq"`
	SecListReqs      []SecurityListRequest         `xml:"SecListReq"`

	// responses
	UserResponses        []UserResponse                  `xml:"UserRsp"`
	BizMsgRejs           []BusinessMessageReject         `xml:"BizMsgRej"`
	MktDataIncrRefreshs  []MarketDataIncrementalRefresh  `xml:"MktDataInc"`
	MktDataFulls         []MarketDataSnapshotFullRefresh `xml:"MktDataFull"`
	MktDataReqRejs       []MarketDataRequestReject       `xml:"MktDataReqRej"`
	SecLists             []SecurityList                  `xml:"SecList"`
	TrdgSesStats         []TradingSessionStatus          `xml:"TrdgSesStat"`
	ExecRpts             []ExecutionReport               `xml:"ExecRpt"`
	OrdCxlRejs          []OrderCancelReject             `xml:"OrdCxlRej"`
	Heartbeats           []Heartbeat                     `xml:"Heartbeat"`
	NewsList             []News                          `xml:"News"`
	Statements           []Statement                     `xml:"Statement"`
	ApplMsgRpts          []ApplMsgRpt                    `xml:"AppIMsgRpt"`
}

// Unpack extracts the single inner message from this Fixml envelope.
// Returns FixmlElementNotFoundError if the envelope is empty.
func (f *Fixml) Unpack() (FixmlMessage, error) {
	switch {
	case len(f.UserRequests) > 0:
		return f.UserRequests[0], nil
	case len(f.MktDataReqs) > 0:
		return f.MktDataReqs[0], nil
	case len(f.TrdgSesStatReqs) > 0:
		return f.TrdgSesStatReqs[0], nil
	case len(f.Orders) > 0:
		return f.Orders[0], nil
	case len(f.OrdCxlRplcReqs) > 0:
		return f.OrdCxlRplcReqs[0], nil
	case len(f.OrdCxlReqs) > 0:
		return f.OrdCxlReqs[0], nil
	case len(f.OrdStatReqs) > 0:
		return f.OrdStatReqs[0], nil
	case len(f.SecListReqs) > 0:
		return f.SecListReqs[0], nil
	case len(f.UserResponses) > 0:
		return f.UserResponses[0], nil
	case len(f.BizMsgRejs) > 0:
		return f.BizMsgRejs[0], nil
	case len(f.MktDataIncrRefreshs) > 0:
		return f.MktDataIncrRefreshs[0], nil
	case len(f.MktDataFulls) > 0:
		return f.MktDataFulls[0], nil
	case len(f.MktDataReqRejs) > 0:
		return f.MktDataReqRejs[0], nil
	case len(f.SecLists) > 0:
		return f.SecLists[0], nil
	case len(f.TrdgSesStats) > 0:
		return f.TrdgSesStats[0], nil
	case len(f.ExecRpts) > 0:
		return f.ExecRpts[0], nil
	case len(f.OrdCxlRejs) > 0:
		return f.OrdCxlRejs[0], nil
	case len(f.Heartbeats) > 0:
		return f.Heartbeats[0], nil
	case len(f.NewsList) > 0:
		return f.NewsList[0], nil
	case len(f.Statements) > 0:
		return f.Statements[0], nil
	case len(f.ApplMsgRpts) > 0:
		return f.ApplMsgRpts[0], nil
	default:
		return nil, &FixmlElementNotFoundError{Msg: "FIXML Message has no content!"}
	}
}

// Serialize marshals a Fixml to compact XML with no indentation.
func Serialize(f *Fixml) (string, error) {
	b, err := xml.Marshal(f)
	if err != nil {
		return "", fmt.Errorf("serialize fixml: %w", err)
	}
	return string(b), nil
}

// Deserialize unmarshals a FIXML XML string into a Fixml struct.
func Deserialize(xmlStr string) (*Fixml, error) {
	var f Fixml
	if err := xml.Unmarshal([]byte(xmlStr), &f); err != nil {
		return nil, fmt.Errorf("deserialize fixml: %w", err)
	}
	return &f, nil
}

// SerializeIndented marshals a Fixml to indented XML for human readability.
func SerializeIndented(f *Fixml) (string, error) {
	b, err := xml.MarshalIndent(f, "", "  ")
	if err != nil {
		return "", fmt.Errorf("serialize fixml indented: %w", err)
	}
	return string(b), nil
}

// ── Pack() methods for all request types ─────────────────────────────────────

// Pack wraps a UserRequest in a new Fixml envelope.
func (u UserRequest) Pack() *Fixml {
	f := NewFixml()
	f.UserRequests = append(f.UserRequests, u)
	return f
}

// Pack wraps a MarketDataRequest in a new Fixml envelope.
func (m MarketDataRequest) Pack() *Fixml {
	f := NewFixml()
	f.MktDataReqs = append(f.MktDataReqs, m)
	return f
}

// Pack wraps a TradingSessionStatusRequest in a new Fixml envelope.
func (t TradingSessionStatusRequest) Pack() *Fixml {
	f := NewFixml()
	f.TrdgSesStatReqs = append(f.TrdgSesStatReqs, t)
	return f
}

// Pack wraps a NewOrderSingle in a new Fixml envelope.
func (o NewOrderSingle) Pack() *Fixml {
	f := NewFixml()
	f.Orders = append(f.Orders, o)
	return f
}

// Pack wraps an OrderCancelReplaceRequest in a new Fixml envelope.
func (o OrderCancelReplaceRequest) Pack() *Fixml {
	f := NewFixml()
	f.OrdCxlRplcReqs = append(f.OrdCxlRplcReqs, o)
	return f
}

// Pack wraps an OrderCancelRequest in a new Fixml envelope.
func (o OrderCancelRequest) Pack() *Fixml {
	f := NewFixml()
	f.OrdCxlReqs = append(f.OrdCxlReqs, o)
	return f
}

// Pack wraps an OrderStatusRequest in a new Fixml envelope.
func (o OrderStatusRequest) Pack() *Fixml {
	f := NewFixml()
	f.OrdStatReqs = append(f.OrdStatReqs, o)
	return f
}

// Pack wraps a SecurityListRequest in a new Fixml envelope.
func (s SecurityListRequest) Pack() *Fixml {
	f := NewFixml()
	f.SecListReqs = append(f.SecListReqs, s)
	return f
}
