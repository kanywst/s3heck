# s3heck (The Unified PKI Orchestrator)

![GitHub Actions](https://github.com/kanywst/s3heck/actions/workflows/ci.yml/badge.svg)
![Go Version](https://img.shields.io/github/go-mod/go-version/kanywst/s3heck)
![License](https://img.shields.io/github/license/kanywst/s3heck)

![s3heck Demo](demo.gif)

**s3heck** is the next-generation CLI tool designed for the 2026 DevSecOps landscape. It transforms the complex, fragmented world of Public Key Infrastructure (PKI) and mTLS into a simple, unified workflow.

It replaces obscure `openssl` commands and manual Kubernetes log hunting with intelligent, actionable diagnostics.

## 🧠 Why s3heck?

### 🚫 The Problem
- **OpenSSL is hard**: Remembering `openssl s_client -connect -servername ...` flags is painful.
- **Diagnostics are manual**: Checking expiration, chain completeness, and algorithm strength requires expert knowledge.
- **K8s PKI is fragmented**: Debugging `autocert` or `step-issuer` involves checking multiple Pods, Events, and Webhooks manually.

### ✅ The s3heck Solution
- **One Command Diagnosis**: Smartly analyze local certs, remote endpoints, and K8s clusters.
- **Proactive Warnings**: Warns about expiring certs, weak algorithms (SHA1/MD5), and missing intermediate CAs.
- **Production Scaffolding**: Generates best-practice configurations for the Smallstep ecosystem.

## 🚀 Features

- **🔍 Smart Diagnostics (`check`)**: 
  - **Remote**: Auto-detects TLS version, SNI issues, chain validity, and expiration.
  - **Local**: Validates PEM format, parses details, and checks security standards.
  - **Kubernetes**: Scans `autocert` and `step-issuer` status, checking Pod logs, K8s Events, and Webhook configurations for common errors.
- **🏗️ Scaffolding (`scaffold`)**: 
  - Generate production-ready K8s manifests for `autocert` and `step-issuer`.
  - Create working Go mTLS client/server examples instantly.
- **📚 Integrated Knowledge (`learn`)**: 
  - Access documentation for the entire ecosystem (`autocert`, `step-issuer`, `kms`) without context switching.

## 📦 Installation

```bash
go install github.com/kanywst/s3heck@latest
```

## 🛠 Usage

### 1. Intelligent Diagnostics

**Check a remote endpoint:**
```bash
# Detects SNI, TLS version, Chain issues, and Expiry
s3heck check --remote google.com
```

**Diagnose Kubernetes Cluster:**
```bash
# Scans Pods, Events, and Webhooks for PKI errors
s3heck check --k8s
```

**Inspect a local certificate:**
```bash
s3heck check --file ./server.crt
```

### 2. Instant Scaffolding

**Deploy Autocert to K8s:**
```bash
s3heck scaffold autocert > deployment.yaml
```

**Configure Step Issuer:**
```bash
s3heck scaffold issuer > issuer.yaml
```

**Generate mTLS Code:**
```bash
s3heck scaffold mtls
# Creates mtls-server.go and mtls-client.go
```

### 3. In-Terminal Learning

```bash
s3heck learn autocert
s3heck learn kms
```

## 📂 Included Ecosystem

This tool unifies the best tools from Smallstep:

- **[Autocert](https://github.com/smallstep/autocert)**: Automatic K8s certificate injection.
- **[Step Issuer](https://github.com/smallstep/step-issuer)**: Cert-manager issuer for private CAs.
- **[Smallstep CLI](https://github.com/smallstep/cli)**: The swiss-army knife of crypto.
- **[Step KMS Plugin](https://github.com/smallstep/step-kms-plugin)**: Hardware key protection.

## 🤝 Contributing

We welcome contributions! Please see [`CONTRIBUTING.md`](CONTRIBUTING.md) for details on how to get started.

1. Fork the repo
2. Create a feature branch
3. Commit your changes
4. Submit a PR

---
*Built with ❤️ for the 2026 Developer.*
