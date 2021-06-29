package office365

import (
	"context"
	"log"
	"terraform-provider-office365/client"
	val "terraform-provider-office365/validate"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"userprincipalname": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: val.ValidateEmail,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mail_nick_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"given_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"job_title": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mail": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"office_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mobile_phone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"preferred_language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"surname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("DataSOurce Read called")
	c := m.(*client.Client)
	var diags diag.Diagnostics
	userName := d.Get("userprincipalname").(string)
	UserInfo, err := c.GetUser(userName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get User Information",
			Detail:   err.Error(),
		})
		return diags
	}
	d.Set("display_name", UserInfo.DisplayName)
	d.Set("job_title", UserInfo.JobTitle)
	d.Set("mail", UserInfo.Mail)
	d.Set("userprincipalname", UserInfo.UserPrincipalName)
	d.Set("mail", UserInfo.Mail)
	d.Set("office_location", UserInfo.OfficeLocation)
	d.Set("mobile_phone", UserInfo.MobilePhone)
	d.Set("preferred_language", UserInfo.PreferredLanguage)
	d.Set("surname", UserInfo.Surname)
	d.SetId(UserInfo.UserPrincipalName)
	return diags
}
