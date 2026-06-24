package msg

import "encoding/xml"

// Heartbeat is an empty keep-alive message (<Heartbeat>).
type Heartbeat struct {
	XMLName xml.Name `xml:"Heartbeat"`
}

func (h Heartbeat) MsgName() string { return "Heartbeat" }

// News is a broker news message with text body (<News>).
type News struct {
	XMLName  xml.Name `xml:"News"`
	OrigTm   string   `xml:"OrigTm,attr,omitempty"`
	Headline string   `xml:"Headline,attr,omitempty"`
	Message  string   `xml:",chardata"`
}

func (n News) MsgName() string { return "News" }

// ApplMsgRpt reports message delivery delay (<ApplMsgRpt>).
type ApplMsgRpt struct {
	XMLName   xml.Name `xml:"ApplMsgRpt"`
	ApplRepID string   `xml:"ApplRepID,attr,omitempty"`
	Txt       string   `xml:"Txt,attr,omitempty"`
}

func (a ApplMsgRpt) MsgName() string { return "ApplMsgRpt" }

// Fund holds a named financial value in a Statement (<Fund>).
type Fund struct {
	XMLName xml.Name `xml:"Fund"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:"value,attr,omitempty"`
}

// Position holds an asset position in a Statement (<Position>).
type Position struct {
	XMLName  xml.Name     `xml:"Position"`
	Acc110   string       `xml:"Acc110,attr,omitempty"`
	Acc120   string       `xml:"Acc120,attr,omitempty"`
	Instrmts []Instrument `xml:"Instrmt"`
}

// Statement is a non-FIX BOS account statement message (<Statement>).
type Statement struct {
	XMLName   xml.Name   `xml:"Statement"`
	Acct      string     `xml:"Acct,attr,omitempty"`
	AcctType  string     `xml:"type,attr,omitempty"`
	Ike       string     `xml:"ike,attr,omitempty"`
	Funds     []Fund     `xml:"Fund"`
	Positions []Position `xml:"Position"`
}

func (s Statement) MsgName() string { return "Statement" }

// AccountType values for Statement.AcctType.
const (
	AccountTypeCash       = 'M'
	AccountTypeDerivative = 'P'
)
