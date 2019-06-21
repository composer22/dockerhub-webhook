package cmd

import (
	"strings"

	"github.com/composer22/dockerhub-webhook/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command for starting the server.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the dockerhub webhook server",
	Long:  "Starts the dockerhub webhook server and waits to forward a request",
	Run: func(cmd *cobra.Command, args []string) {
		opt := server.OptionsNew(func(o *server.Options) {
			o.Hostname = hostname
			o.Port = port
			o.ProfPort = profPort
			o.MaxConn = maxConn
			o.MaxProcs = maxProcs
			o.Debug = debug
			o.Namespace = namespace
			o.ValidTokens = strings.Split(validTokens, ",")
			o.AlivePath = alivePath
			o.NotifyPath = notifyPath
			o.StatusPath = statusPath
			o.TargetHost = targetHost
			o.TargetPort = targetPort
			o.TargetPath = targetPath
			o.TargetToken = targetToken
		})
		server.New(opt).Start()
	},
	Example: `dockerhub-webhook start --target-host 127.0.0.1 --target-port 8080 --target-path "/generic-webhook-trigger/invoke/" --target-token A12378
`,
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.SetUsageTemplate(getUsageTemplate())
}

// Override help template.
func getUsageTemplate() string {
	return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags] KEY"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
