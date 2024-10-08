# CLI

`otf` is the CLI for OTF.

Download a [release](https://github.com/jpetrucciani/otf/releases). Ensure you select the client component, `otf`. The release is a zip file. Extract the `otf` binary to a directory in your system PATH.

Run `otf` with no arguments to receive usage instructions:

```bash
Usage:
  otf [command]

Available Commands:
  agents        Agent management
  help          Help about any command
  organizations Organization management
  runs          Runs management
  state         State version management
  teams         Team management
  users         User account management
  workspaces    Workspace management

Flags:
      --address string   Address of OTF server (default "localhost:8080")
  -h, --help             help for otf
      --token string     API authentication token

Use "otf [command] --help" for more information about a command.
```

Credentials are sourced from the same file the terraform CLI uses (`~/.terraform.d/credentials.tfrc.json`). To populate credentials, run:

```bash
terraform login <otfd_hostname>
```
