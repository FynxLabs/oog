package cmd

import (
	"github.com/fynxlabs/oog/lib/oog"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run OOG",
	Long:  `Run OOG Bot w/ any passed flags or configs`,
	Run: func(cmd *cobra.Command, args []string) {
		oog.run()
	},
}
