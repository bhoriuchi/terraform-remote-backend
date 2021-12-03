package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	serverAddressFlag = "address"
)

// server subcommand
func initServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Server commands",
	}

	cmd.AddCommand(initServerStartCommand())
	return cmd
}

// start the server
func initServerStartCommand() *cobra.Command {
	var configFile string
	v := newViper()

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := setViperConfig(v, configFile); err != nil {
				return err
			}

			fmt.Println(v.GetString("http.address"))

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&configFile, "config", "c", os.Getenv("TFRB_CONFIG"), "Path to the configuration file")

	cmd.Flags().StringP(serverAddressFlag, "a", ":9443", "Address for server to listen on")
	v.BindPFlag("http.address", cmd.Flags().Lookup(serverAddressFlag))

	return cmd
}
