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
   akamai-overview - A CLI to interact with Akamai account information

USAGE:
   akamai-overview [global options] command [command options] [arguments...]

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     contracts       List associated account contracts
     edge-hostnames  Lists all edge hostnames available under a contract
     groups          List associated account groups
     products        List associated contract products
     properties      Lists properties available for the current contract and group
     help, h         Shows a list of commands or help for one command

   CPCodes actions:
     cpcodes  List associated contract/group cpcodes

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/rpieniazek/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```
