/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"kafka-workshop/internal/app/producer"

	"github.com/spf13/cobra"
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Runs sample of Kafka producer client",
	Long: `
Runs Kafka producer client that writes a few sample messages to 
"my-topic" topic. Use to demonstrate how to implement a simple Kafka
producer client using "kafka-go" library.`,
	Run: func(cmd *cobra.Command, args []string) {
		producer.Run()
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)
}
