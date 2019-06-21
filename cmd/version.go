package cmd

import (
	"fmt"

	"github.com/composer22/dockerhub-webhook/server"
	"github.com/spf13/cobra"
)

// versionCmd returns the version of the application
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the application",
	Long:  "Returns the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", server.Version())
	},
	Example: `dockerhub-webhook version`,
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.SetUsageTemplate(versionUsageTemplate())
}

// Override help template.
func versionUsageTemplate() string {
	return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
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
