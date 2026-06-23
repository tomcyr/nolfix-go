// /home/tomcyr/code/nolfix/cmd/example/main.go
package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	nolfix "github.com/tomcyr/nolfix-go"
	"github.com/tomcyr/nolfix-go/msg"
)

func main() {
	factory := nolfix.NewRequestFactory(nolfix.UUIDGenerator{})
	syncClient := nolfix.NewNolRequestReplyClient(nolfix.Host, nolfix.SyncPort)

	login(factory, syncClient)
	// Uncomment to run other demos:
	// printInstruments(factory, syncClient)
	// disableAsyncAndPrintSession(factory, syncClient)
	// addInstrumentToFilter(factory, syncClient)
	// clearFilter(factory, syncClient)
	// newOrder(factory, syncClient)
	// cancelOrder(factory, syncClient)
	// startAsyncReceiver()
}

func login(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Logging in...")
	req := factory.CreateUserRequest()
	req.UserReqTyp = intPtr(msg.UserReqTypLogin)
	req.Username = nolfix.UserName
	req.Password = nolfix.UserPasswd

	resp, err := client.Send(req.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, err := resp.Unpack()
	if err != nil {
		fmt.Println("Unpack error:", err)
		return
	}
	fmt.Println("Response:", inner)
}

func printInstruments(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Getting instruments...")
	req := factory.CreateSecurityListRequest()
	req.ListReqTyp = intPtr(msg.SecListReqTypeOneMarketType)
	req.MktID = msg.MarketIDNM

	resp, err := client.Send(req.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, err := resp.Unpack()
	if err != nil {
		fmt.Println("Unpack error:", err)
		return
	}
	sl, ok := inner.(msg.SecurityList)
	if !ok || sl.SecL == nil {
		fmt.Println("Unexpected response:", inner)
		return
	}
	for _, instr := range sl.SecL.Instrmts {
		fmt.Println(instr)
	}
}

func disableAsyncAndPrintSession(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Getting session state...")
	req := factory.CreateTradingSessionStatusRequest()
	req.SubReqTyp = string(msg.SessSubReqTypOffline)

	resp, err := client.Send(req.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, _ := resp.Unpack()
	fmt.Println("Session status:", inner)
}

func addInstrumentToFilter(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Adding instrument to filter...")
	depth := msg.MarketDepthBestOffer
	req := factory.CreateMarketDataRequest()
	req.SubReqTyp = string(msg.SubReqTypAddToFilter)
	req.MktDepth = &depth
	req.Reqs = []msg.MdReqGrp{
		{Typ: string(msg.EntryTypeOffer)},
		{Typ: string(msg.EntryTypeLastTrade)},
		{Typ: string(msg.EntryTypeTradeVolume)},
		{Typ: string(msg.EntryTypeOpeningPrice)},
		{Typ: string(msg.EntryTypeIndexVolume)},
	}
	req.InstReq = &msg.InstReq{
		Instrmts: []msg.Instrument{{Sym: "COMARCH"}},
	}

	resp, err := client.Send(req.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, _ := resp.Unpack()
	fmt.Println("MarketDataSnapshotFullRefresh:", inner)
}

func clearFilter(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Clearing filter...")
	req := factory.CreateMarketDataRequest()
	req.SubReqTyp = string(msg.SubReqTypClearFilter)

	resp, err := client.Send(req.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, _ := resp.Unpack()
	fmt.Println("MarketDataSnapshotFullRefresh:", inner)
}

func newOrder(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Placing order...")
	src := 4
	qty := float32(1.0)
	order := factory.CreateNewOrderSingle()
	order.TrdDt = "20130502"
	order.Acct = "00-55-123456"
	order.Side = string(msg.SideTypeBuy)
	order.TxnTm = "20130502-14:04:00"
	order.OrdTyp = string(msg.OrderTypePEG)
	order.Ccy = "PLN"
	order.TmInForce = string(msg.TimeInForceDay)
	order.Instrmt = &msg.Instrument{Sym: "COMARCH", ID: "PLCOMAR00012", Src: &src}
	order.OrdQty = &msg.OrderQtyData{Qty: qty}

	resp, err := client.Send(order.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, _ := resp.Unpack()
	fmt.Println("ExecutionReport:", inner)
}

func cancelOrder(factory *nolfix.RequestFactory, client *nolfix.NolRequestReplyClient) {
	fmt.Println("Cancelling order...")
	src := 4
	qty := float32(1.0)
	cancel := factory.CreateOrderCancelRequest()
	cancel.Acct = "00-55-123456"
	cancel.OrdID = "335753983"
	cancel.Side = string(msg.SideTypeBuy)
	cancel.TxnTm = "20130502-13:56:00"
	cancel.Instrmt = &msg.Instrument{Sym: "COMARCH", ID: "PLCOMAR00012", Src: &src}
	cancel.OrdQty = &msg.OrderQtyData{Qty: qty}

	resp, err := client.Send(cancel.Pack())
	if err != nil {
		printErr(err)
		return
	}
	inner, _ := resp.Unpack()
	fmt.Println("ExecutionReport:", inner)
}

func startAsyncReceiver() {
	fmt.Println("Starting async receiver (Ctrl+C to stop)...")
	client := nolfix.NewNolClient(nolfix.Host, nolfix.AsyncPort)
	if err := client.Connect(); err != nil {
		fmt.Println("Connect error:", err)
		os.Exit(1)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	msgCh := make(chan *msg.Fixml)
	errCh := make(chan error, 1)
	go func() {
		for {
			resp, err := client.Receive()
			if err != nil {
				errCh <- err
				return
			}
			msgCh <- resp
		}
	}()

	for {
		select {
		case f := <-msgCh:
			inner, err := f.Unpack()
			if err != nil {
				fmt.Println("Unpack error:", err)
			} else {
				fmt.Println("Async message:", inner)
			}
		case err := <-errCh:
			fmt.Println("Receive error:", err)
			return
		case <-stop:
			fmt.Println("Stopping...")
			client.Disconnect() //nolint:errcheck
			return
		}
	}
}

func printErr(err error) {
	var bizErr *nolfix.BusinessMessageRejectError
	var mdrErr *nolfix.MarketDataRequestRejectError
	switch {
	case errors.As(err, &bizErr):
		fmt.Println("BusinessMessageReject:", bizErr.Message)
	case errors.As(err, &mdrErr):
		fmt.Println("MarketDataRequestReject:", mdrErr.Message)
	default:
		fmt.Println("Error:", err)
	}
}

func intPtr(v int) *int { return &v }
