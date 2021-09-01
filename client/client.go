package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type User struct {
	DisplayName       string `json:"displayName"`
	JobTitle          string `json:"jobTitle"`
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
	GivenName         string `json:"givenName"`
	MobilePhone       string `json:"mobilePhone"`
	OfficeLocation    string `json:"officeLocation"`
	PreferredLanguage string `json:"preferredLanguage"`
	Surname           string `json:"surname"`
	ObjectId          string `json:"id"`
}

type PasswordProfileModel struct {
	ForceChangePasswordNextSignIn bool   `json:"forceChangePasswordNextSignIn,omitempty"`
	Password                      string `json:"password,omitempty"`
}

type CreatUser struct {
	AccountEnabled    bool                 `json:"accountEnabled"`
	DisplayName       string               `json:"displayName"`
	MailNickName      string               `json:"mailNickname"`
	UserPrincipalName string               `json:"userPrincipalName"`
	PasswordProfile   PasswordProfileModel `json:"passwordProfile"`
	City              string               `json:"city,omitempty"`
	Country           string               `json:"country,omitempty"`
	Department        string               `json:"department,omitempty"`
	GivenName         string               `json:"givenName,omitempty"`
	JobTitle          string               `json:"jobTitle,omitempty"`
	OfficeLocation    string               `json:"officeLocation,omitempty"`
	PostalCode        string               `json:"postalCode,omitempty"`
	PreferredLanguage string               `json:"preferredLanguage,omitempty"`
	State             string               `json:"state,omitempty"`
	StreetAddress     string               `json:"streetAddress,omitempty"`
	Surname           string               `json:"surname,omitempty"`
	MobilePhone       string               `json:"mobilePhone,omitempty"`
	UsageLocation     string               `json:"usageLocation,omitempty"`
	Mail              string               `json:"mail"`
}

type UpdateUser struct {
	AccountEnabled    bool   `json:"accountEnabled"`
	DisplayName       string `json:"displayName,omitempty"`
	MailNickName      string `json:"mailNickname,omitempty"`
	GivenName         string `json:"givenName,omitempty"`
	JobTitle          string `json:"jobTitle,omitempty"`
	Mail              string `json:"mail,omitempty"`
	MobilePhone       string `json:"mobilePhone,omitempty"`
	OfficeLocation    string `json:"officeLocation,omitempty"`
	PreferredLanguage string `json:"preferredLanguage,omitempty"`
	Surname           string `json:"surname,omitempty"`
	City              string `json:"city,omitempty"`
	Country           string `json:"country,omitempty"`
	Department        string `json:"department,omitempty"`
	PostalCode        string `json:"postalCode,omitempty"`
	State             string `json:"state,omitempty"`
	StreetAddress     string `json:"streetAddress,omitempty"`
	UsageLocation     string `json:"usageLocation,omitempty"`
}

type AssignLicenses struct {
	DisabledPlans []string `json:"disabledPlans"`
	SkuId         string   `json:"skuId"`
}

type License struct {
	AddLicenses    []AssignLicenses `json:"addLicenses"`
	RemoveLicenses []string         `json:"removeLicenses"`
}

type Response struct {
	Value []map[string]interface{} `json:"value"`
}

var (
	Errors = make(map[int]string)
)

func init() {
	Errors[400] = "Bad Request, StatusCode = 400"
	Errors[404] = "User Does Not Exist , StatusCode = 404"
	Errors[409] = "User Already Exist, StatusCode = 409"
	Errors[401] = "Unautharized Access, StatusCode = 401"
	Errors[429] = "User Has Sent Too Many Request, StatusCode = 429"
}

type Client struct {
	authToken  string
	httpClient *http.Client
	BaseUrl    string
}

func NewClient(token string) *Client {
	return &Client{
		authToken:  token,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		BaseUrl:    "https://graph.microsoft.com/v1.0/users/",
	}
}

func (c *Client) GetUser(UserId string) (*User, error) {
	principalName := UserId
	URL := c.BaseUrl + principalName
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	req.Header.Set("authorization", c.authToken)
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	recieveduser := User{}
	err = json.Unmarshal(body, &recieveduser)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return &recieveduser, nil
	} else {
		return nil, fmt.Errorf(string(body) + fmt.Sprintf(", StatusCode: %v", res.StatusCode))
	}
}

func (c *Client) CreateUser(userCreateInfo CreatUser) (*User, error) {
	log.Println("Create Called")
	reqb, err := json.Marshal(userCreateInfo)
	if err != nil {
		return nil, err
	}
	URL := c.BaseUrl
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(reqb)))
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	req.Header.Set("authorization", c.authToken)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	log.Println("Create Completed")
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return &user, nil
	} else {
		return nil, fmt.Errorf(string(body) + fmt.Sprintf(", StatusCode: %v", res.StatusCode))
	}
}

func (c *Client) UpdateUser(UserId string, userUpdateInfo UpdateUser) error {
	log.Println("Update called")
	reqb, err := json.Marshal(userUpdateInfo)
	if err != nil {
		return err
	}
	URL := c.BaseUrl + UserId
	req, err := http.NewRequest("PATCH", URL, strings.NewReader(string(reqb)))
	if err != nil {
		return err
	}
	req.Header.Set("authorization", c.authToken)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println("[ERROR]: ", err)
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer res.Body.Close()
	log.Println("Update completed")
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	} else {
		log.Printf("Error my : %v \n Body : %s", Errors[res.StatusCode], res.Body)
		return fmt.Errorf("%v \nbody: %s", res.StatusCode, res.Body)
	}
}

func (c *Client) DeleteUser(UserId string) error {
	log.Println("Delete Called")
	URL := "https://graph.microsoft.com/v1.0/users/" + UserId
	req, err := http.NewRequest("DELETE", URL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("authorization", c.authToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer res.Body.Close()
	log.Println("Delete completed")
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	} else {
		log.Println(Errors[res.StatusCode], err)
		return fmt.Errorf("%s, %v", res.Body, res.StatusCode)
	}
}

func (c *Client) AddRemoveLicenses(userPrincipalName string, licensesArray License) error {
	body, err := json.Marshal(licensesArray)
	if err != nil {
		return err
	}
	URL := c.BaseUrl + userPrincipalName + "/assignLicense"
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(body)))
	if err != nil {
		log.Println("[ERROR]: ", err)
		return err
	}
	req.Header.Set("authorization", c.authToken)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println("[ERROR]: ", err)
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	} else {
		return fmt.Errorf(string(respBody) + fmt.Sprintf(", StatusCode: %v", resp.StatusCode))
	}
}

func (c *Client) GetLicenseDetails(userPrincipalName string) ([]AssignLicenses, error) {
	URL := c.BaseUrl + userPrincipalName + "/licenseDetails"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	req.Header.Set("authorization", c.authToken)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(string(body) + fmt.Sprintf(", StatusCode: %v", resp.StatusCode))
	}
	licensesMap := Response{}
	err = json.Unmarshal(body, &licensesMap)
	if err != nil {
		return nil, err
	}
	licensesValue := licensesMap.Value
	var assignLicenses []AssignLicenses
	for _, license := range licensesValue {
		assignLicense := AssignLicenses{
			SkuId: license["skuId"].(string),
		}
		for _, l := range license["servicePlans"].([]interface{}) {
			servicePlan := l.(map[string]interface{})
			if servicePlan["provisioningStatus"].(string) == "Disabled" {
				assignLicense.DisabledPlans = append(assignLicense.DisabledPlans, servicePlan["servicePlanId"].(string))
			}
		}
		assignLicenses = append(assignLicenses, assignLicense)
	}
	return assignLicenses, nil
}

func (c *Client) IsRetry(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "429") == true {
			return true
		}
	}
	return false
}
