package awscredswrap

import (
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AwsCredsWrap struct {
	Session     *session.Session
	Credentials credentials.Value
	Region      string
}

// New
func New() *AwsCredsWrap {
	sess := newSession()
	awsCredsWrap := &AwsCredsWrap{
		Session: sess,
		Region:  *sess.Config.Region,
	}

	return awsCredsWrap
}

func (a *AwsCredsWrap) GetCredentials(roleArn, roleSessionName, mfaSerial string, durationSeconds time.Duration) (err error) {
	creds := stscreds.NewCredentials(a.Session, roleArn, assumeRoleProvider(roleSessionName, mfaSerial, durationSeconds))

	a.Credentials, err = creds.Get()
	if err != nil {
		return err
	}

	return nil
}

func (a *AwsCredsWrap) ExportEnvironments() []string {
	s := []string{
		"export AWS_ACCESS_KEY_ID='" + a.Credentials.AccessKeyID + "'",
		"export AWS_SECRET_ACCESS_KEY='" + a.Credentials.SecretAccessKey + "'",
		"export AWS_SESSION_TOKEN='" + a.Credentials.SessionToken + "'",
		"export AWS_DEFAULT_REGION='" + a.Region + "'",
	}

	return s
}

func (a *AwsCredsWrap) ExecuteCommand(com string, args ...string) (err error) {
	a.setEnvironments()

	err = execCommand(com, args...)

	return err
}

func newSession() (sess *session.Session) {
	sess = session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState:       session.SharedConfigEnable,
				AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
			},
		),
	)

	return sess
}

func assumeRoleProvider(roleSessionName, mfaSerial string, durationSeconds time.Duration) (f func(p *stscreds.AssumeRoleProvider)) {
	f = func() (f func(p *stscreds.AssumeRoleProvider)) {
		if mfaSerial != "" {
			return func(p *stscreds.AssumeRoleProvider) {
				p.RoleSessionName = roleSessionName
				p.Duration = durationSeconds
				p.SerialNumber = aws.String(mfaSerial)
				p.TokenProvider = stscreds.StdinTokenProvider
			}
		} else {
			return func(p *stscreds.AssumeRoleProvider) {
				p.RoleSessionName = roleSessionName
				p.Duration = durationSeconds
			}
		}
	}()

	return f
}

func (a *AwsCredsWrap) setEnvironments() {
	os.Setenv("AWS_ACCESS_KEY_ID", a.Credentials.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", a.Credentials.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", a.Credentials.SessionToken)
	os.Setenv("AWS_DEFAULT_REGION", a.Region)
}

func execCommand(com string, args ...string) (err error) {
	var command *exec.Cmd

	if args != nil {
		command = exec.Command(com, args...)
	} else {
		command = exec.Command(com)
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if err = command.Start(); err != nil {
		return err
	}

	command.Wait()

	return nil
}
