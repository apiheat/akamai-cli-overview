# Akamai CLI for Property Overview
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

### Local Install, if you choose not to use the akamai package manager
If you want to compile it from source, you will need Go 1.9 or later, and the [Glide](https://glide.sh) package manager installed:
1. Fetch the package:
   `go get https://github.com/partamonov/akamai-cli-overview`
1. Change to the package directory:
   `cd $GOPATH/src/github.com/partamonov/akamai-cli-overview`
1. Install dependencies using Glide:
   `glide install`
1. Compile the binary:
   `go build -ldflags="-s -w -X main.version=X.X.X" -o akamai-overview`

### Credentials
In order to use this configuration, you need to:
* Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Expects `default` section in .edgerc, can be changed via --section parameter

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

## Overview
This CLI intended to be used as overview tool. It support list commands for PAPI API.

## Main Command Usage
```shell
NAME:
   akamai overview - A CLI to list contracts, groups, properties

USAGE:
   akamai overview [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     contracts, ctr     Lists the set of contracts
     cpcodes, cpc       Lists CP codes available within your contract/group pairing
     custom-overrides   Lists the set of custom XML metadata overrides configured for you by Akamai representatives
     edge-hostnames, e  Lists all edge hostnames available under a contract
     groups, grp        Lists the set of groups for account
     products, pr       Lists the set of products that are available under a given contract
     properties, prp    Lists properties available for the current contract and group
     rule-formats, rf   List all available rule formats
     help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "$HOME/.edgerc") [$AKAMAI_EDGERC]
   --debug                  Debug info
   --no-color               Disable color output
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```
