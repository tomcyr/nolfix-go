# BOSSA NOL3 client

Go client library for the **Nol3 API** — the FIXML/TCP interface exposed by the BOSSA Polish stock broker. This is a port of [nolfix4java](https://github.com/tomcyr/nolfix4java).

## Requirements

- Go 1.26+
- A running Nol3 application (127.0.0.1:24444 sync, 127.0.0.1:24445 async by default)

## Installation

```bash
go get github.com/tomcyr/nolfix-go
```

## Quick start

```go
factory := nolfix.NewRequestFactory(nolfix.UUIDGenerator{})
client  := nolfix.NewNolRequestReplyClient(nolfix.Host, nolfix.SyncPort)

// Login
req := factory.CreateUserRequest()
req.UserReqTyp = intPtr(msg.UserReqTypLogin)
req.Username   = nolfix.UserName
req.Password   = nolfix.UserPasswd

resp, err := client.Send(req.Pack())
if err != nil {
    // inspect with errors.As — see Error handling below
}
inner, _ := resp.Unpack() // returns *msg.UserResponse
```

See [cmd/example/main.go](cmd/example/main.go) for a runnable demo covering login, instruments, orders, market data, and async receive.

## Package structure

```
nolfix/
├── settings.go        # Default connection constants (Host, SyncPort, AsyncPort, …)
├── limits.go          # API limits (max securities in filter, max orders/day, …)
├── id.go              # IdGenerator interface + UUIDGenerator
├── client.go          # NolClient — raw TCP send/receive (Sender + Receiver interfaces)
├── request_reply.go   # NolRequestReplyClient — one-call sync send+receive
├── factory.go         # RequestFactory — creates pre-stamped request structs
├── errors.go          # Error types
└── msg/
    ├── fixml.go       # Fixml envelope — Serialize, Deserialize, Unpack, Pack
    ├── access.go      # UserRequest, UserResponse
    ├── orders.go      # NewOrderSingle, cancel, replace, status, ExecutionReport
    ├── filter.go      # MarketDataRequest, incremental/full refresh, reject
    ├── securities.go  # SecurityListRequest, SecurityList
    ├── session.go     # TradingSessionStatusRequest, TradingSessionStatus
    ├── common.go      # Instrument, BusinessMessageReject
    └── misc.go        # Heartbeat, News, Statement, ApplMsgRpt
```

## Usage

### Synchronous request-reply

`NolRequestReplyClient.Send` opens a TCP connection to the sync port, sends the message, reads the response, and closes the connection — all in one call.

```go
client := nolfix.NewNolRequestReplyClient(nolfix.Host, nolfix.SyncPort)

// Security list
req := factory.CreateSecurityListRequest()
req.MktID = msg.MarketIDNM

resp, err := client.Send(req.Pack())
sl, _ := resp.Unpack()               // returns msg.SecurityList
```

### Orders

```go
src := 4
order := factory.CreateNewOrderSingle()
order.TrdDt    = "20130502"
order.Acct     = "00-55-123456"
order.Side     = string(msg.SideTypeBuy)
order.OrdTyp   = string(msg.OrderTypeLimit)
order.Ccy      = "PLN"
order.TmInForce = string(msg.TimeInForceDay)
order.Instrmt  = &msg.Instrument{Sym: "COMARCH", ID: "PLCOMAR00012", Src: &src}
order.OrdQty   = &msg.OrderQtyData{Qty: 10.0}

resp, err := client.Send(order.Pack())
```

### Market data filter

```go
depth := msg.MarketDepthBestOffer
req := factory.CreateMarketDataRequest()
req.SubReqTyp = string(msg.SubReqTypAddToFilter)
req.MktDepth  = &depth
req.Reqs = []msg.MdReqGrp{
    {Typ: string(msg.EntryTypeOffer)},
    {Typ: string(msg.EntryTypeLastTrade)},
}
req.InstReq = &msg.InstReq{
    Instrmts: []msg.Instrument{{Sym: "COMARCH"}},
}

client.Send(req.Pack())
```

### Async receiver

Use `NolClient` directly on the async port. Push-notifications arrive continuously once subscribed.

```go
client := nolfix.NewNolClient(nolfix.Host, nolfix.AsyncPort)
client.Connect()

stop := make(chan os.Signal, 1)
signal.Notify(stop, os.Interrupt)

go func() {
    for {
        f, err := client.Receive()
        if err != nil { return }
        inner, _ := f.Unpack()
        fmt.Println(inner)
    }
}()
<-stop
client.Disconnect()
```

### Custom ID generator

```go
type MyIDGen struct{ n int }
func (g *MyIDGen) NextID() string { g.n++; return strconv.Itoa(g.n) }

factory := nolfix.NewRequestFactory(&MyIDGen{})
```

## Error handling

```go
var bizErr *nolfix.BusinessMessageRejectError
var mdrErr *nolfix.MarketDataRequestRejectError
var conErr *nolfix.NolConnectionError

switch {
case errors.As(err, &bizErr):
    fmt.Println("rejected:", bizErr.Message.Txt)
case errors.As(err, &mdrErr):
    fmt.Println("market data rejected:", mdrErr.Message.ReqRejResn)
case errors.As(err, &conErr):
    fmt.Println("connection error:", conErr.Cause)
}
```

| Error type | When |
|---|---|
| `NolConnectionError` | TCP connect / send / receive fails |
| `BusinessMessageRejectError` | Server returns `<BizMsgRej>` |
| `MarketDataRequestRejectError` | Server returns `<MktDataReqRej>` |

## Wire protocol

The Nol3 protocol is **asymmetric**:

- **Send** (client → server): ASCII decimal length string, then XML payload + null byte.
- **Receive** (server → client): 4-byte header (little-endian uint24 in bytes 0–2 = message size), then XML payload + null byte.

## Default connection settings

| Constant | Value |
|---|---|
| `Host` | `127.0.0.1` |
| `SyncPort` | `24444` |
| `AsyncPort` | `24445` |
| `UserName` | `BOS` |
| `UserPasswd` | `BOS` |
