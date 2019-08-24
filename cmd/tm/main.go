package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/v1ruc/decentralized-ticket-management/cmd/tm/commands"
)

func main() {
	p := flags.NewParser(&commands.TicketManagement, flags.Default)

	_, err := p.Parse()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			}
		}

		os.Exit(1)
	}
}
