# Terraformer

This is a fork of [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) with added support for the Rootly provider.

## Usage

1. Install the [latest release](https://github.com/rootlyhq/terraformer/releases). Terraform is also required.
2. Prepare a Terraform working directory with the rootly provider installed.
3. Run `terraformer-rootly import rootly --resources=* --compact --path-pattern=generated` to import.
4. Add `source = "rootlyhq/rootly"` to `generated/provider.tf` and change directory to `generated/`.
5. Run `terraform state replace-provider -- -/rootly rootlyhq/rootly` to fix Terraform 0.13+.
6. Run `terraform init` to initialize the imported configuration.
7. Run `terraform plan` to verify imported configuration.

## Demo

[![asciicast](https://asciinema.org/a/Gv8LCrdpGX0mqISHReAQfJV7N.svg)](https://asciinema.org/a/Gv8LCrdpGX0mqISHReAQfJV7N)

## Documentation

Terraformer documentation is available at [github.com/GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer).

## Development

Upgrade the terraform-provider-rootly Go dependency and commit to main branch. A new release will be built automatically using GitHub actions.
