package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// apiKey shows how to use an API key to authenticate.
func apiKey() error {
	client, err := pubsub.NewClient(context.Background(), "PROJECT_ID", option.WithAPIKey("TOKEN"))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	// Use the authenticated client.
	_ = client

	return nil
}

func getACL() error {
	client, err := pubsub.NewClient(context.Background(), "eng-venture-320213", option.WithAPIKey("AIzaSyBOgoKN8ab-LloliWG6H-lXAqvTTmkyb2A"))
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()
	// Use the authenticated client.
	_ = client

	return nil
}

type Response struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	State       string `json:"state"`
	Disabled    string `json:"disabled"`
}

func main() {
	apiKey()

	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://iam.googleapis.com/v1beta/{parent=projects/*/locations/*}/workloadIdentityPools", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}
