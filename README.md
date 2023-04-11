# s3heck

- [s3heck](#s3heck)
  - [Install](#install)
  - [Usage](#usage)
    - [Help](#help)
    - [Example](#example)

s3heck is a command line interface that can perform checks on the certificate chain.

## Install

```bash
git clone https://github.com/kanywst/s3heck.git
cd s3heck
go install
```

## Usage

### Help

```bash
$ s3heck x509 -h
display necessary information about x509 certificates

Usage:
  s3heck x509 [flags]

Flags:
  -d, --dns          display subject alternative name
  -h, --help         help for x509
  -i, --issuer       display issuer
  -p, --pem string   specify certificate input pem
  -s, --subject      display subject
  -v, --validity     display validity
```

### Example

```bash
$ s3heck x509 -p ./data/cert.pem --issuer --subject --validity --dns
Certificate [1]
  Issuer: R3
  Subject: kanywst.top
  Validity:
    NotBefore: 2023-03-31 13:32:00 +0000 UTC
    NotAfter: 2023-06-29 13:31:59 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name: kanywst.top, tbow.kanywst.top
Certificate [2]
  Issuer: ISRG Root X1
  Subject: R3
  Validity:
    NotBefore: 2020-09-04 00:00:00 +0000 UTC
    NotAfter: 2025-09-15 16:00:00 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name:
Certificate [3]
  Issuer: DST Root CA X3
  Subject: ISRG Root X1
  Validity:
    NotBefore: 2021-01-20 19:14:03 +0000 UTC
    NotAfter: 2024-09-30 18:14:03 +0000 UTC
  X509v3 extensions
    X509v3 Subject Alternative Name:
```
