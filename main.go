package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"example_resource": resourceExample(),
				},
			}
		},
	})
}

func resourceExample() *schema.Resource {
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			d.SetId("example_id")
			return nil
		},
		Read:   func(d *schema.ResourceData, m interface{}) error { return nil },
		Update: func(d *schema.ResourceData, m interface{}) error { return nil },
		Delete: func(d *schema.ResourceData, m interface{}) error { return nil },
	}
}
