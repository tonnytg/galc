package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	menu := "choose:\n\t--type Roles\n\t--type ServiceAccount\n\t--type ServiceAccounts\n"

	f := flag.String("type", "", menu)
	flag.Parse()

	switch *f {
	case "Roles":
		GetRoles()
	case "ServiceAccount":
		GetServiceAccount()
	case "ServiceAccounts":
		GetServiceAccounts()
	default:
		fmt.Printf(menu)
	}
}

type Roles struct {
	Role []struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Etag        string `json:"etag"`
	} `json:"roles"`
}

type ServiceAccount struct {
	Name           string `json:"name"`
	ProjectID      string `json:"projectId"`
	UniqueID       string `json:"uniqueId"`
	Email          string `json:"email"`
	DisplayName    string `json:"displayName"`
	Etag           string `json:"etag"`
	Description    string `json:"description"`
	Oauth2ClientID string `json:"oauth2ClientId"`
	Disabled       string `json:"disabled"`
}

type ServiceAccounts struct {
	Accounts []struct {
		Name           string `json:"name"`
		ProjectID      string `json:"projectId"`
		UniqueID       string `json:"uniqueId"`
		Email          string `json:"email"`
		DisplayName    string `json:"displayName"`
		Etag           string `json:"etag"`
		Oauth2ClientID string `json:"oauth2ClientId"`
		Description    string `json:"description,omitempty"`
	} `json:"accounts"`
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
		var role Roles
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, &role)
		fmt.Println("Roles:")

		for i := 0; i < len(role.Role); i++{
			fmt.Printf("Path:\t\t %s \n", role.Role[i].Name)
			fmt.Printf("Name:\t\t %s \n", role.Role[i].Title)
			fmt.Printf("Description:\t %s \n", role.Role[i].Description)
			fmt.Println("---")
		}
	}
}

// GetServiceAccount get info about one service account, you need export vars GCP_API_KEY PROJECT_ID and SERVICE_ACCOUNT
func GetServiceAccount() {
	token := os.Getenv("GCP_API_KEY")
	project := os.Getenv("PROJECT_ID")
	serviceAccount := os.Getenv("SERVICE_ACCOUNT")
	url := fmt.Sprintf("https://iam.googleapis.com/v1/projects/%s/serviceAccounts/%s", project, serviceAccount)

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
		var sc ServiceAccount
		data, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(data, &sc)

		fmt.Println("ServiceAccount:")
		fmt.Printf("Project:\t %s \n", sc.ProjectID)
		fmt.Printf("DisplayName:\t %s \n", sc.DisplayName)
		fmt.Printf("Email:\t\t %s \n", sc.Email)
		fmt.Printf("Description:\t\t %s \n", sc.Description)
		fmt.Println("---")
	}
}

func GetServiceAccounts() {
	token := os.Getenv("GCP_API_KEY")
	project := os.Getenv("PROJECT_ID")
	url := fmt.Sprintf("https://iam.googleapis.com/v1/projects/%s/serviceAccounts", project)

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
		data, _ := ioutil.ReadAll(resp.Body)
		var scs ServiceAccounts
		json.Unmarshal(data, &scs)

		fmt.Println("ServiceAccounts:")
		for i := 0; i < len(scs.Accounts); i++ {
			fmt.Printf("Project:\t %s \n", scs.Accounts[i].ProjectID)
			fmt.Printf("DisplayName:\t %s \n", scs.Accounts[i].DisplayName)
			fmt.Printf("Email:\t\t %s \n", scs.Accounts[i].Email)
			fmt.Printf("Description:\t %s \n", scs.Accounts[i].Description)
			fmt.Println("---")
		}

	}
}