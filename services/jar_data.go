package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/everest1508/mcserver-db/apimodels"
	"github.com/everest1508/mcserver-db/constants"
	"github.com/everest1508/mcserver-db/db"
	"github.com/everest1508/mcserver-db/models"
	utils "github.com/everest1508/mcserver-db/utils/api"
)

func FetchAndStoreJarData(client *utils.APIClient) {
	log.Println("Cron Started")
	var wg sync.WaitGroup
	var jarTypeResponse apimodels.TypeResponse

	resp, err := client.Get(constants.TYPE_ENDPOINT, nil)
	if err != nil {
		fmt.Printf("Error fetching jar types: %s\n", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal([]byte(resp.Body), &jarTypeResponse)
		if err != nil {
			fmt.Println("Error decoding jar types response:", err)
			return
		}
	}

	for jarType, subTypes := range jarTypeResponse.Response {
		wg.Add(1)
		go func(jarType string, subTypes []string) {
			defer wg.Done()
			fetchAndStoreSubtypeDetails(jarType, subTypes, client)
		}(jarType, subTypes)
	}
	wg.Wait()
	log.Println("Cron completed")
}

func fetchAndStoreSubtypeDetails(jarType string, subTypes []string, client *utils.APIClient) error {
	for _, subType := range subTypes {
		endpoint := strings.ReplaceAll(constants.DETAILS_ENDPOINT, "{type}", jarType)
		endpoint = strings.ReplaceAll(endpoint, "{subType}", subType)

		resp, err := client.Get(endpoint, nil)
		if err != nil {
			return fmt.Errorf("error fetching details for %s-%s: %w", jarType, subType, err)
		}

		if resp.StatusCode == http.StatusOK {
			var subTypeResponse apimodels.SubTypeResponse
			if err := json.Unmarshal([]byte(resp.Body), &subTypeResponse); err != nil {
				return fmt.Errorf("error unmarshaling subtype response for %s-%s: %w", jarType, subType, err)
			}

			for _, file := range subTypeResponse.Response["files"] {
				serverRecord := models.Server{
					Type:        jarType,
					SubType:     subType,
					Version:     file.Version,
					File:        file.File,
					DisplaySize: file.Size.Display,
					ByteSize:    uint(file.Size.Bytes),
					MD5:         file.MD5,
					Built:       file.Built,
					Stability:   file.Stability,
				}
				serverRecord.CreateRecord(db.DB)
			}
		}
	}
	return nil
}
