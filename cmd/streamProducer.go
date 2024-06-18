/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"fmt"
	"kafka-workshop/internal/app/producer"

	"github.com/spf13/cobra"
)

// streamProducerCmd represents the streamProducer command
var streamProducerCmd = &cobra.Command{
	Use:   "streamProducer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("streamProducer called")
		producer.RunStream()
	},
}

func init() {
	rootCmd.AddCommand(streamProducerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamProducerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streamProducerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
