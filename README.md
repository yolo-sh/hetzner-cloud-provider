<p align="center">
  <img src="https://user-images.githubusercontent.com/1233275/187983134-17b69812-f4a6-4359-969c-5384ee296ecf.png" alt="Hetzner's logo" width="180" height="180" />
</p>

<p align="center">
    <h1 align="center">Hetzner Cloud Provider</h1>
    <p align="center">This repository contains the source code that implements the Hetzner cloud provider for the <a href="https://github.com/yolo-sh/cli">Yolo CLI</a>.</p>
</p>

```bash
yolo hetzner --context production --region fsn1 init yolo-sh/api
```
<p align="center">
  <img width="759" src="https://user-images.githubusercontent.com/1233275/187968190-0ce41e41-4612-486e-bc3d-64f03fc25ac5.png" alt="Example of use of the Yolo CLI" />
</p>

## Table of contents
- [Usage](#usage)
    - [Authentication](#authentication)
        - [--context](#--context)
        - [--region](#--region)
    - [Permissions](#permissions)
    - [Authorized instance types](#authorized-instance-types)
- [Infrastructure lifecycle](#infrastructure-lifecycle)
    - [Init](#init)
    - [Edit](#edit)
    - [Open port](#open-port)
    - [Close port](#close-port)
    - [Remove](#remove)
    - [Uninstall](#uninstall)
- [Infrastructure costs](#infrastructure-costs)
- [License](#license)

## Usage

```console
To begin, create your first environment using the command:

  yolo hetzner init <repository>

Once initialized, you may want to connect it using the command: 

  yolo hetzner edit <repository>

If you don't plan to use this environment again, you could remove it using the command:
	
  yolo hetzner remove <repository>

<repository> may be relative to your personal GitHub account (eg: cli) or fully qualified (eg: my-organization/api).

Usage:
  yolo hetzner [command]

Examples:
  yolo hetzner init yolo-sh/api --instance-type cx11
  yolo hetzner edit yolo-sh/api
  yolo hetzner remove yolo-sh/api

Available Commands:
  close-port  Close a port in an environment
  edit        Connect your editor to an environment
  init        Initialize a new environment
  open-port   Open a port in an environment
  remove      Remove an environment
  uninstall   Uninstall Yolo from your Hetzner account

Flags:
      --context string   the configuration context to use to access your Hetzner account
  -h, --help             help for hetzner
      --region string    the region to use to access your Hetzner account

Use "yolo hetzner [command] --help" for more information about a command.
```

### Authentication

In order to access your Hetzner account, the Yolo CLI will first look for credentials in the following environment variables:

- `HCLOUD_TOKEN`

- `HCLOUD_REGION`

Otherwise, the configuration files created by the Hetzner CLI (via `hcloud context create`) will be used.

#### --context

If you have configured the Hetzner CLI with multiple configuration contexts, you could tell Yolo which one to use via the `--context` flag:

```shell
yolo hetzner --context production init yolo-sh/api
```

**By default, Yolo will use the active context.**

#### --region

If you want to overwrite the region resolved by the Yolo CLI, you could use the `--region` flag:

```shell
yolo hetzner --region fsn1 init yolo-sh/api
```

```shell
yolo hetzner --context production --region fsn1 init yolo-sh/api
```

### Permissions

Yolo requires your API token to have `Read & Write` permissions. See the next sections to learn more about the actions that will be done on your behalf.

### Authorized instance types

To be used with Yolo, the chosen instance must be **a cloud instance**.

#### Examples

```shell
cx11, cpx31, cx51...
```

## Infrastructure lifecycle

![infra](https://user-images.githubusercontent.com/1233275/187925670-e06790b5-0084-4d91-a18e-160c771b4f4a.png)

The schema above describe all the components that may be created in your Hetzner account. The next sections will describe their lifetime according to your use of the Yolo CLI.

### Init

#### The first time Yolo is used in a region

```bash
yolo hetzner init yolo-sh/api --instance-type cx11
```

When running the `init` command for the first time in a region, the following components will be created:

- A `network` named `yolo-network` with an IPv4 CIDR block equals to `10.0.0.0/16` to isolate your infrastructure.

- A `subnet`with an IPv4 CIDR block equals to `10.0.0.0/24` that will contain your environments' instances.

#### On each init

When running the `init` command, the following components will be created:

- A `firewall` named `yolo-${ENV_NAME}-firewall` to let your environment accepts `SSH` connections on port `2200`.

- An `SSH key` named `yolo-${ENV_NAME}-ssh-key` to let you access your environment via `SSH`.

- A `server` named `yolo-${ENV_NAME}-server` with a type equals to the one passed via the `--instance-type` flag or `cx11` by default.

### Edit

```bash
yolo hetzner edit yolo-sh/api
```

When running the `edit` command, nothing will be done to your infrastructure.

### Open port

```bash
yolo hetzner open-port yolo-sh/api 8080
```

When running the `open-port` command, an `inbound` rule will be added to the `firewall` of the environment. 

This rule will allow all `TCP` trafic from `any IP address` to the specified port.

### Close port

```bash
yolo hetzner close-port yolo-sh/api 8080
```

When running the `close-port` command, the `inbound` rule added by the `open-port` command will be removed.

### Remove

```bash
yolo hetzner remove yolo-sh/api
```

When running the `remove` command, all the components associated with the environment will be removed.

In other words:

- The `server`.

- The `SSH key`.

- The `firewall`.

### Uninstall

```bash
yolo hetzner uninstall
```

When running the `uninstall` command, all the shared components will be removed. 

In other words:

- The `network`.

- The `subnet`.

## Infrastructure costs

The cost of running an environment on Hetzner is equal to the cost of the `server` used.

## License

Yolo is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
