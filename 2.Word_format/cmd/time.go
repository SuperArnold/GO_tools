package cmd

import (
	"fmt"
	"log"

	"github.com/SuperArnold/GO_tools/2.Word_format/internal/timer"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "時間格式處理",
	Long:  "時間格式處理",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HI~")
	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "取得現在時間",
	Long:  "取得現在時間",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("現在時間為 %s, %d", nowTime.Format("2021-01-01 00:00:00"), nowTime.Unix())
	},
}
