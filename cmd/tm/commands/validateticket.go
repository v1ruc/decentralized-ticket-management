package commands

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/monetha/go-verifiable-data/cmd/privatedata-exchange/commands/flag"
)

// ValidateTicketCommand handles `validate-ticket` command
type ValidateTicketCommand struct {
	EventDIDAddress flag.EthereumAddress `long:"eventaddr"        required:"true" description:"Ethereum address of event DID contract"`
	BackendURL      string               `long:"backendurl"       required:"true" description:"Ethereum backend URL"    default:"https://ropsten.infura.io"`
	Verbosity       log.Lvl              `long:"verbosity"                        description:"log verbosity (0-9)"     default:"2"`
	VModule         string               `long:"vmodule"                          description:"log verbosity pattern"`
}

// Execute implements flags.Commander interface
func (c *ValidateTicketCommand) Execute(args []string) error {
	panic("implement me")
}
