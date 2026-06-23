package msg

import "encoding/xml"

// Instrument identifies a financial instrument.
type Instrument struct {
	XMLName xml.Name `xml:"Instrmt"`
	Sym     string   `xml:"Sym,attr,omitempty"`
	ID      string   `xml:"ID,attr,omitempty"`
	Src     *int     `xml:"Src,attr,omitempty"`
	CFI     string   `xml:"CFI,attr,omitempty"`
	SecGrp  string   `xml:"SecGrp,attr,omitempty"`
}

// BusinessMessageReject is a generic negative response (<BizMsgRej>).
type BusinessMessageReject struct {
	XMLName   xml.Name `xml:"BizMsgRej"`
	RefMsgTyp string   `xml:"RefMsgTyp,attr,omitempty"`
	BizRejRsn *int     `xml:"BizRejRsn,attr,omitempty"`
	Txt       string   `xml:"Txt,attr,omitempty"`
}

func (m BusinessMessageReject) MsgName() string { return "BizMsgRej" }

// RefMsgType values for BusinessMessageReject.RefMsgTyp.
const (
	RefMsgTypeLoginLogout   = "BE"
	RefMsgTypeOrderNew      = "D"
	RefMsgTypeOrderCancel   = "F"
	RefMsgTypeOrderChange   = "G"
	RefMsgTypeOrderStatus   = "H"
	RefMsgTypeOnlineQuotes  = "V"
	RefMsgTypeSessionStatus = "g"
)

// BusinessRejectReason values for BusinessMessageReject.BizRejRsn.
const (
	BizRejRsnOther          = 0
	BizRejRsnUnknownID      = 1
	BizRejRsnUnknownInstr   = 2
	BizRejRsnUnknownMsgType = 3
	BizRejRsnAccessDenied   = 4
	BizRejRsnXMLError       = 5
	BizRejRsnNotAuthorized  = 6
	BizRejRsnNoCommunication = 7
)

// MarketDepth values for subscription requests.
const (
	MarketDepthAllOffers   = 0
	MarketDepthBestOffer   = 1
	MarketDepthBest5Offers = 2
)
