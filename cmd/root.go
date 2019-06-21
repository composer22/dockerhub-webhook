package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/composer22/dockerhub-webhook/server"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Used globally for all commands
var (
	cfgFile     string
	hostname    string
	port        int
	profPort    int
	maxConn     int
	maxProcs    int
	debug       bool
	namespace   string
	validTokens string
	alivePath   string
	notifyPath  string
	statusPath  string
	targetHost  string
	targetPort  int
	targetPath  string
	targetToken string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   server.ApplicationName,
	Short: "Runs a small proxy server for relaying webhook calls from dockerhub",
	Long:  "Small server that obuscates a URL path and relays requests from dockerhub to jenkins",
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/."+server.ApplicationName+")")
	RootCmd.PersistentFlags().StringP("hostname", "O", server.DefaultHostname, "Host or IP for this server")
	RootCmd.PersistentFlags().IntP("port", "L", server.DefaultPort, "Listen port for this server")
	RootCmd.PersistentFlags().IntP("profile-port", "P", server.DefaultProfPort, "Profile port for this server")
	RootCmd.PersistentFlags().IntP("max-conn", "C", server.DefaultMaxConnections, "Maximum conn for this server")
	RootCmd.PersistentFlags().IntP("max-procs", "X", server.DefaultMaxProcs, "Maximum processors for this server")
	RootCmd.PersistentFlags().BoolP("debug", "D", false, "Debug")
	RootCmd.PersistentFlags().StringP("valid-tokens", "v", "", "List of valid tokens to access notifications (comma delim)")
	RootCmd.PersistentFlags().StringP("namespace", "n", server.DefaultNamespace, "Namespace for pricessing requests")
	RootCmd.PersistentFlags().StringP("alive-path", "a", server.DefaultAlivePath, "Path to handle health checks")
	RootCmd.PersistentFlags().StringP("notify-path", "y", server.DefaultNotifyPath, "Path to handle notification events")
	RootCmd.PersistentFlags().StringP("status-path", "s", server.DefaultStatusPath, "Path to get server status")
	RootCmd.PersistentFlags().StringP("target-host", "e", server.DefaultTargetHost, "Host to relay request")
	RootCmd.PersistentFlags().IntP("target-port", "o", server.DefaultTargetPort, "Port to access")
	RootCmd.PersistentFlags().StringP("target-path", "g", server.DefaultTargetPath, "Path to webhook")
	RootCmd.PersistentFlags().StringP("target-token", "t", "", "Authentication token")

	// Get values from config file.
	viper.BindPFlag("hostname", RootCmd.PersistentFlags().Lookup("hostname"))
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("profile-port", RootCmd.PersistentFlags().Lookup("profile-port"))
	viper.BindPFlag("max-conn", RootCmd.PersistentFlags().Lookup("max-conn"))
	viper.BindPFlag("max-procs", RootCmd.PersistentFlags().Lookup("max-procs"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("namespace", RootCmd.PersistentFlags().Lookup("namespace"))
	viper.BindPFlag("alive-path", RootCmd.PersistentFlags().Lookup("alive-path"))
	viper.BindPFlag("notify-path", RootCmd.PersistentFlags().Lookup("notify-path"))
	viper.BindPFlag("status-path", RootCmd.PersistentFlags().Lookup("status-path"))
	viper.BindPFlag("valid-tokens", RootCmd.PersistentFlags().Lookup("valid-tokens"))
	viper.BindPFlag("target-host", RootCmd.PersistentFlags().Lookup("target-host"))
	viper.BindPFlag("target-port", RootCmd.PersistentFlags().Lookup("target-port"))
	viper.BindPFlag("target-path", RootCmd.PersistentFlags().Lookup("target-path"))
	viper.BindPFlag("target-token", RootCmd.PersistentFlags().Lookup("target-token"))
	viper.SetDefault("hostname", server.DefaultHostname)
	viper.SetDefault("port", server.DefaultPort)
	viper.SetDefault("profile-port", server.DefaultProfPort)
	viper.SetDefault("max-conn", server.DefaultMaxConnections)
	viper.SetDefault("max-procs", server.DefaultMaxProcs)
	viper.SetDefault("debug", false)
	viper.SetDefault("valid-tokens", make([]string, 0))
	viper.SetDefault("namespace", server.DefaultNamespace)
	viper.SetDefault("alive-path", server.DefaultAlivePath)
	viper.SetDefault("notify-path", server.DefaultNotifyPath)
	viper.SetDefault("status-path", server.DefaultStatusPath)
	viper.SetDefault("target-host", server.DefaultTargetHost)
	viper.SetDefault("target-port", server.DefaultTargetPort)
	viper.SetDefault("target-path", server.DefaultTargetPath)
	viper.SetDefault("target-token", "")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home) // adding home directory as first search path.
		viper.AddConfigPath(".")
		viper.SetConfigName("." + server.ApplicationName) // name of config file (without extension).
	}

	viper.AutomaticEnv() // read in environment variables that match.

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Cannot find configuration file.\n\nERR: %s\n", err.Error())
		os.Exit(0)
	}
	hostname = viper.GetString("hostname")
	port = viper.GetInt("port")
	profPort = viper.GetInt("profile-port")
	maxConn = viper.GetInt("max-conn")
	maxProcs = viper.GetInt("max-procs")
	debug = viper.GetBool("debug")
	validTokens = viper.GetString("valid-tokens")
	namespace = viper.GetString("namespace")
	alivePath = viper.GetString("alive-path")
	notifyPath = viper.GetString("notify-path")
	statusPath = viper.GetString("status-path")
	targetHost = viper.GetString("target-host")
	targetPort = viper.GetInt("target-port")
	targetPath = viper.GetString("target-path")
	targetToken = viper.GetString("target-token")
	fmt.Printf("validTokens: %s\n", validTokens[1])
	fmt.Printf("len(validTokens): %d\n", len(validTokens))
	// We allow for some config settings via a filepath, so that it loads from
	// individual files if mounted.
	if v, err := ioutil.ReadFile(validTokens); err == nil {
		validTokens = strings.Trim(string(v), "\n")
	}

	if v, err := ioutil.ReadFile(targetToken); err == nil {
		targetToken = strings.Trim(string(v), "\n")
	}
}
