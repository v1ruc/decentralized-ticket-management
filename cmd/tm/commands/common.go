package commands

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth/backend"
	"golang.org/x/crypto/sha3"
)

var (
	signUpKey [32]byte
)

func init() {
	copy(signUpKey[:], "signup")
}

type signUpData struct {
	ParticipantDIDAddress common.Address `json:"participant_did_address"`
	FullName              string         `json:"participant_full_name"`
}

func (d signUpData) ToJSONBytes() []byte {
	bs, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return bs
}

func (d *signUpData) FromJSONBytes(bs []byte) error {
	return json.Unmarshal(bs, d)
}

type ticketData struct {
	EventDIDAddress       common.Address `json:"event_did_address"`
	ParticipantDIDAddress common.Address `json:"participant_did_address"`
}

func (d ticketData) ToJSONBytes() []byte {
	bs, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return bs
}

type qrCode struct {
	TicketData      []byte `json:"ticket_data"`
	TicketSignature []byte `json:"ticket_signature"`
}

func (d qrCode) ToJSONBytes() []byte {
	bs, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return bs
}

func hash(bs []byte) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	hw.Write(bs)
	hw.Sum(h[:0])
	return h
}

func initLogging(level log.Lvl, vmodule string) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(level)
	_ = glogger.Vmodule(vmodule)
	log.Root().SetHandler(glogger)
}

// createCtrlCContext returns context which's Done channel is closed when application should be terminated (on SIGINT, SIGTERM, SIGHUP, SIGQUIT signal)
func createCtrlCContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-sigChan
		log.Warn("got interrupt signal")
	}()

	return ctx
}

func readOwnerAddress(ctx context.Context, contractAddress common.Address, backend backend.Backend) (ownerAddress common.Address, err error) {
	didContract := contracts.InitPassportLogicContract(contractAddress, backend)
	ownerAddress, err = didContract.Owner(&bind.CallOpts{Context: ctx})
	return
}
