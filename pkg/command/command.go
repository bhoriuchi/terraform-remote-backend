package command

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(initServerCommand())
	return cmd
}

func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("tfrb")

	return v
}

func setViperConfig(v *viper.Viper, config string) error {
	if config != "" {
		configPath, err := filepath.Abs(config)
		if err != nil {
			return fmt.Errorf("failed to resolve configuration file path: %s", err)
		}
		v.SetConfigFile(configPath)
	} else {
		v.AddConfigPath(".")
	}

	return nil
}
