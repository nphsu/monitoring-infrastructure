package cmd

import (
	"shunp/api/server"

	"github.com/spf13/cobra"
)

var port int

var RootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := server.NewServer()
		if err != nil {
			return err
		}
		s.Run(port)
		cmd.AddCommand(DbCmd())
		return nil
	},
}

func init() {
	RootCmd.Flags().IntVarP(&port, "port", "p", 8080, "port number")
}
