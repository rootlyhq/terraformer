# Rootly Terraformer

This is a fork of [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) with added support for the Rootly provider.

[![asciicast](https://asciinema.org/a/Gv8LCrdpGX0mqISHReAQfJV7N.svg)](https://asciinema.org/a/Gv8LCrdpGX0mqISHReAQfJV7N)

## Usage

### 1. Installation

Terraform 0.13+ is required.

#### Linux

```
curl -LO "https://github.com/rootlyhq/terraformer/releases/download/$(curl -s https://api.github.com/repos/rootlyhq/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-rootly-linux-amd64"
chmod +x terraformer-rootly-linux-amd64
sudo mv terraformer-rootly-linux-amd64 /usr/local/bin/terraformer
```

#### MacOS

```
curl -LO "https://github.com/rootlyhq/terraformer/releases/download/$(curl -s https://api.github.com/repos/rootlyhq/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-rootly-darwin-amd64"
chmod +x terraformer-rootly-darwin-amd64
sudo mv terraformer-rootly-darwin-amd64 /usr/local/bin/terraformer-rootly
```

### 2. Prepare working directory

Prepare a working directory containing `versions.tf`:

```
terraform {
  required_providers {
    rootly = {
      source = "rootlyhq/rootly"
    }
  }
}
```

Set the `ROOTLY_API_TOKEN` environment variable.

### 3. Import Terraform configuration

Import all resources:

```
terraformer-rootly import rootly --resources=*
```

Or import specific resources:

```
terraformer-rootly import rootly --resources=environment,severity
```

See all the command line options using `terraformer-rootly --help`

### 4. Set qualified name for Rootly provider

Generated `provider.tf` files need `source = "rootlyhq/rootly"`.

Generated `.tfstate` files need to be updated to use a qualified name:

```
terraform state replace-provider -- -/rootly rootlyhq/rootly
```

### 5. Verify

Use `terraform plan` to verify imported configuration matches your Rootly configuration.

## Documentation

Terraformer documentation is available at [github.com/GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer).

## Development

### Upgrade Rootly API support

Run `go get -u github.com/rootlyhq/terraform-provider-rootly` and commit to main branch. A new release will be built automatically.
