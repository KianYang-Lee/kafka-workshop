/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"kafka-workshop/internal/app/producer"
	"log"

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
		streamName, err := cmd.Flags().GetString("streamName")
		if err != nil {
			log.Fatalln(err)
		}
		producer.RunStream(streamName)
	},
}

func init() {
	rootCmd.AddCommand(streamProducerCmd)
	streamProducerCmd.Flags().StringP("streamName", "n", "test", "Wiki event stream to read from(same name will be used as topic name)")
}
