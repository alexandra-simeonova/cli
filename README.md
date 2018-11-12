# Kymactl

A command line tool to support Kyma developers

## Available Commands

Some of the following commands are still work in progress and are not yet usable

- `version`: Shows the Kyma CLI version. The version is set at compile time passing it to the go linker as a flag:
    ```bash
    go build -o kyma-cli -ldflags "-X github.wdf.sap.corp/SAP-CP-Extension-Factory/kyma-cli/cmd.Version=1.5.0"
    ```
- `status`: Tracks the status of a Kyma cluster (replaces the `is-installed.sh` script)
- `ctx`: Allows managing Kyma contexts (WIP).
- `help`: Displays usage for the given command (e.g. `kyma-cli help`, `kyma-cli help ctx`, etc...)