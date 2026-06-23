package msg

import "encoding/xml"

// EntryType values for market data entry types.
const (
	EntryTypeBid                     = '0'
	EntryTypeOffer                   = '1'
	EntryTypeLastTrade               = '2'
	EntryTypeTradeVolume             = 'B'
	EntryTypeOpenInterest            = 'C'
	EntryTypeIndexVolume             = '3'
	EntryTypeOpeningPrice            = '4'
	EntryTypeClosingPrice            = '5'
	EntryTypeTradingSessionHighPrice = '7'
	EntryTypeTradingSessionLowPrice  = '8'
	EntryTypeReferencePrice          = 'r'
)

// SubscriptionRequestType values.
const (
	SubReqTypAddToFilter = '1'
	SubReqTypClearFilter = '2'
)

// MdUpdateAction values for MdIncGrp.UpdtAct.
const (
	MdUpdateActionNew    = '0'
	MdUpdateActionChange = '1'
	MdUpdateActionDelete = '2'
)

// MDReqRejReason values for MarketDataRequestReject.ReqRejResn.
const (
	MDReqRejReasonUnknownSecurity      = '0'
	MDReqRejReasonIDDuplicate          = '1'
	MDReqRejReasonErrorSubscrReq       = '4'
	MDReqRejReasonUnsupportedNumOffers = '5'
	MDReqRejReasonErrorMdUpdateType    = '6'
	MDReqRejReasonUnsupportedQuotes    = '8'
)

// MdReqGrp is a single entry type in a market data request (<req>).
type MdReqGrp struct {
	XMLName xml.Name `xml:"req"`
	Typ     string   `xml:"Typ,attr"`
}

// InstReq wraps a list of instruments in a market data request (<InstReq>).
type InstReq struct {
	XMLName  xml.Name     `xml:"InstReq"`
	Instrmts []Instrument `xml:"Instrmt"`
}

// MarketDataRequest subscribes/unsubscribes from market data (<MktDataReq>).
type MarketDataRequest struct {
	XMLName   xml.Name   `xml:"MktDataReq"`
	ReqID     string     `xml:"ReqID,attr,omitempty"`
	SubReqTyp string     `xml:"SubReqTyp,attr,omitempty"`
	MktDepth  *int       `xml:"MktDepth,attr,omitempty"`
	Reqs      []MdReqGrp `xml:"req"`
	InstReq   *InstReq   `xml:"InstReq"`
}

func (m MarketDataRequest) MsgName() string { return "MktDataReq" }

// MdIncGrp is a single incremental market data update entry (<Inc>).
type MdIncGrp struct {
	XMLName   xml.Name     `xml:"Inc"`
	UpdtAct   string       `xml:"UpdtAct,attr,omitempty"`
	Typ       string       `xml:"Typ,attr,omitempty"`
	Px        string       `xml:"Px,attr,omitempty"`
	MdPxLv    *int         `xml:"MDPxLvl,attr,omitempty"`
	Ccy       *float32     `xml:"Ccy,attr,omitempty"`
	Sz        *float32     `xml:"Sz,attr,omitempty"`
	NumOfOrds *int         `xml:"NumOfOrds,attr,omitempty"`
	Dt        string       `xml:"Dt,attr,omitempty"`
	Tm        string       `xml:"Tm,attr,omitempty"`
	Tov       *float32     `xml:"Tov,attr,omitempty"`
	Instrmts  []Instrument `xml:"Instrmt"`
}

// MarketDataIncrementalRefresh is an async market data update (<MktDataInc>).
type MarketDataIncrementalRefresh struct {
	XMLName xml.Name   `xml:"MktDataInc"`
	ReqID   string     `xml:"MDReqID,attr,omitempty"`
	Incs    []MdIncGrp `xml:"Inc"`
}

func (m MarketDataIncrementalRefresh) MsgName() string { return "MktDataInc" }

// MarketDataSnapshotFullRefresh is the response to a filter add/clear (<MktDataFull>).
type MarketDataSnapshotFullRefresh struct {
	XMLName xml.Name `xml:"MktDataFull"`
	ReqID   string   `xml:"ReqID,attr,omitempty"`
}

func (m MarketDataSnapshotFullRefresh) MsgName() string { return "MktDataFull" }

// MarketDataRequestReject is the response when a market data request is rejected (<MktDataReqRej>).
type MarketDataRequestReject struct {
	XMLName    xml.Name `xml:"MktDataReqRej"`
	ReqID      string   `xml:"ReqID,attr,omitempty"`
	ReqRejResn string   `xml:"ReqRejResn,attr,omitempty"`
	Txt        string   `xml:"Txt,attr,omitempty"`
}

func (m MarketDataRequestReject) MsgName() string { return "MktDataReqRej" }
