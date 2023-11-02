package cli

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const DEFAULT_PORT = 8080

func Setup() {
	port()
}

func port() {
	viper.SetDefault("port", DEFAULT_PORT)

	pflag.StringP("port", "p", fmt.Sprintf("%d", DEFAULT_PORT), "Port to run the Autoconfig API on")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}
