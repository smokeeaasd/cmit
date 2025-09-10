package cmit

import (
	"fmt"
	"log"

	"github.com/smokeeaasd/cmit/internal/commit"
	"github.com/smokeeaasd/cmit/internal/form"
	"github.com/smokeeaasd/cmit/internal/version"
	"github.com/spf13/cobra"
)

var (
	showVersion bool
	detailed    bool
)

var rootCmd = &cobra.Command{
	Use:   "cmit",
	Short: "cmit is a Git commit helper",
	Long:  "cmit is a CLI tool to generate conventional commit messages interactively.",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println("cmit version", version.Version)
			return
		}

		f := form.CreateForm(detailed)

		if err := f.Run(); err != nil {
			log.Fatal(err)
		}

		commit.ExecuteCommit(ExtraArgs)
	},
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version and exit")
	rootCmd.PersistentFlags().BoolVar(&detailed, "detailed", false, "Enable detailed commit mode")
}
