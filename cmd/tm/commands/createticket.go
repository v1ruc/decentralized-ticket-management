package commands

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/pkg/errors"
)

// CreateTicketCommand handles `create-ticket` command
type CreateTicketCommand struct {
	EventDIDAddress       flag.EthereumAddress         `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	ParticipantDIDAddress flag.EthereumAddress         `long:"participantaddr"  required:"true" description:"Ethereum address of participant DID contract"`
	PrivateKey            flag.ECDSAPrivateKeyFromFile `long:"privatekey"       required:"true" description:"private key filename of event owner"`
	BackendURL            string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	IPFSURL               string                       `long:"ipfsurl"                          description:"IPFS node URL"           default:"https://ipfs.infura.io:5001"`
	Verbosity             log.Lvl                      `long:"verbosity"                        description:"log verbosity (0-9)"     default:"2"`
	VModule               string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *CreateTicketCommand) Execute(args []string) error {
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

	eventOwnerKey := c.PrivateKey.AsECDSAPrivateKey()

	// creating ticket QR code
	ticketBytes := ticketData{
		EventDIDAddress:       c.EventDIDAddress.AsCommonAddress(),
		ParticipantDIDAddress: c.ParticipantDIDAddress.AsCommonAddress(),
	}.ToJSONBytes()
	ticketSignature, err := crypto.Sign(hash(ticketBytes).Bytes(), eventOwnerKey)
	if err != nil {
		return errors.Wrap(err, "failed to sign the ticket")
	}

	ticketQrCode := qrCode{
		ticketData:      ticketBytes,
		ticketSignature: ticketSignature,
	}.ToJSONBytes()

	// writing ticket to participants DID contract
	eventOwnerSession := e.NewSession(eventOwnerKey)

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "failed to create IPFS client")
	}

	wr := facts.NewPrivateDataWriter(eventOwnerSession, fs)
	res, err := wr.WritePrivateData(ctx, c.EventDIDAddress.AsCommonAddress(), ticketKey, ticketQrCode, rand.Reader)
	if err != nil {
		return errors.Wrap(err, "failed to write ticket data")
	}

	log.Info("Ticket info", "tx_hash", res.TransactionHash.Hex(), "ipfs_hash", res.DataIPFSHash)

	panic("implement me")
}
