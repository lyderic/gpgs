package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gpgs",
	Short: "wrapper to various, hard to remember gpg commands",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gpgs.yaml)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	/*
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

			// Search config in home directory with name ".gpgs" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigName(".gpgs")
		}
	*/

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetDefault("keyring", filepath.Join(".gnupg", "pubring.kbx"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
