package cmd

import (
	"strings"

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
	keyinfo := script.Exec(`gpg-connect-agent "keyinfo --list" /bye`)
	var lines []string
	keyinfo.EachLine(func(line string, output *strings.Builder) {
		lines = append(lines, line)
	})
	for _, line := range lines {
		if !strings.HasPrefix(line, "S KEYINFO") {
			continue
		}
		bits := strings.Fields(line)
		grip := bits[2]
		status := bits[6]
		Magenta("%s %s\n", grip, status)
	}
	for _, uid := range []string{"passs", "flipouk"} {
		Yellow("--- %s ---\n", uid)
		script.Exec("gpg --list-secret-keys --with-keygrip --with-colons " + uid).Stdout()
	}
}

func init() {
	rootCmd.AddCommand(agentCmd)
	agentCmd.PersistentFlags().BoolP("kill", "k", false, "kill gpg agent")
	viper.BindPFlag("kill", agentCmd.PersistentFlags().Lookup("kill"))
}
