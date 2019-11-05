package cmd

import (
	"shunp/api/database"

	"github.com/spf13/cobra"
)

func DbCmd() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, arg []string) error {
			database.InitDB()
			return nil
		},
	}
	return cmd
}

func init() {
}
