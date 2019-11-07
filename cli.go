package awscredswrap

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run
func Run(cmd *cobra.Command, args []string) (err error) {
	roleArn := viper.GetString("role-arn")
	roleSessionName := viper.GetString("role-session-name")
	durationSeconds := time.Duration(viper.GetInt("duration-seconds")) * time.Second
	mfaSerial := viper.GetString("mfa-serial")

	awsCredsWrap := New()
	if err := awsCredsWrap.GetCredentials(roleArn, roleSessionName, mfaSerial, durationSeconds); err != nil {
		return errors.Wrap(err, "can not get credentials")
	}

	switch len(args) {
	case 0:
		envs := awsCredsWrap.ExportEnvironments()
		for _, v := range envs {
			fmt.Println(v)
		}

	case 1:
		if err = awsCredsWrap.ExecuteCommand(args[0], nil...); err != nil {
			return errors.Wrap(err, "failed to execute command")
		}

	default:
		if err = awsCredsWrap.ExecuteCommand(args[0], args[1:]...); err != nil {
			return errors.Wrap(err, "failed to execute command")
		}

	}

	return nil
}

//PreRun
func PreRun(cmd *cobra.Command, args []string) {
	roleSessionName := viper.GetString("role-session-name")
	if roleSessionName == "" {
		roleSessionName = "awscredswrap-session-" + time.Now().Format("20060102150405")
		viper.Set("role-session-name", roleSessionName)
	}
}
