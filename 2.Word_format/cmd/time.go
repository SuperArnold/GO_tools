package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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
		log.Printf("現在時間為 %s, %d", nowTime.Format("2021-01-0100:00:00"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "cal",
	Short: "計算時間",
	Long:  "計算所需時間",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-0215:04:05"

		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "Jan 2, 2006 at 3:04pm "
			}

			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		fmt.Println(calculateTime)
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("結果： %s, %d", t.Format(layout), t.Unix())
	},
}
