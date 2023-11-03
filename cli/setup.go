package cli

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const DEFAULT_BIND_IP = "127.0.0.1"
const DEFAULT_PORT = "8080"

func Setup() {
	pflag.StringP("bind", "b", DEFAULT_BIND_IP, "IP to run the Autoconfig API on")
	pflag.StringP("port", "p", DEFAULT_PORT, "Port to run the Autoconfig API on")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.SetDefault("bind", DEFAULT_BIND_IP)
	viper.SetDefault("port", DEFAULT_PORT)

}
