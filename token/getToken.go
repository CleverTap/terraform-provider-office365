package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetToken(clientID, clientScret, tenantID string) error {
	log.Println("geting Token..")
	url_path := "https://login.microsoftonline.com/" + tenantID + "/oauth2/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientScret)
	data.Set("resource", "https://graph.microsoft.com")
	encodedData := data.Encode()
	req, _ := http.NewRequest("POST", url_path, strings.NewReader(encodedData))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("something is wrong as status code is %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	respone := TokenResponse{}
	json.Unmarshal(body, &respone)
	if respone.AccessToken == "" {
		return fmt.Errorf("Please check your credentials")
	}
	bearer := "Bearer " + respone.AccessToken
	os.Setenv("bearer", bearer)
	log.Println("Token aquired")
	return nil
}
