package office365

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)
func TestAccUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.office365_user.test_user", "display_name", "Melissa Darrow"),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "job_title", "Marketing Director"),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "mail", ""),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "mobile_phone", "+1 206 555 0110"),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "office_location", "131/1105"),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "preferred_language", "en-US"),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "surname", "Darrow"),
					resource.TestCheckResourceAttr("data.office365_user.test_user", "userprincipalname", "testread@clevertap1.onmicrosoft.com"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(`	  
	    data "office365_user" "test_user" {
               userprincipalname="testread@clevertap1.onmicrosoft.com"
		 }
	`)
}
