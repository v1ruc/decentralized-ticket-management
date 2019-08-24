package commands

import (
	"os"

	"github.com/v1ruc/decentralized-ticket-management/cmd"
)

// TicketManagementCommand is a root command used to parse command line arguments
type TicketManagementCommand struct {
	Version    func()            `short:"v" long:"version" description:"Print the version of tool and exit"`
	SignUp     SignUpCommand     `command:"signup"         description:"Sign up to participate in event (called by participant)"`
	SignUpList SignUpListCommand `command:"signup-list"    description:"List of registered participants (called by event organizer)"`
}

// TicketManagement is a prepared command to be used in command line arguments parsing
var TicketManagement TicketManagementCommand

func init() {
	TicketManagement.Version = func() {
		cmd.PrintVersion()
		os.Exit(0)
	}
}
