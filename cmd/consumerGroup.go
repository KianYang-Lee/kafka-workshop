/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"fmt"
	"kafka-workshop/internal/app/consumer"
	"log"

	"github.com/spf13/cobra"
)

// consumerGroupCmd represents the consumerGroup command
var consumerGroupCmd = &cobra.Command{
	Use:   "consumerGroup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consumerGroup called")
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			log.Fatalln(err)
		}
		sleep, err := cmd.Flags().GetInt64("sleep")
		if err != nil {
			log.Fatalln(err)
		}
		consumer.RunGroups(num, sleep)
	},
}

func init() {
	rootCmd.AddCommand(consumerGroupCmd)
	consumerGroupCmd.Flags().IntP("num", "n", 1, "number of consumers")
	consumerGroupCmd.Flags().Int64P("sleep", "s", 0, "number of seconds to sleep (to simulate I/O load)")
}
