/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"kafka-workshop/internal/app/producer"

	"github.com/spf13/cobra"
)

// streamProducerCmd represents the streamProducer command
var streamProducerCmd = &cobra.Command{
	Use:   "streamProducer",
	Short: "Runs Kafka producer client that send Wiki stream to broker",
	Long: `
The application, when run, will establish an event stream to 
receive stream from Wiki event. It will write every messages that it 
receive to Kafka broker for topic "wiki-test".`,
	Run: func(cmd *cobra.Command, args []string) {
		producer.RunStream()
	},
}

func init() {
	rootCmd.AddCommand(streamProducerCmd)
}
