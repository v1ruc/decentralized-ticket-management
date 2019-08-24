package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/contracts"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/pkg/errors"
)

// ReadTicketCommand handles `read-ticket` command
type ReadTicketCommand struct {
	EventDIDAddress       flag.EthereumAddress         `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	ParticipantDIDAddress flag.EthereumAddress         `long:"participantaddr"           required:"true" description:"Ethereum address of participant DID contract"`
	PrivateKey            flag.ECDSAPrivateKeyFromFile `long:"privatekey"       required:"true" description:"private key filename of participant"`
	BackendURL            string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	IPFSURL               string                       `long:"ipfsurl"                          description:"IPFS node URL"           default:"https://ipfs.infura.io:5001"`
	Verbosity             log.Lvl                      `long:"verbosity"                        description:"log verbosity (0-9)"     default:"2"`
	VModule               string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ReadTicketCommand) Execute(args []string) error {
	initLogging(c.Verbosity, c.VModule)
	ctx := createCtrlCContext()

	b, err := ethclient.DialContext(ctx, c.BackendURL)
	if err != nil {
		return errors.Wrap(err, "failed to create Ethereum backend")
	}

	e := eth.New(b, log.Info)
	if err := e.UpdateSuggestedGasPrice(ctx); err != nil {
		return errors.Wrap(err, "failed to update suggested gas price")
	}

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "failed to create IPFS client")
	}

	// reading event owner address
	didContract := contracts.InitPassportLogicContract(c.EventDIDAddress.AsCommonAddress(), b)
	eventOwnerAddress, err := didContract.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "failed to get event owner address")
	}

	// reading ticket QR code
	var factKey [32]byte
	copy(factKey[:], c.EventDIDAddress.AsCommonAddress().Bytes())

	rd := facts.NewPrivateDataReader(e, fs)
	ticketQrCodeBytes, err := rd.ReadPrivateData(ctx, c.PrivateKey.AsECDSAPrivateKey(),
		c.ParticipantDIDAddress.AsCommonAddress(), eventOwnerAddress, factKey)
	if err != nil {
		return errors.Wrap(err, "failed to read ticket data")
	}

	fmt.Println("Ticket QR:", string(ticketQrCodeBytes))

	return nil
}
