package gearlist

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Service struct {
	Sheets *sheets.Service
	Drive  *drive.Service
}

var singletonService *Service

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)

	if err != nil {
		tok, err = tokenFromAuthCode(config)
		saveToken(tokFile, tok)
	}

	if tok == nil {
		printAuthUrl(config)
		return nil
	}

	return config.Client(context.Background(), tok)
}

// Parse token from local file with authCode
func tokenFromAuthCode(config *oauth2.Config) (*oauth2.Token, error) {
	f, err := os.Open("authcode.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	authCode := bufio.NewScanner(f).Text()
	return config.Exchange(context.TODO(), authCode)
}

// Display auth url to get authCode
func printAuthUrl(config *oauth2.Config) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("AuthURL: \n%v\n", authURL)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetGoogleService() *Service {
	if singletonService != nil {
		return singletonService
	}

	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	sheetsInstance, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	driveInstance, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	singletonService := &Service{
		Sheets: sheetsInstance,
		Drive:  driveInstance,
	}

	return singletonService
}

func GetHomeSpreadsheet() *drive.File {
	service := GetGoogleService()
	resp, err := service.Drive.Files.List().Q("name = 'HoME'").Fields("files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	var homeSheet *drive.File
	for _, file := range resp.Files {
		if file.Name == "HoME" {
			homeSheet = file
		}
	}
	if homeSheet == nil {
		log.Fatalf("Unable to get HoME spreadsheet")
	}

	return homeSheet
}
