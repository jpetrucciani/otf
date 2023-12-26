# Installation

## Requirements

- Linux - the server and agent components are tested on Linux only; the client CLI should work on all platforms.
- PostgreSQL - at least version 12.
- Terraform >= 1.2.0
- An SSL certificate.

## Download

There are three components that can be downloaded:

- `otfd` - the server daemon
- `otf` - the client CLI
- `otf-agent` - the agent daemon

Download them from [Github releases](https://github.com/jpetrucciani/otf/releases).

The server and agent components are also available as docker images:

- `jpetrucciani/otfd`
- `jpetrucciani/otf-agent`

## Install helm chart

You can install an `otfd` cluster on Kubernetes using the helm chart. See the [helm chart repository](https://github.com/jpetrucciani/otf-charts) for further information.

## Install from source

You'll need [Go](https://golang.org/doc/install).

Clone the repo, then build and install using the make task:

```bash
git clone https://github.com/jpetrucciani/otf
cd otf
make install
```

That'll install the binaries inside your go bin directory (defaults to `$HOME/go/bin`).
