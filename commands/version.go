package commands

import (
	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/spf13/cobra"
)

var VerSion = &cobra.Command{
	Use:   "version",
	Short: "Display Eris's Platform Version.",
	Long:  `Display the current installed version of Eris.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Eris CLI Version: " + VERSION)
	},
}
