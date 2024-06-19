/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"kafka-workshop/internal/app/consumer"
	"log"

	"github.com/spf13/cobra"
)

// consumerGroupCmd represents the consumerGroup command
var consumerGroupCmd = &cobra.Command{
	Use:   "consumerGroup",
	Short: "Runs Kafka consumer group client",
	Long: `
Runs Kafka consumer group client that reads from "wiki-test" topic
from Kafka broker. User can configure how many consumers to run.
Optional argument can be set to simulate I/O processing using "sleep"
flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			log.Fatalln(err)
		}
		sleep, err := cmd.Flags().GetInt64("sleep")
		if err != nil {
			log.Fatalln(err)
		}
		consumer.RunGroup(num, sleep)
	},
}

func init() {
	rootCmd.AddCommand(consumerGroupCmd)
	consumerGroupCmd.Flags().IntP("num", "n", 1, "number of consumers")
	consumerGroupCmd.Flags().Int64P("sleep", "s", 0, "number of seconds to sleep (to simulate I/O load)")
}
