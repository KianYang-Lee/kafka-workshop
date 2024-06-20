/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"kafka-workshop/internal/app/consumer"

	"github.com/spf13/cobra"
)

// consumerExerciseCmd represents the consumerExercise command
var consumerExerciseCmd = &cobra.Command{
	Use:   "consumerExercise",
	Short: "Run Kafka consumer client in workshop exercise",
	Long: `consumerExercise runs a Kafka consumer client in the exercise
file. At the start it should only print out a line to sout and is designed as a
blank space for user to work on. User should code on the file until it become a
working Kafka consumer client`,
	Run: func(cmd *cobra.Command, args []string) {
		consumer.RunExercise()
	},
}

func init() {
	rootCmd.AddCommand(consumerExerciseCmd)
}
