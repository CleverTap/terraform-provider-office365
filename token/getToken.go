package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetToken(clientID, clientScret, tenantID string) (string, error) {
	url_path := "https://login.microsoftonline.com/" + tenantID + "/oauth2/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientScret)
	data.Set("resource", "https://graph.microsoft.com")
	encodedData := data.Encode()
	req, err := http.NewRequest("POST", url_path, strings.NewReader(encodedData))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("something is wrong as status code is %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	respone := TokenResponse{}
	err = json.Unmarshal(body, &respone)
	if err != nil {
		return "", err
	}
	bearer := "Bearer " + respone.AccessToken
	return bearer, nil
}
