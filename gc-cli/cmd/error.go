/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// errorCmd represents the error command
var errorCmd = &cobra.Command{
	Use:   "error",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("error called")
		// Get env var or crash if key not present

		// groupId := "CL7V2Pyv6IT59gE"
		// errorEvents := api.GetErrorEventsAll(groupId, 0, 5)
		// renderer.RenderErrors(errorEvents.ErrorEvent)

	},
}

func init() {
	viewCmd.AddCommand(errorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// errorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// errorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
