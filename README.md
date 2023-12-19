# Terraformer

This is a fork of [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) with added support for the Rootly provider.

## Usage

### Installation

Please download prebuilt binaries from GitHub releases. Terraform is also required.

### Environment

Prepare a working directory with a `versions.tf` requiring the Rootly Terraform provider:

```tf
terraform {
  required_providers {
    rootly = {
      source = "rootlyhq/rootly"
      version = "1.2.9"
    }
  }
}
```

Run `terraform init` to initialize the Terraform lockfile and download dependencies.

### Import

Import all Rootly resources:

```sh
terraformer-rootly import rootly --resources=*
```

Import specific resources:

```sh
terraformer-rootly import rootly --resources=environment,severity
```

For a full list of options:

```sh
terraformer-rootly import rootly --help

Import current state to Terraform configuration from Rootly

Usage:
   import rootly [flags]
   import rootly [command]

Available Commands:
  list        List supported resources for rootly provider

Flags:
  -b, --bucket string         gs://terraform-state
  -C, --compact
  -c, --connect                (default true)
  -x, --excludes strings      environment
  -f, --filter strings
  -h, --help                  help for rootly
  -S, --no-sort               set to disable sorting of HCL
  -O, --output string         output format hcl or json (default "hcl")
  -o, --path-output string     (default "generated")
  -p, --path-pattern string   {output}/{provider}/ (default "{output}/{provider}/{service}/")
  -r, --resources strings     environment
  -n, --retry-number int      number of retries to perform when refresh fails (default 5)
  -m, --retry-sleep-ms int    time in ms to sleep between retries (default 300)
  -s, --state string          local or bucket (default "local")
  -t, --token string          env param ROOTLY_API_TOKEN
  -v, --verbose
```

Terraformer documentation is available at [github.com/GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer).
