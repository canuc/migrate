package provider

import (
	"github.com/canuc/terraform-migrate-provider/resource_changed"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		DataSourcesMap: map[string]*schema.Resource{
			"migrate_resource_changed": resource_changed.DataSourceResourceChanged(),
		},
	}
}
