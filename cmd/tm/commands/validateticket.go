package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
	"github.com/pkg/errors"
)

// ValidateTicketCommand handles `validate-ticket` command
type ValidateTicketCommand struct {
	EventDIDAddress flag.EthereumAddress `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	TicketQRCode    string               `long:"ticket"           required:"true" description:"Ticket QR code to validate"`
	BackendURL      string               `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	Verbosity       log.Lvl              `long:"verbosity"                        description:"log verbosity (0-9)"     default:"2"`
	VModule         string               `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ValidateTicketCommand) Execute(args []string) error {
	initLogging(c.Verbosity, c.VModule)
	ctx := createCtrlCContext()

	// decoding ticket QR code
	ticketQRBytes, err := hexutil.Decode(c.TicketQRCode)
	if err != nil {
		return errors.Wrap(err, "failed to decode ticket QR code")
	}
	qr := &qrCode{}
	if err := qr.FromJSONBytes(ticketQRBytes); err != nil {
		return errors.Wrap(err, "failed to unmarshal QR code JSON")
	}

	// validating event DID address
	td := &ticketData{}
	if err := td.FromJSONBytes(qr.TicketData); err != nil {
		return errors.Wrap(err, "failed to unmarshal ticked data JSON")
	}
	if td.EventDIDAddress != c.EventDIDAddress.AsCommonAddress() {
		return fmt.Errorf("ticket event DID (%v) does not match provided event ID (%v)", td.EventDIDAddress.Hex(), c.EventDIDAddress.AsCommonAddress().Hex())
	}

	// recovering event owner address from signature
	recoveredEventOwnerPubKey, err := crypto.SigToPub(hash(qr.TicketData).Bytes(), qr.TicketSignature)
	if err != nil {
		return errors.Wrap(err, "failed to recover event owner public key")
	}
	recoveredEventOwnerAddress := crypto.PubkeyToAddress(*recoveredEventOwnerPubKey)

	b, err := ethclient.DialContext(ctx, c.BackendURL)
	if err != nil {
		return errors.Wrap(err, "failed to create Ethereum backend")
	}

	// reading event owner address
	eventOwnerAddress, err := readOwnerAddress(ctx, c.EventDIDAddress.AsCommonAddress(), b)
	if err != nil {
		return errors.Wrap(err, "failed to get event owner address")
	}

	// validating signer address
	if recoveredEventOwnerAddress != eventOwnerAddress {
		return fmt.Errorf("the ticket is signed not by the owner of the event, expected signer address (%v), got (%v)",
			eventOwnerAddress.Hex(), recoveredEventOwnerAddress.Hex())
	}

	fmt.Printf("Valid ticket for participant DID: %v\n", td.ParticipantDIDAddress.Hex())
	return nil
}
