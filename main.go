package main

import (
	"github.com/canuc/terraform-migrate-provider/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	opts := &plugin.ServeOpts{ProviderFunc: provider.Provider}
	plugin.Serve(opts)
}
