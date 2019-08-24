package commands

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
	"github.com/monetha/go-verifiable-data/eth"
	"github.com/monetha/go-verifiable-data/eth/backend/ethclient"
	"github.com/monetha/go-verifiable-data/facts"
	"github.com/monetha/go-verifiable-data/ipfs"
	"github.com/pkg/errors"
)

// SignUpCommand handles `signup` command
type SignUpCommand struct {
	EventDIDAddress       flag.EthereumAddress         `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	ParticipantDIDAddress flag.EthereumAddress         `long:"participantaddr"           required:"true" description:"Ethereum address of participant DID contract"`
	ParticipantFullName   string                       `long:"participantname"           required:"true" description:"Full name of participant"`
	PrivateKey            flag.ECDSAPrivateKeyFromFile `long:"privatekey"       required:"true" description:"private key filename of participant"`
	BackendURL            string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	IPFSURL               string                       `long:"ipfsurl"                          description:"IPFS node URL"           default:"https://ipfs.infura.io:5001"`
	Verbosity             log.Lvl                      `long:"verbosity"                        description:"log verbosity (0-9)"     default:"2"`
	VModule               string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *SignUpCommand) Execute(args []string) error {
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

	participantSession := e.NewSession(c.PrivateKey.AsECDSAPrivateKey())

	fs, err := ipfs.New(c.IPFSURL)
	if err != nil {
		return errors.Wrap(err, "failed to create IPFS client")
	}

	wr := facts.NewPrivateDataWriter(participantSession, fs)

	signUpPrivateData := signUpData{
		ParticipantDIDAddress: c.ParticipantDIDAddress.AsCommonAddress(),
		FullName:              c.ParticipantFullName,
	}.ToJSONBytes()
	res, err := wr.WritePrivateData(ctx, c.EventDIDAddress.AsCommonAddress(), signUpKey, signUpPrivateData, rand.Reader)
	if err != nil {
		return errors.Wrap(err, "failed to write private signup data")
	}

	log.Info("Signup info", "tx_hash", res.TransactionHash.Hex(), "ipfs_hash", res.DataIPFSHash)

	return nil
}
