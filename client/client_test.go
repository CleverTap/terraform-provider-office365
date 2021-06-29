package client

import (
	"log"
	"os"
	"terraform-provider-office365/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	client_id := os.Getenv("OFFICE365_CLIENT_ID")
	client_secret := os.Getenv("OFFICE365_CLIENT_SECRET")
	tenant_id := os.Getenv("OFFICE365_TENANT_ID")
	err := token.GetToken(client_id, client_secret, tenant_id)
	if err != nil {
		log.Fatal(err)
	}
}
func TestClient_NewUser(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *CreatUser
		seedData  *User
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &CreatUser{
				AccountEnabled:    true,
				DisplayName:       "test create",
				MailNickName:      "test",
				UserPrincipalName: "clienttest@clevertap1.onmicrosoft.com",
				PasswordProfile: PasswordProfileModel{
					ForceChangePasswordNextSignIn: false,
					Password:                      "qwertyu1jj@",
				},
				JobTitle:       "intern",
				MobilePhone:    "+91 882164100",
				OfficeLocation: "pune",
				GivenName:      "first",
				Surname:        "second",
			},

			seedData: &User{
				DisplayName:       "test create",
				JobTitle:          "intern",
				Mail:              "",
				UserPrincipalName: "clienttest@clevertap1.onmicrosoft.com",
				GivenName:         "first",
				MobilePhone:       "+91 882164100",
				OfficeLocation:    "pune",
				PreferredLanguage: "",
				Surname:           "second",
			},
			expectErr: false,
		},
		{
			testName: "item already exists",
			newItem: &CreatUser{
				AccountEnabled:    true,
				DisplayName:       "test create",
				MailNickName:      "test",
				UserPrincipalName: "clienttest@clevertap1.onmicrosoft.com",
				PasswordProfile: PasswordProfileModel{
					ForceChangePasswordNextSignIn: false,
					Password:                      "qwertyu1jj@",
				},
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("bearer"))
			item, err := client.CreateUser(*tc.newItem)
			if err == nil {
				item.ObjectId = ""
			}
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.seedData, item)
		})
	}
}
func TestClient_GetUser(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		seedData     map[string]User
		expectErr    bool
		expectedResp *User
	}{
		{
			testName:  "user exists",
			itemName:  "clienttest@clevertap1.onmicrosoft.com",
			expectErr: false,
			expectedResp: &User{
				DisplayName:       "test create",
				JobTitle:          "intern",
				Mail:              "",
				UserPrincipalName: "clienttest@clevertap1.onmicrosoft.com",
				GivenName:         "first",
				MobilePhone:       "+91 882164100",
				OfficeLocation:    "pune",
				PreferredLanguage: "",
				Surname:           "second",
			},
		},
		{
			testName:     "user does not exist",
			itemName:     "testenotexist@clevertap1.onmicrosoft.com",
			seedData:     nil,
			expectErr:    true,
			expectedResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("bearer"))
			item, err := client.GetUser(tc.itemName)
			if err == nil {
				item.ObjectId = ""
			}
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}
func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName      string
		updatedItem   *UpdateUser
		principalname string
		seedData      User
		expectErr     bool
	}{
		{
			testName:      "item exists",
			principalname: "clienttest@clevertap1.onmicrosoft.com",
			updatedItem: &UpdateUser{
				GivenName:   "test update123",
				JobTitle:    "tester123",
				MobilePhone: "8821640100",
				Surname:     "intern",
			},
			seedData: User{
				GivenName:         "test update123",
				JobTitle:          "tester123",
				MobilePhone:       "8821640100",
				Surname:           "intern",
				DisplayName:       "test create",
				Mail:              "",
				UserPrincipalName: "clienttest@clevertap1.onmicrosoft.com",
				OfficeLocation:    "pune",
				PreferredLanguage: "",
			},
			expectErr: false,
		},
		{
			testName:      "item does not exist",
			principalname: "raghu7984@clevertap1.onmicrosft.com",
			updatedItem: &UpdateUser{
				GivenName:   "test update",
				JobTitle:    "tester",
				MobilePhone: "8821640100",
				Surname:     "intern",
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("bearer"))
			err := client.UpdateUser(tc.principalname, *tc.updatedItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetUser(tc.principalname)
			if err == nil {
				item.ObjectId = ""
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.seedData, *item)
		})
	}
}
func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemName  string
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName:  "user exists",
			itemName:  "clienttest@clevertap1.onmicrosoft.com",
			expectErr: false,
		},
		{
			testName:  "user Does not exists",
			itemName:  "raghugautam45656@clevertap1.onmicrosoft.com",
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("bearer"))
			err := client.DeleteUser(tc.itemName)
			log.Println(err)
			if tc.expectErr {
				log.Println("[DELETE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			log.Println("[DELETE ERROR]: ", err)
			assert.NoError(t, err)
		})
	}
}
