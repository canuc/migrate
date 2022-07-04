package resource_changed

import (
	"hash/fnv"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceResourceChanged() *schema.Resource {
	return &schema.Resource{
		Read: readHasChanges,

		Schema: map[string]*schema.Schema{
			"db_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				Description: "The current version of the project's db.",
				Required:    true,
				Computed:    false,
			},
			"changed": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    false,
				Default:     nil,
				Description: "Has the db_version in state changed since last run?",
				Computed:    true,
			},
		},
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// dataSourceAwsAmiDescriptionRead performs the AMI lookup.
func readHasChanges(d *schema.ResourceData, meta interface{}) error {
	changed := d.HasChanges("db_version")
	d.Set("changed", changed)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
