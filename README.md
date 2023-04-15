# s3heck

- [s3heck](#s3heck)
  - [Usage](#usage)
    - [Subcommands](#subcommands)
      - [Example](#example)
  - [Installation](#installation)
  - [Contributing](#contributing)

s3heck is a command line interface that can perform checks on the certificate chain.

## Usage

```bash
$ s3heck -h
s3heck is a command line interface that can perform checks on the certificate chain.

Usage:
  s3heck [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  x509        display necessary information about x509 certificates

Flags:
  -h, --help   help for s3heck

Use "s3heck [command] --help" for more information about a command.
```

### Subcommands

Use the `x509` to have outputs the desired information about the x509 certificate

#### Example

```bash
$ s3heck x509 -p ./data/cert.pem --issuer --subject --validity --dns
Certificate [0]
  Issuer: R3
  Subject: kanywst.top
  Validity:
    NotBefore: 2023-03-31 13:32:00 +0000 UTC
    NotAfter: 2023-06-29 13:31:59 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name: kanywst.top, tbow.kanywst.top
Certificate [1]
  Issuer: ISRG Root X1
  Subject: R3
  Validity:
    NotBefore: 2020-09-04 00:00:00 +0000 UTC
    NotAfter: 2025-09-15 16:00:00 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name:
Certificate [2]
  Issuer: DST Root CA X3
  Subject: ISRG Root X1
  Validity:
    NotBefore: 2021-01-20 19:14:03 +0000 UTC
    NotAfter: 2024-09-30 18:14:03 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name:
```

## Installation

```bash
git clone https://github.com/kanywst/s3heck.git
cd s3heck
go install
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am "Add some feature"`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
