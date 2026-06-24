package msg

import "encoding/xml"

// SessSubReqTyp values.
const (
	SessSubReqTypOnline  = '1'
	SessSubReqTypOffline = '2'
)

// SessionPhase values for TradingSessionStatus.SesSub.
const (
	SessionPhaseConsultation          = "C" // Konsultacja nadzoru
	SessionPhasePreOpen               = "P" // Przed otwarciem
	SessionPhaseIntervention          = "E" // Interwencja
	SessionPhaseOpening               = "O" // Otwarcie
	SessionPhasePlayOff               = "R" // Dogrywka
	SessionPhaseContinuous            = "S" // Sesja notowań ciągłych
	SessionPhaseRegulatoryIntervention = "N" // Interwencja nadzoru
	SessionPhaseClosingConsultation   = "F" // Konsultacja nadzoru (zamknięcie)
	SessionPhasePostSession           = "B" // Po sesji
)

// SessionState values for TradingSessionStatus.Stat.
const (
	SessionStateBalancing    = "AR"
	SessionStateSecInUse     = "A"
	SessionStateSecLocked    = "AG"
	SessionStateNoPlayOff    = "IR"
	SessionStateClosedAS     = "AS"
	SessionStateClosedIS     = "IS"
	SessionStateClosedI      = "I"
)

// SessionRejectReason values for TradingSessionStatus.StatRejRsn.
const (
	SessionRejectReasonWrongSessionID = 1
	SessionRejectReasonOther          = 99
)

// TradingSessionStatusRequest requests session status (<TrdgSesStatReq>).
type TradingSessionStatusRequest struct {
	XMLName   xml.Name `xml:"TrdgSesStatReq"`
	ReqID     string   `xml:"ReqID,attr,omitempty"`
	SubReqTyp string   `xml:"SubReqTyp,attr,omitempty"`
}

func (t TradingSessionStatusRequest) MsgName() string { return "TrdgSesStatReq" }

// TradingSessionStatus is the response with current session info (<TrdgSesStat>).
type TradingSessionStatus struct {
	XMLName    xml.Name     `xml:"TrdgSesStat"`
	ReqID      string       `xml:"ReqID,attr,omitempty"`
	SesID      string       `xml:"SesID,attr,omitempty"`
	SesSub     string       `xml:"SesSub,attr,omitempty"`
	Stat       string       `xml:"Stat,attr,omitempty"`
	StatRejRsn *int         `xml:"StatRejRsn,attr,omitempty"`
	Instrmts   []Instrument `xml:"Instrmt"`
}

func (t TradingSessionStatus) MsgName() string { return "TrdgSesStat" }
