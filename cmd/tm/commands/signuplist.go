package commands

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/monetha/go-verifiable-data/types/change"
	"github.com/monetha/go-verifiable-data/types/data"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// SignUpListCommand handles `signup-list` command
type SignUpListCommand struct {
	EventDIDAddress flag.EthereumAddress         `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	PrivateKey      flag.ECDSAPrivateKeyFromFile `long:"privatekey"       required:"true" description:"private key filename of event owner"`
	BackendURL      string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	IPFSURL         string                       `long:"ipfsurl"                          description:"IPFS node URL"           default:"https://ipfs.infura.io:5001"`
	Verbosity       log.Lvl                      `long:"verbosity"                        description:"log verbosity (0-9)"     default:"3"`
	VModule         string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *SignUpListCommand) Execute(args []string) error {
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

	historian := facts.NewHistorian(e)
	eventPassportAddress := c.EventDIDAddress.AsCommonAddress()
	// filtering signup changes
	filterOpts := &facts.ChangesFilterOpts{
		Context:    ctx,
		ChangeType: []change.Type{change.Updated},
		DataType:   []data.Type{data.PrivateData},
		Key:        [][32]byte{signUpKey},
	}
	it, err := historian.FilterChanges(filterOpts, eventPassportAddress)
	if err != nil {
		return errors.Wrap(err, "failed to filter sing-up changes")
	}

	// creating list of signup providers
	signUpChanges := make(map[common.Address]types.Log)
	for it.Next() {
		if err := it.Error(); err != nil {
			return errors.Wrap(err, "failed to iterate sing-up changes")
		}

		ch := it.Change
		signUpChanges[ch.FactProvider] = ch.Raw
	}

	// reading signup private data
	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "failed to create IPFS client")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Participant DID", "Full Name"})

	rd := facts.NewPrivateDataReader(e, fs)
	eventOwnerKey := c.PrivateKey.AsECDSAPrivateKey()
	for _, txRaw := range signUpChanges {
		decryptedSignUpData, err := rd.ReadHistoryPrivateData(ctx, eventOwnerKey, eventPassportAddress, txRaw.TxHash)
		if err != nil {
			log.Warn("failed to read private data", "tx_hash", txRaw.TxHash, "err", err)
			continue
		}

		d := &signUpData{}
		if err := d.FromJSONBytes(decryptedSignUpData); err != nil {
			log.Warn("failed to parse signup JSON", "bytes", decryptedSignUpData, "err", err)
			continue
		}

		table.Append([]string{d.ParticipantDIDAddress.Hex(), d.FullName})
	}

	table.Render()

	return nil
}
