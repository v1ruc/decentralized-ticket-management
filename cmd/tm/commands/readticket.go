package commands

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
)

// ReadTicketCommand handles `read-ticket` command
type ReadTicketCommand struct {
	EventDIDAddress       flag.EthereumAddress         `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	ParticipantDIDAddress flag.EthereumAddress         `long:"myaddr"           required:"true" description:"Ethereum address of participant DID contract"`
	PrivateKey            flag.ECDSAPrivateKeyFromFile `long:"privatekey"       required:"true" description:"private key filename of participant"`
	BackendURL            string                       `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	IPFSURL               string                       `long:"ipfsurl"                          description:"IPFS node URL"           default:"https://ipfs.infura.io:5001"`
	Verbosity             log.Lvl                      `long:"verbosity"                        description:"log verbosity (0-9)"     default:"2"`
	VModule               string                       `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ReadTicketCommand) Execute(args []string) error {
	panic("implement me")
}
