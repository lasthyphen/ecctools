package subnetcmd

import (
	"fmt"

	"github.com/lasthyphen/ecctools/pkg/application"
	"github.com/spf13/cobra"
)

// var app *application.GoGoTools

func NewCmd(injectedApp *application.GoGoTools) *cobra.Command {
	// app = injectedApp

	cmd := &cobra.Command{
		Use:    "subnet",
		Hidden: true,
		Short:  "Coming Soon!",
		Long:   ``,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	// cmd.AddCommand(newSubnetidCmd())
	return cmd
}
