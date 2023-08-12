package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gc-cli/models"
	"github.com/go-resty/resty/v2"
)

func GetErrorGroups(startPage int, endPage int) *models.GroupStats {
	// pageIteration 0 = 5 // default
	var result = []models.GroupStat{}
	var resp *models.GroupStats
	token := "" // first call without token
	for i := startPage; i < endPage; i++ {
		resp = GetErrorGroupsWithToken(token)
		result = append(result, resp.ErrorGroupStats...)
		token = resp.NextPageToken
		if token == "" {
			// GCP API behaviour; if no other data available
			break
		}
	}
	resp.ErrorGroupStats = result
	return resp

}

// this calls GCP API and gets the response

func GetErrorGroupsWithToken(nextToken string) *models.GroupStats {
	// Get env var or crash if key not present
	var token, present = os.LookupEnv("GCLOUD_ACCESS_TOKEN")
	if !present {
		log.Fatalln(`Token environment variable not set: GCLOUD_ACCESS_TOKEN. Possibly this token can be set with: 
		export GCLOUD_ACCESS_TOKEN=$(gcloud auth print-access-token)`)
	}
	// fmt.Println(token)

	// 	enum": [
	// "PERIOD_UNSPECIFIED",
	// "PERIOD_1_HOUR",
	// "PERIOD_6_HOURS",
	// "PERIOD_1_DAY",
	// "PERIOD_1_WEEK",
	// "PERIOD_30_DAYS"
	// Make call to gcp error reporting API to get group stats
	foo := "https://clouderrorreporting.googleapis.com/v1beta1/projects/sonova-marketing/groupStats?timeRange.period=PERIOD_1_WEEK"
	foo += "&order=LAST_SEEN_DESC"
	if nextToken != "" {
		foo = foo + "&pageToken=" + nextToken
	}
	var result = &models.GroupStats{}
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(token).
		SetResult(result).
		Get(foo)
	// fmt.Println(resp)

	if resp.StatusCode() == 401 {
		log.Fatalln(`Token environment variable not set: GCLOUD_ACCESS_TOKEN. Possibly this token can be set with: 
		export GCLOUD_ACCESS_TOKEN=$(gcloud auth print-access-token)`)
	}
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatalln("Shutting down process")
	}
	if !resp.IsSuccess() {
		fmt.Println(resp.Status())
		log.Fatalln("Failed Network Call to GCP API")
	}
	return result
}

func GetErrorEventsAll(groupId string, startPage int, endPage int) *models.ErrorEvents {
	// pageIteration 0 = 5 // default
	var result = []models.ErrorEvent{}
	token := "" // first call without token
	var resp *models.ErrorEvents
	for i := startPage; i < endPage; i++ {
		resp = getErrorEvents(groupId, token)
		result = append(result, resp.ErrorEvent...)
		token = resp.NextPageToken
		if token == "" {
			// GCP API behaviour; if no other data available
			break
		}
	}
	resp.ErrorEvent = result
	return resp

}

func getErrorEvents(groupId string, nextToken string) *models.ErrorEvents {
	var token, present = os.LookupEnv("GCLOUD_ACCESS_TOKEN")
	if !present {
		log.Fatalln(`Token environment variable not set: GCLOUD_ACCESS_TOKEN. Possibly this token can be set with: 
			export GCLOUD_ACCESS_TOKEN=$(gcloud auth print-access-token)`)
	}

	// Make call to gcp error reporting API to get group stats
	var foo = "https://clouderrorreporting.googleapis.com/v1beta1/projects/sonova-marketing/events?groupId=" + groupId

	errorEvents := &models.ErrorEvents{}
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(token).
		SetResult(errorEvents).
		Get(foo)
	// fmt.Println(resp)

	if !resp.IsSuccess() {
		fmt.Println("Error: ", err)
		fmt.Println(resp.Status())
		log.Fatalln("Failed Network Call to GCP API")
	}
	return errorEvents
}
