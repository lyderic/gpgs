package cmd

import (
	"fmt"
	. "gpgs/internal"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list keys in gpg keyring",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	if viper.GetBool("full") {
		listKeys("secret")
		listKeys("public")
	} else {
		listUids()
	}
}

func listUids() {
	for _, uid := range GetUids() {
		fmt.Println(">", uid)
	}
}

func listKeys(keytype string) {
	Blue("[%s keys] ", keytype)
	command := fmt.Sprintf("gpg --list-%s-keys", keytype)
	script.Exec(command).Reject(viper.GetString("keyring")).Stdout()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("full", "f", false, "provide detailed information")
	viper.BindPFlag("full", listCmd.Flags().Lookup("full"))
}
