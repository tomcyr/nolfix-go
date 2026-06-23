package msg

import "encoding/xml"

// MarketID values.
const (
	MarketIDNM = "NM" // cash market
	MarketIDDN = "DN" // derivatives market
)

// SecurityListRequestType values.
const (
	SecListReqTypeOneInstrmt     = 0
	SecListReqTypeOneTypeList    = 1
	SecListReqTypeFullList       = 4
	SecListReqTypeOneMarketType  = 5
)

// SecurityRequestResult values.
const (
	SecReqResultOK              = 0
	SecReqResultNoSec           = 1
	SecReqResultNoSecList       = 4
	SecReqResultNoSecListOfType = 5
)

// SecurityListRequest requests a list of securities (<SecListReq>).
type SecurityListRequest struct {
	XMLName    xml.Name     `xml:"SecListReq"`
	ReqID      string       `xml:"ReqID,attr,omitempty"`
	ListReqTyp *int         `xml:"ListReqTyp,attr,omitempty"`
	MktID      string       `xml:"MktID,attr,omitempty"`
	Instrmts   []Instrument `xml:"Instrmt"`
}

func (s SecurityListRequest) MsgName() string { return "SecListReq" }

// SecL is the wrapper element inside SecurityList containing instruments.
type SecL struct {
	XMLName  xml.Name     `xml:"SecL"`
	Instrmts []Instrument `xml:"Instrmt"`
}

// SecurityList is the response containing a list of securities (<SecList>).
type SecurityList struct {
	XMLName       xml.Name `xml:"SecList"`
	RptID         *int     `xml:"RptID,attr,omitempty"`
	ReqID         string   `xml:"ReqID,attr,omitempty"`
	MktID         string   `xml:"MktID,attr,omitempty"`
	TotNoReltdSym *int     `xml:"TotNoReltdSym,attr,omitempty"`
	ReqRslt       *int     `xml:"ReqRslt,attr,omitempty"`
	SecL          *SecL    `xml:"SecL"`
}

func (s SecurityList) MsgName() string { return "SecList" }
