# Rootly Terraformer (DEPRECATED)

## ⚠️ Please use [Terraform import blocks](https://developer.hashicorp.com/terraform/language/import/generating-configuration) to generate configuration

Generate Terraform configuration from Rootly's API. Terraform in reverse.

[![asciicast](https://asciinema.org/a/630898.svg)](https://asciinema.org/a/630898)

## Usage

### 1. Install

#### Homebrew

    brew tap rootlyhq/homebrew-tap
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

    terraformer-rootly import rootly --resources=*

Or import specific resources:

    terraformer-rootly import rootly --resources=environment,severity

See all available resources using `terraformer-rootly import rootly list`

### Next steps

Generated `.tfstate` files need to be upgraded to Terraform 0.13+ format:

    terraform state replace-provider -auto-approve -- -/rootly rootlyhq/rootly

## Development

### Upgrade Rootly API support

Run `make build`, commit changes, and tag with the next semantic version.
