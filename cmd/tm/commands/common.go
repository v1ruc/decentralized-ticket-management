package commands

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var (
	signUpKey [32]byte
)

func init() {
	copy(signUpKey[:], "signup")
}

type signUpData struct {
	ParticipantDIDAddress common.Address `json:"did_address"`
	FullName              string         `json:"full_name"`
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
