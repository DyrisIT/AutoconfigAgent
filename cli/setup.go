package cli

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const DEFAULT_BIND_IP = "127.0.0.1"
const DEFAULT_PORT = 8080

func Setup() {
	viper.SetDefault("port", DEFAULT_PORT)
	viper.SetDefault("bind", DEFAULT_PORT)

	pflag.StringP("port", "p", fmt.Sprintf("%d", DEFAULT_PORT), "Port to run the Autoconfig API on")
	pflag.StringP("bind", "b", DEFAULT_BIND_IP, "IP to run the Autoconfig API on")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}
