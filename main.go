package main

import (
	provider "github.com/canuc/terraform-provider-migrationchangedetection/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	opts := &plugin.ServeOpts{ProviderFunc: provider.Provider}
	plugin.Serve(opts)
}
