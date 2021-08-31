package office365

import (
	"context"
	"time"
	"strings"
	"terraform-provider-office365/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserLicense() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user_principal_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"licenses": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled_plans": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
						},
						"skuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceUserLicenseImporter,
		},
		CreateContext: resourceUserLicenseCreate,
		ReadContext:   resourceUserLicenseRead,
		UpdateContext: resourceUserLicenseUpdate,
		DeleteContext: resourceUserLicenseDelete,
	}
}

func resourceUserLicenseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	var diags diag.Diagnostics
	userPrincipalName := d.Get("user_principal_name").(string)
	var assignLicenseArray []client.AssignLicenses
	licenses := d.Get("licenses").(*schema.Set).List()
	for _, v := range licenses {
		license := v.(map[string]interface{})
		disabledPlans := license["disabled_plans"].(*schema.Set).List()
		disabledPlansData := make([]string, len(disabledPlans))
		for i, data := range disabledPlans {
			disabledPlansData[i] = data.(string)
		}
		assignLicense := client.AssignLicenses{
			SkuId:         license["skuid"].(string),
			DisabledPlans: disabledPlansData,
		}
		assignLicenseArray = append(assignLicenseArray, assignLicense)
	}
	licensesArray := client.License{
		AddLicenses: assignLicenseArray,
	}
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		err := c.AddRemoveLicenses(userPrincipalName, licensesArray)
		if err != nil {
			if c.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	d.SetId(userPrincipalName)
	return diags
}

func resourceUserLicenseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		body, err := c.GetLicenseDetails(d.Id())
		if err != nil {
			if c.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		var licenses []interface{}
		for _, v := range body {
			license := make(map[string]interface{})
			license["disabled_plans"] = v.DisabledPlans
			license["skuid"] = v.SkuId
			licenses = append(licenses, license)
		}
		d.Set("licenses", licenses)
		d.Set("user_principal_name", d.Id())
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		if strings.Contains(retryErr.Error(), "ResourceNotFound") == true {
			d.SetId("")
			return diags
		}
		return diag.FromErr(retryErr)
	}
	return diags
}

func resourceUserLicenseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	if d.HasChange("user_principal_name") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't update User Principal Name",
			Detail:   "Can't update User Principal Name",
		})
		return diags
	}
	userPrincipalName := d.Id()
	var assignLicenseArray []client.AssignLicenses
	licenses := d.Get("licenses").(*schema.Set).List()
	for _, v := range licenses {
		license := v.(map[string]interface{})
		disabledPlans := license["disabled_plans"].(*schema.Set).List()
		disabledPlansData := make([]string, len(disabledPlans))
		for i, data := range disabledPlans {
			disabledPlansData[i] = data.(string)
		}
		assignLicense := client.AssignLicenses{
			SkuId:         license["skuid"].(string),
			DisabledPlans: disabledPlansData,
		}
		assignLicenseArray = append(assignLicenseArray, assignLicense)
	}
	var body []client.AssignLicenses
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		body, err = c.GetLicenseDetails(userPrincipalName)
		if err != nil {
			if c.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	checkMap := make(map[string]bool)
	for _, v := range assignLicenseArray {
		checkMap[v.SkuId] = true
	}
	var removeLicenses []string
	for _, v := range body {
		if _, ok := checkMap[v.SkuId]; !ok {
			removeLicenses = append(removeLicenses, v.SkuId)
		}
	}
	licensesArray := client.License{
		AddLicenses:    assignLicenseArray,
		RemoveLicenses: removeLicenses,
	}
	err = c.AddRemoveLicenses(userPrincipalName, licensesArray)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserLicenseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	userPrincipalName := d.Id()
	var body []client.AssignLicenses
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		body, err = c.GetLicenseDetails(userPrincipalName)
		if err != nil {
			if c.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	var removeLicenses []string
	for _, v := range body {
		removeLicenses = append(removeLicenses, v.SkuId)
	}
	licensesArray := client.License{
		RemoveLicenses: removeLicenses,
	}
	err = c.AddRemoveLicenses(userPrincipalName, licensesArray)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func resourceUserLicenseImporter(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*client.Client)
	body, err := c.GetLicenseDetails(d.Id())
	if err != nil {
		return nil, err
	}
	var licenses []interface{}
	for _, v := range body {
		license := make(map[string]interface{})
		license["disabled_plans"] = v.DisabledPlans
		license["skuid"] = v.SkuId
		licenses = append(licenses, license)
	}
	d.Set("licenses", licenses)
	d.Set("user_principal_name", d.Id())
	return []*schema.ResourceData{d}, nil
}
