# Rootly Terraformer

This is a fork of [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) with added support for the Rootly provider.

[![asciicast](https://asciinema.org/a/630898.svg)](https://asciinema.org/a/630898)

## Usage

### 1. Installation

#### Homebrew

    brew tap rootlyhq/homebrew-tap git@github.com:rootlyhq/homebrew-tap.git
    brew install terraformer-rootly

#### Download

Terraform 0.13+ is required.

    export ARCH=darwin-arm64 # Mac Apple silicon. For Mac Intel silicon use darwin-amd64. For Linux use linux-amd64.
    curl -LO "https://github.com/rootlyhq/terraformer/releases/download/$(curl -s https://api.github.com/repos/rootlyhq/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-rootly-${ARCH}"
    chmod +x terraformer-rootly-${ARCH}
    sudo mv terraformer-rootly-${ARCH} /usr/local/bin/terraformer-rootly

### 2. Prepare working directory

Prepare a Terraform working directory with the Rootly provider installed.

If starting from scratch create a `versions.tf` file:

```tf
terraform {
  required_providers {
    rootly = {
      source = "rootlyhq/rootly"
    }
  }
}
```

and run `terraform init` to initialize Terraform.

### 3. Import Terraform configuration

Set the `ROOTLY_API_TOKEN` environment variable or use `--token=` CLI flag when running `terraformer-rootly`.

Import all resources:

```sh
terraformer-rootly import rootly --resources=*
```

Or import specific resources:

```sh
terraformer-rootly import rootly --resources=environment,severity
```

See all the command line options using `terraformer-rootly --help`

### 4. Set qualified name for Rootly provider

Generated `provider.tf` files need `source = "rootlyhq/rootly"`.

Generated `.tfstate` files need to be updated to use a qualified name:

```sh
terraform state replace-provider -- -/rootly rootlyhq/rootly
```

### 5. Verify

Use `terraform plan` to verify imported configuration matches your Rootly configuration.

## Documentation

Terraformer documentation is available at [github.com/GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer).

## Development

### Upgrade Rootly API support

Run `make build`, commit changes, and tag with the next semantic version.
