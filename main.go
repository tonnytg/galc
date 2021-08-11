package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	GetRoles()
}

type Roles struct {
	Role []struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Etag        string `json:"etag"`
	} `json:"roles"`
}


// GetRoles POST and return JSON with all roles of project, need export GCP_API_KEY and PROJECT_ID to work
// GCP_API_KEY you can obtain with "gcloud auth print-access-token"
func GetRoles() {

	token := os.Getenv("GCP_API_KEY")
	project := os.Getenv("PROJECT_ID")
	url := fmt.Sprintf("https://iam.googleapis.com/v1/projects/%s/roles", project)

	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}


	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	} else {
		defer resp.Body.Close()
		var roles Roles
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, &roles)

		for i := 0; i < 2; i++{
			fmt.Println("Path:", roles.Role[i].Name)
			fmt.Println("Name:", roles.Role[i].Title)
			fmt.Println("Description:", roles.Role[i].Description)
			fmt.Println("---")
		}
	}
}
