package cmit

import (
	"log"

	"github.com/smokeeaasd/cmit/internal/commit"
	"github.com/smokeeaasd/cmit/internal/form"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmit",
	Short: "cmit is a Git commit helper",
	Long:  "cmit is a CLI tool to generate conventional commit messages interactively.",
	Run: func(cmd *cobra.Command, args []string) {
		f := form.CreateForm()
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
