package office365

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("office365_user.test_user", "job_title", "intern"),
					resource.TestCheckResourceAttr("office365_user.test_user", "mobile_phone", "+91 88216 10 10"),
					resource.TestCheckResourceAttr("office365_user.test_user", "office_location", "131/1105"),
					resource.TestCheckResourceAttr("office365_user.test_user", "preferred_language", "en-US"),
					resource.TestCheckResourceAttr("office365_user.test_user", "surname", "gautam"),
					resource.TestCheckResourceAttr("office365_user.test_user", "display_name", "Test abc"),
					resource.TestCheckResourceAttr("office365_user.test_user", "user_principal_name", "testingbasic@clevertap1.onmicrosoft.com"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
	resource "office365_user" "test_user" {
		 display_name        = "Test abc"
		 mail_nick_name = "Tester"
		 password="testqwerty1@"
		 user_principal_name = "testingbasic@clevertap1.onmicrosoft.com"
		 account_enabled = true
		 given_name = "raghu"
   		 surname = "gautam"
    	 mobile_phone = "+91 88216 10 10"
		 job_title = "intern"
		 office_location = "131/1105"
		 preferred_language = "en-US"
	  }
`)
}
func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("office365_user.test_user", "job_title", "intern"),
					resource.TestCheckResourceAttr("office365_user.test_user", "mobile_phone", "+91 88216 10 10"),
					resource.TestCheckResourceAttr("office365_user.test_user", "office_location", "131/1105"),
					resource.TestCheckResourceAttr("office365_user.test_user", "preferred_language", "en-US"),
					resource.TestCheckResourceAttr("office365_user.test_user", "surname", "gautam"),
					resource.TestCheckResourceAttr("office365_user.test_user", "display_name", "Test abc"),
					resource.TestCheckResourceAttr("office365_user.test_user", "user_principal_name", "testingupdate@clevertap1.onmicrosoft.com"),
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("office365_user.test_user", "job_title", "student"),
					resource.TestCheckResourceAttr("office365_user.test_user", "mobile_phone", "+91 88216 20"),
					resource.TestCheckResourceAttr("office365_user.test_user", "office_location", "131/1105"),
					resource.TestCheckResourceAttr("office365_user.test_user", "preferred_language", "en-US"),
					resource.TestCheckResourceAttr("office365_user.test_user", "surname", "gautam"),
					resource.TestCheckResourceAttr("office365_user.test_user", "display_name", "Test update"),
					resource.TestCheckResourceAttr("office365_user.test_user", "user_principal_name", "testingupdate@clevertap1.onmicrosoft.com"),
				),
			},
		},
	})
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
	
	resource "office365_user" "test_user" {
		display_name        = "Test abc"
		mail_nick_name = "Tester"
		password="testqwerty1123@"
		user_principal_name = "testingupdate@clevertap1.onmicrosoft.com"
		account_enabled = true
		given_name = "raghu"
		surname = "gautam"
		mobile_phone = "+91 88216 10 10"
		job_title = "intern"
		office_location = "131/1105"
		preferred_language = "en-US"
	 }
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
	resource "office365_user" "test_user" {
		display_name        = "Test update"
		mail_nick_name = "Tester"
		password="testqwerty1123@"
		user_principal_name = "testingupdate@clevertap1.onmicrosoft.com"
		account_enabled = true
		given_name = "raghu"
		surname = "gautam"
		mobile_phone = "+91 88216 20"
		job_title = "student"
		office_location = "131/1105"
		preferred_language = "en-US"
	 }
`)
}
