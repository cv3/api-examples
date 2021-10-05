package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	API_BASE_URL = "https://service.commercev3.com/rest"
	API_METHOD   = "POST"
	REST_KEY     = "REST_KEY"
	REST_SECRET  = "REST_SECRET"
	AUTH_SCOPE   = "products" // products, orders, categories, etc.
)

type Auth struct {
	Token   string `json:"access_token"`
	Expires int    `json:"exires_in"`
	Type    string `json:"token_type"`
	Scope   string `json:"scope"`
}

func main() {

	// Auth Request
	// in a real app you would check first to see
	// if you're auth token is still good. Here we just
	// get a new one every time.
	authInfo, err := doAuthRequest()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Data Request
	// in this case we are getting all products
	dataResponse, err := doDataRequest(authInfo.Token)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Use Results
	// in a real app you'll likely have structs representing
	// the JSON data you want to work with, and here you would
	// decode the JSON into those structs, like we did above
	// when we decoded the OAuth2 response into an Auth object.
	// Here we just print the body of the response.
	fmt.Println(string(dataResponse))

}

// doAuthRequest hits the CV3 REST OAuth2 endpoint to get
// a token to use to make data requests. In this case we are
// validating for the "products" endpoint. Tokens are good for
// an hour.
func doAuthRequest() (*Auth, error) {

	// create request
	auth_payload := strings.NewReader("grant_type=client_credentials&scope=" + AUTH_SCOPE)

	client := &http.Client{}
	req, err := http.NewRequest(API_METHOD, API_BASE_URL+"/oauth2/v2.0/token", auth_payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// set headers
	req.SetBasicAuth(REST_KEY, REST_SECRET)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	// extract Auth object from response body
	var a Auth

	err = json.NewDecoder(res.Body).Decode(&a)
	if err != nil {
		fmt.Println(err)
	}

	return &a, nil
}

// doDataRequest takes a valid token for the products
// endpoing and requests an export of all products. Here
// we just return the body of the response.
func doDataRequest(token string) ([]byte, error) {

	// create JSON payload
	data_payload := `{
		"data": {
				"exportProducts": {
					"export_by_range": {
						"start": 1
					}
				}
			}
		}
	`

	// create client
	client := &http.Client{}
	req, err := http.NewRequest(API_METHOD, API_BASE_URL+"/products", strings.NewReader(data_payload))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// set bearer authorization header with the OAuth token value
	req.Header.Add("Authorization", "Bearer "+token)

	// send request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	// get response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
