package office365

import (
	"context"
	"log"
	"os"
	"terraform-provider-office365/client"
	t "terraform-provider-office365/token"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OFFICE365_CLIENT_ID", ""),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OFFICE365_CLIENT_SECRET", ""),
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OFFICE365_TENANT_ID", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"office365_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"office365_user": dataSourceUsers(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Println("provider.go called")
	var diags diag.Diagnostics
	clienId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	tenantId := d.Get("tenant_id").(string)
	if clienId == "" || clientSecret == "" || tenantId == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Please Re-check you credentials",
			Detail:   "Please Re-check you credentials",
		})
		return nil, diags
	}
	err := t.GetToken(clienId, clientSecret, tenantId)
	if err != nil {
		os.Setenv("bearer", "")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
			Detail:   "please enter valid credentials",
		})
		return nil, diags
	}
	bearer := os.Getenv("bearer")
	return client.NewClient(bearer), diags
}
