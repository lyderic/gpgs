package cmd

import (
	"fmt"

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
	}
	if viper.GetBool("status") {
		agentStatus()
	}
}

func killAgent() {
	script.Exec("gpgconf --kill gpg-agent").Stdout()
}

func agentStatus() {
	fmt.Println("agentStatus called")
}

func init() {
	rootCmd.AddCommand(agentCmd)
	agentCmd.PersistentFlags().BoolP("status", "s", false, "show status of gpg agent")
	viper.BindPFlag("status", agentCmd.PersistentFlags().Lookup("status"))
	agentCmd.PersistentFlags().BoolP("kill", "k", false, "kill gpg agent")
	viper.BindPFlag("kill", agentCmd.PersistentFlags().Lookup("kill"))
}
