/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gc-cli/api"
	"github.com/gc-cli/models"
	"github.com/gc-cli/renderer"
	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Returns the GCP Error Dashboard View",
	Long:  `Returns the GCP Error Dashboard View`,
	Run: func(cmd *cobra.Command, args []string) {

		result := api.GetErrorGroups(0, 1)
		filterErrorGroups(result.ErrorGroupStats)
		// {
		// 	"group": {
		// 	  "name": "projects/sonova-marketing/groups/CPOl4M2Tk_-iCw",
		// 	  "groupId": "CPOl4M2Tk_-iCw",
		// 	  "resolutionStatus": "ACKNOWLEDGED"
		// 	},
		// 	"count": "82",
		// 	"firstSeenTime": "2023-06-15T13:52:11.077876Z",
		// 	"lastSeenTime": "2023-06-28T21:16:44Z",
		// 	"affectedServices": [
		// 	  {
		// 		"service": "prod-de-lgw-http-server",
		// 		"version": "prod-de-lgw-http-server-00057-bav",
		// 		"resourceType": "cloud_run_revision"
		// 	  },
		// 	  {
		// 		"service": "prod-uk-lgw-http-server",
		// 		"version": "prod-uk-lgw-http-server-00048-riw",
		// 		"resourceType": "cloud_run_revision"
		// 	  }
		// 	],
		// 	"numAffectedServices": 2,
		// 	"representative": {
		// 	  "eventTime": "2023-06-23T20:55:59.088041Z",
		// 	  "serviceContext": {
		// 		"service": "prod-uk-lgw-http-server",
		// 		"version": "prod-uk-lgw-http-server-00047-nap",
		// 		"resourceType": "cloud_run_revision"
		// 	  },
		// 	  "message": "Form not found book-hearing-test-hclearer [Error]\n    at FormsProvider.get (/app/dist/services/lead-generation-server/dist/services/formsProvider.js:40:19)\n    at runMicrotasks (\u003canonymous\u003e)\n    at processTicksAndRejections (internal/process/task_queues.js:95:5)\n    at async FormsController.forms (/app/dist/services/lead-generation-server/dist/controllers/forms.controller.js:50:26)",
		// 	  "context": {
		// 		"httpRequest": {
		// 		  "method": "GET",
		// 		  "url": "https://forms.hearingclearer.co.uk/api/v1/forms/book-hearing-test-hclearer?env=",
		// 		  "responseStatusCode": 500,
		// 		  "remoteIp": "94.1.27.63"
		// 		}
		// 	  }
		// 	}
		//   },

		// renderer.RenderDashboard(result)
		renderer.RenderTree()
	},
}

func filterErrorGroups(resp []models.GroupStat) {
	resolutionStatus := "OPEN"
	// 	"enum": [
	// "RESOLUTION_STATUS_UNSPECIFIED",
	// "OPEN",
	// "ACKNOWLEDGED",
	// "RESOLVED",
	// "MUTED"
	// ],
	filteredGroups := []models.GroupStat{}

	for _, elem := range resp {
		if elem.Group.ResolutionStatus == resolutionStatus {
			filteredGroups = append(filteredGroups, elem)
		}
	}
}

func init() {
	viewCmd.AddCommand(dashboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashboardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
