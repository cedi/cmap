package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version represents the Version of the kkpctl binary, should be set via ldflags -X
	Version string

	// Date represents the Date of when the kkpctl binary was build, should be set via ldflags -X
	Date string

	// Commit represents the Commit-hash from which kkpctl binary was build, should be set via ldflags -X
	Commit string

	// BuiltBy represents who build the binary, should be set via ldflags -X
	BuiltBy string

	outputType string
	sortBy     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmap",
	Short: "map my network in a more efficient way",
	Long:  `This is a CLI for interacting with the nmap API, as I'm just too lazy to learn nmap properly and because I hate the way nmap returns results on the CLI...`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit, date, builtBy string) {

	// asign build flags for version info
	Version = version
	Date = date
	Commit = commit
	BuiltBy = builtBy

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "text", "The output type to use")
	rootCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"text", "json", "yaml"}, cobra.ShellCompDirectiveDefault
	})
}
