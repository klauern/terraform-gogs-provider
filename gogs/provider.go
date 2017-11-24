package gogs

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is the Terraform provider for Gogs.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOGS_TOKEN", nil),
				Description: "Token key for Gogs API access",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOGS_URL", nil),
				Description: "URL to the Gogs Server (https://gogs.user.com:8888)",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gogs_user": resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Token:   d.Get("token").(string),
		BaseURL: d.Get("base_url").(string),
	}

	return config.Client(), nil
}
