package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/youyo/awscredswrap"
)

var Version string

var rootCmd = &cobra.Command{
	Use:          "awscredswrap",
	Short:        "awscredswrap uses temporary credentials for the specified iam role to set a shell environment variable or execute a command.",
	Version:      Version,
	PreRun:       awscredswrap.PreRun,
	RunE:         awscredswrap.Run,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringP("role-arn", "r", "", "The arn of the role to assume.")
	rootCmd.Flags().StringP("role-session-name", "n", "", "An identifier for the assumed role session.")
	rootCmd.Flags().IntP("duration-seconds", "d", 3600, "The duration, in seconds, of the role session.")
	rootCmd.Flags().StringP("mfa-serial", "m", "", "The identification number of the MFA device that is associated with the user who is making the AssumeRole call.")

	viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {}
