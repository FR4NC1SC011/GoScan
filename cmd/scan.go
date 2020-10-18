package cmd

import (
	"github.com/spf13/cobra"
)

var scan = &cobra.Command{
	Use:   "scan",
	Short: "Scan the provided host",
	Run: func(cmd *cobra.Command, args []string) {
		//host := strings.Join(args[1], "")
		Scanner(args[0])
	},
}

func init() {
	RootCmd.AddCommand(scan)
}
