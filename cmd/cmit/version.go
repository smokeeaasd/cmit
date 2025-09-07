package cmit

import (
	"fmt"

	"github.com/smokeeaasd/cmit/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of cmit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmit version", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
