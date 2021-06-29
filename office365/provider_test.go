package office365

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"office365": testAccProvider,
	}
}
func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ", err)
		t.Fatalf("err: %s", err)
	}
}
func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("OFFICE365_TENANT_ID"); v == "" {
		t.Fatal("TENANT_ID must be set for acceptance tests")
	}
	if v := os.Getenv("OFFICE365_CLIENT_SECRET"); v == "" {
		t.Fatal("CLIENT_SECRET must be set for acceptance tests")
	}
	if v := os.Getenv("OFFICE365_CLIENT_ID"); v == "" {
		t.Fatal("CLIENT_ID must be set for acceptance tests")
	}
}
