package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(nowTimeCmd)
	rootCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要計算的時間")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持續時間")

}
