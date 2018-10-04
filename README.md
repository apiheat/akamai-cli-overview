# Akamai CLI Overview
The Akamai overview is a CLI that wraps Akamai's {OPEN} APIs to let you quickly view account details ( conttracts / groups / etc ).

Should you miss something we *gladly accept patches* :)

CLI uses custom [Akamai API client](https://github.com/apiheat/go-edgegrid)

# Configuration & Installation

## Credentials
Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in edgerc file which example is shown below

>*NOTE:* Default file location is *~/.edgerc*

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change section which is being actively used you can
* change it via `--config parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_CONFIG=/Users/jsmitsh/.edgerc`

In order to change section which is being actively used you can
* change it via `--section parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

>*NOTE:* Make sure your API client do have appropriate scopes enabled

## Installation
The tool can be used as completly standalone binary or in conjuction with akamai-cli 

### Akamai-cli ( recommended )

1.  Execute the following from console
    `akamai install https://github.com/apiheat/akamai-cli-overview`

### Standalone
As part of automated releases/builds you can download latest version from the project release page

# Actions

```shell
NAME:
   akamai-akamai-cli-overview - A CLI to interact with Akamai account information

USAGE:
   akamai-akamai-cli-overview [global options] command [command options] [arguments...]

VERSION:
   3.0.0

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     contracts         List associated account contracts
     cpcodes           List associated contract/group cpcodes
     custom-behaviors  Lists the set of custom XML metadata behaviors configured for you by Akamai representativess
     custom-overrides  Lists the set of custom XML metadata overrides configured for you by Akamai representatives
     edge-hostnames    Lists all edge hostnames available under a contract
     groups            List associated account groups
     products          List associated contract products
     properties        Lists properties available for the current contract and group
     rule-formats      List all available rule formats
     help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/rpieniazek/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --debug value            Debug Level [$AKAMAI_EDGERC_DEBUGLEVEL]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

# Development
In order to develop the tool with us do the following:
1. Fork repository
1. Clone it to your folder ( within *GO* path )
1. Ensure you can restore dependencies by running 
   ```shell
   dep ensure
   ```
1. Make necessary changes
1. Make sure solution builds properly ( feel free to add tests )
   ```shell
   go build -ldflags="-s -w -X main.appVer=v1.2.3 -X main.appName=akamai-cli-overview" -o akamai-cli-overview
   ```
   