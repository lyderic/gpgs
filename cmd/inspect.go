package cmd

import (
	"os"

	. "gpgs/internal"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect a gpg encrypted file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			Red("Please supply file as argument!\n")
			return
		}
		file := args[0]
		if _, err := os.Stat(file); os.IsNotExist(err) {
			Red("%s: file not found!\n", file)
			return
		}
		inspect(file)
	},
}

func inspect(file string) {
	command := `gpg --pinentry-mode=loopback --passphrase "" --list-packets ` + file
	pipe := script.Exec(command)
	pipe.SetError(nil)
	output, err := pipe.String()
	if err != nil {
		panic(err)
	}
	script.Echo(output).First(2).Stdout()
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
