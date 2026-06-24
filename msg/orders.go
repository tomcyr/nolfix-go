package msg

import "encoding/xml"

// SideType values.
const (
	SideTypeBuy  = '1'
	SideTypeSell = '2'
)

// OrderType values.
const (
	OrderTypePKC       = '1'
	OrderTypeLimit     = 'L'
	OrderTypeStopLoss  = '3'
	OrderTypeStopLimit = '4'
	OrderTypePCR       = 'K'
	OrderTypePEG       = 'E'
	OrderTypePEGL      = 'G'
)

// TimeInForce values.
const (
	TimeInForceDay          = '0'
	TimeInForceGTC          = '1'
	TimeInForceBeforeOpen   = '2'
	TimeInForceWiA          = '3'
	TimeInForceWuA          = '4'
	TimeInForceGoodTillDate = '6'
	TimeInForceAtClose      = '7'
	TimeInForceNextFixing   = 'f'
	TimeInForceGoodTillTime = 't'
)

// OrderStatus values for ExecutionReport.Stat.
const (
	OrderStatusNew             = '0'
	OrderStatusPartiallyFilled = '1'
	OrderStatusFilled          = '2'
	OrderStatusCanceled        = '4'
	OrderStatusPendingCancel   = '6'
	OrderStatusRejected        = '8'
	OrderStatusPendingNew      = 'A'
	OrderStatusExpired         = 'C'
)

// ExecType values for ExecutionReport.ExecTyp.
const (
	ExecTypeNew            = '0'
	ExecTypeTransaction    = 'F'
	ExecTypeCancellation   = '4'
	ExecTypeModified       = 'E'
	ExecTypeBeingCancelled = '6'
	ExecTypeRejected       = '8'
	ExecTypeOrderStatus    = 'I'
)

// CommType values for ExecutionReport commission data.
const (
	CommTypeFactor  = '1'
	CommTypePercent = '2'
	CommTypeAbs     = '3'
)

// DisplayInstruction holds the visible quantity for an order (<DsplyInstr>).
type DisplayInstruction struct {
	XMLName    xml.Name `xml:"DsplyInstr"`
	DisplayQty float32  `xml:"DisplayQty,attr"`
}

// OrderQtyData holds order quantity (<OrdQty>).
type OrderQtyData struct {
	XMLName xml.Name `xml:"OrdQty"`
	Qty     float32  `xml:"Qty,attr"`
}

// TriggeringInstruction holds DDM+ trigger data (<TrgrInstr>).
type TriggeringInstruction struct {
	XMLName   xml.Name `xml:"TrgrInstr"`
	TrgrTyp   string   `xml:"TrgrTyp,attr,omitempty"`
	TrgrActn  string   `xml:"TrgrActn,attr,omitempty"`
	TrgrPx    *float32 `xml:"TrgrPx,attr,omitempty"`
	TrgrPxTyp string   `xml:"TrgrPxTyp,attr,omitempty"`
}

// NewOrderSingle is a new order request (<Order>).
type NewOrderSingle struct {
	XMLName    xml.Name               `xml:"Order"`
	ID         string                 `xml:"ID,attr,omitempty"`
	TrdDt      string                 `xml:"TrdDt,attr,omitempty"`
	Acct       string                 `xml:"Acct,attr,omitempty"`
	MinQty     *float32               `xml:"MinQty,attr,omitempty"`
	Side       string                 `xml:"Side,attr,omitempty"`
	TxnTm      string                 `xml:"TxnTm,attr,omitempty"`
	OrdTyp     string                 `xml:"OrdTyp,attr,omitempty"`
	Px         string                 `xml:"Px,attr,omitempty"`
	StopPx     *float32               `xml:"StopPx,attr,omitempty"`
	Ccy        string                 `xml:"Ccy,attr,omitempty"`
	TmInForce  string                 `xml:"TmInForce,attr,omitempty"`
	ExpireDt   string                 `xml:"ExpireDt,attr,omitempty"`
	DefPayTyp  string                 `xml:"DefPayTyp,attr,omitempty"`
	DsplyInstr *DisplayInstruction    `xml:"DsplyInstr"`
	Instrmt    *Instrument            `xml:"Instrmt"`
	OrdQty     *OrderQtyData          `xml:"OrdQty"`
	TrgrInstr  *TriggeringInstruction `xml:"TrgrInstr"`
}

func (o NewOrderSingle) MsgName() string { return "Order" }

// OrderCancelReplaceRequest is an order modification request (<OrdCxlRplcReq>).
type OrderCancelReplaceRequest struct {
	XMLName    xml.Name               `xml:"OrdCxlRplcReq"`
	ID         string                 `xml:"ID,attr,omitempty"`
	OrigID     string                 `xml:"OrigID,attr,omitempty"`
	OrdID      string                 `xml:"OrdID,attr,omitempty"`
	OrdID2     string                 `xml:"OrdID2,attr,omitempty"`
	TrdDt      string                 `xml:"TrdDt,attr,omitempty"`
	Acct       string                 `xml:"Acct,attr,omitempty"`
	MinQty     *float32               `xml:"MinQty,attr,omitempty"`
	Side       string                 `xml:"Side,attr,omitempty"`
	TxnTm      string                 `xml:"TxnTm,attr,omitempty"`
	OrdTyp     string                 `xml:"OrdTyp,attr,omitempty"`
	Px         string                 `xml:"Px,attr,omitempty"`
	StopPx     *float32               `xml:"StopPx,attr,omitempty"`
	Ccy        string                 `xml:"Ccy,attr,omitempty"`
	TmInForce  string                 `xml:"TmInForce,attr,omitempty"`
	ExpireDt   string                 `xml:"ExpireDt,attr,omitempty"`
	ExpireTm   string                 `xml:"ExpireTm,attr,omitempty"`
	Txt        string                 `xml:"Txt,attr,omitempty"`
	DefPayTyp  string                 `xml:"DefPayTyp,attr,omitempty"`
	DsplyInstr *DisplayInstruction    `xml:"DsplyInstr"`
	Instrmt    *Instrument            `xml:"Instrmt"`
	OrdQty     *OrderQtyData          `xml:"OrdQty"`
	TrgrInstr  *TriggeringInstruction `xml:"TrgrInstr"`
}

func (o OrderCancelReplaceRequest) MsgName() string { return "OrdCxlRplcReq" }

// OrderCancelRequest is an order cancellation request (<OrdCxlReq>).
type OrderCancelRequest struct {
	XMLName xml.Name      `xml:"OrdCxlReq"`
	ID      string        `xml:"ID,attr,omitempty"`
	OrigID  string        `xml:"OrigID,attr,omitempty"`
	OrdID   string        `xml:"OrdID,attr,omitempty"`
	OrdID2  string        `xml:"OrdID2,attr,omitempty"`
	Acct    string        `xml:"Acct,attr,omitempty"`
	Side    string        `xml:"Side,attr,omitempty"`
	TxnTm   string        `xml:"TxnTm,attr,omitempty"`
	Txt     string        `xml:"Txt,attr,omitempty"`
	Instrmt *Instrument   `xml:"Instrmt"`
	OrdQty  *OrderQtyData `xml:"OrdQty"`
}

func (o OrderCancelRequest) MsgName() string { return "OrdCxlReq" }

// OrderStatusRequest queries for order status (<OrdStatReq>).
type OrderStatusRequest struct {
	XMLName   xml.Name    `xml:"OrdStatReq"`
	OrdID     string      `xml:"OrdID,attr,omitempty"`
	OrigID    string      `xml:"OrigID,attr,omitempty"`
	ID        string      `xml:"ID,attr,omitempty"`
	StatReqID string      `xml:"StatReqID,attr,omitempty"`
	Acct      string      `xml:"Acct,attr,omitempty"`
	Side      string      `xml:"Side,attr,omitempty"`
	Instrmt   *Instrument `xml:"Instrmt"`
}

func (o OrderStatusRequest) MsgName() string { return "OrdStatReq" }

// CommissionData holds per-transaction commission information (<Comm>).
type CommissionData struct {
	XMLName xml.Name `xml:"Comm"`
	Comm    float32  `xml:"Comm,attr"`
	CommTyp string   `xml:"CommTyp,attr,omitempty"`
}

// ExecutionReport is an async response to order messages (<ExecRpt>).
type ExecutionReport struct {
	XMLName     xml.Name                `xml:"ExecRpt"`
	OrdID       string                  `xml:"OrdID,attr,omitempty"`
	OrdID2      string                  `xml:"OrdID2,attr,omitempty"`
	ID          string                  `xml:"ID,attr,omitempty"`
	StatReqID   string                  `xml:"StatReqID,attr,omitempty"`
	ExecID      string                  `xml:"ExecID,attr,omitempty"`
	ExecTyp     string                  `xml:"ExecTyp,attr,omitempty"`
	Stat        string                  `xml:"Stat,attr,omitempty"`
	RejRsn      *int                    `xml:"RejRsn,attr,omitempty"`
	Acct        string                  `xml:"Acct,attr,omitempty"`
	Side        string                  `xml:"Side,attr,omitempty"`
	OrdTyp      string                  `xml:"OrdTyp,attr,omitempty"`
	Px          string                  `xml:"Px,attr,omitempty"`
	StopPx      *float32                `xml:"StopPx,attr,omitempty"`
	Ccy         string                  `xml:"Ccy,attr,omitempty"`
	TmInForce   string                  `xml:"TmInForce,attr,omitempty"`
	ExpireDt    string                  `xml:"ExpireDt,attr,omitempty"`
	LastPx      *float32                `xml:"LastPx,attr,omitempty"`
	LastQty     *float32                `xml:"LastQty,attr,omitempty"`
	LeavesQty   *float32                `xml:"LeavesQty,attr,omitempty"`
	CumQty      *float32                `xml:"CumQty,attr,omitempty"`
	TxnTm       string                  `xml:"TxnTm,attr,omitempty"`
	NetMny      *float32                `xml:"NetMny,attr,omitempty"`
	MinQty      *float32                `xml:"MinQty,attr,omitempty"`
	Txt         string                  `xml:"Txt,attr,omitempty"`
	DefPayTyp   string                  `xml:"DefPayTyp,attr,omitempty"`
	Instrmts    []Instrument            `xml:"Instrmt"`
	OrdQty      *OrderQtyData           `xml:"OrdQty"`
	DsplyInstrs []DisplayInstruction    `xml:"DsplyInstr"`
	Comms       []CommissionData        `xml:"Comm"`
	TrgrInstrs  []TriggeringInstruction `xml:"TrgrInstr"`
}

func (e ExecutionReport) MsgName() string { return "ExecRpt" }
