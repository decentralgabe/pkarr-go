package cli

import (
	"github.com/spf13/cobra"

	"pkarr-go/internal"
)

func init() {
	rootCmd.AddCommand(identityCmd)
}

var identityCmd = &cobra.Command{
	Use:   "id",
	Short: "Manage identities",
	RunE: func(cmd *cobra.Command, args []string) error {
		identities, err := internal.Read()
		if err != nil {
			return err
		}
		if len(identities) == 0 {
			println("No identities found.")
			return nil
		}
		for id, identity := range identities {
			println(id, identity.Records)
		}
		return nil
	},
}

var identityAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an identity",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
