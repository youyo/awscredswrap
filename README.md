# awscredswrap

[![Go Report Card](https://goreportcard.com/badge/github.com/youyo/awscredswrap)](https://goreportcard.com/report/github.com/youyo/awscredswrap)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](./LICENSE)

AWS assume role credential wrapper.

## Description

awscredswrap uses temporary credentials for the specified iam role to set a shell environment variable or execute a command.

## Install

- Brew

```
$ brew install youyo/tap/awscredswrap
```

Other platforms are download from [github release page](https://github.com/youyo/awscredswrap/releases).

## Usage

```bash
$ awscredswrap --help
awscredswrap uses temporary credentials for the specified iam role to set a shell environment variable or execute a command.

Usage:
  awscredswrap [flags]

Flags:
  -d, --duration-seconds int       The duration, in seconds, of the role session. (default 3600)
  -h, --help                       help for awscredswrap
  -m, --mfa-serial string          The identification number of the MFA device that is associated with the user who is making the AssumeRole call.
  -r, --role-arn string            The arn of the role to assume.
  -n, --role-session-name string   An identifier for the assumed role session.
      --version                    version for awscredswrap
```

### As command wrapper

```console
$ awscredswrap --role-arn arn:aws:iam::00000000:role/foo -- some_command [arg1 arg2...]
```

### As env exporter

When awscredswrap is executed with no arguments, awscredswrap outputs shell script to export AWS credentials environment variables.

```console
$ awscredswrap --role-arn arn:aws:iam::00000000:role/foo
export AWS_ACCESS_KEY_ID='XXXXXXXXXXXXXXXX'
export AWS_SECRET_ACCESS_KEY='zWarBXUtMKJYnC8y4dNAf9e5HQqFTp....'
export AWS_SESSION_TOKEN='Wj3YGuSMwn8aJx4AN6TFsbtB5URKHEpVgdDkPvy7....'
export AWS_DEFAULT_REGION='us-east-1'
```

You can set the credentials in current shell by `eval`.

```console
$ eval $(awscredswrap --role-arn arn:aws:iam::00000000:role/foo)
```

Temporary credentials has expiration time (about 1 hour).

## License

[MIT](LICENSE)

## Author

[youyo](https://github.com/youyo)
