package cli

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const DEFAULT_BIND_IP = "127.0.0.1"
const DEFAULT_PORT = "8080"

var bind = ""
var port = ""

func Setup() {
	viper.SetDefault("port", DEFAULT_PORT)
	viper.SetDefault("bind", DEFAULT_PORT)

	pflag.StringVarP(&port, "port", "p", DEFAULT_PORT, "Port to run the Autoconfig API on")
	viper.Set("port", port)
	pflag.StringVarP(&bind, "bind", "b", DEFAULT_BIND_IP, "IP to run the Autoconfig API on")
	viper.Set("bind", bind)
	// pflag.Lookup("port").NoOptDefVal = DEFAULT_PORT
	// pflag.Lookup("bind").NoOptDefVal = DEFAULT_BIND_IP

	// pflag.Parse()

	// viper.BindPFlags(pflag.CommandLine)
}
