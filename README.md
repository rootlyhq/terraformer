# Terraformer

This is a fork of [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) with added support for the Rootly provider.

## Usage

1. Install the [latest release](https://github.com/rootlyhq/terraformer/releases). Terraform is also required.
2. Prepare a Terraform working directory with the rootly provider installed.
3. Run `terraformer-rootly import rootly --resources=* --compact --path-pattern=generated` to import.
4. Add `source = "rootlyhq/rootly"` to `generated/provider.tf`.
5. Run `terraform -chdir=generated state replace-provider -- -/rootly rootlyhq/rootly` to fix Terraform 0.13+.
6. Run `terraform -chdir=generated init` to initialize the imported configuration.
7. Run `terraform -chdir=generated plan` to verify imported configuration.

## Documentation

Terraformer documentation is available at [github.com/GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer).
