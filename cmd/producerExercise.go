/*
Copyright © 2024 KianYang Lee <ken.lee.kianyang@gmail.com>
*/
package cmd

import (
	"kafka-workshop/internal/app/producer"

	"github.com/spf13/cobra"
)

// producerExerciseCmd represents the producerExercise command
var producerExerciseCmd = &cobra.Command{
	Use:   "producerExercise",
	Short: "Run Kafka producer client in workshop exercise",
	Long: `producerExercise runs a Kafka producer client in the exercise
file. At the start it should only print out a line to sout and is designed as a
blank space for user to work on. User should code on the file until it become a
working Kafka producer client`,
	Run: func(cmd *cobra.Command, args []string) {
		producer.RunExercise()
	},
}

func init() {
	rootCmd.AddCommand(producerExerciseCmd)
}
