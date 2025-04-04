package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v2.8.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Terraformer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Terraformer " + version)
	},
}
