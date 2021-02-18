package cmd

import (
	. "gpgs/internal"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "control gpg-agent",
	Run: func(cmd *cobra.Command, args []string) {
		agent()
	},
}

func agent() {
	if viper.GetBool("kill") {
		killAgent()
	} else {
		agentStatus()
	}
}

func killAgent() {
	script.Exec("gpgconf --kill gpg-agent").Stdout()
}

func agentStatus() {
	for _, keyinfo := range GetKeyinfos() {
		if keyinfo.Status == "1" {
			for _, uid := range GetUids() {
				for _, grip := range GetGrips(uid) {
					if keyinfo.Grip == grip {
						Green("Key Loaded: %s\n", uid)
					}
				}
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(agentCmd)
	agentCmd.PersistentFlags().BoolP("kill", "k", false, "kill gpg agent")
	viper.BindPFlag("kill", agentCmd.PersistentFlags().Lookup("kill"))
}
