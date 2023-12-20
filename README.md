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

#### Generate configuration

Import all Rootly resources:

```sh
terraformer-rootly import rootly --resources=*
```

Import specific resources:

```sh
terraformer-rootly import rootly --resources=environment,severity
```

Combine all generated resource definitions into a single `resources.tf` file:

```sh
terraformer-rootly import rootly --resources=* --compact --path-pattern=generated
```

Terraformer documentation is available at [github.com/GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer).

#### Upgrade state

Terraformer generates state for Terraform 0.13, which uses unqualified provider names. Run the following command to fix for Terraform versions above 0.13.

```sh
terraform state replace-provider -- -/rootly rootlyhq/rootly
```
