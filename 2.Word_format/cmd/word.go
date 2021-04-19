package cmd

import (
	"log"
	"strings"

	"github.com/SuperArnold/GO_tools/2.Word_format/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1 //大寫
	ModeLower            //小寫
)

var str string
var mode int

var desc = strings.Join([]string{
	"選單：",
	"1.轉成大寫",
	"2.轉成小寫",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單字格式轉換",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		default:
			log.Fatalf("Fatal")
		}
		log.Printf("結果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "請說輸入單字內容")
	wordCmd.Flags().IntVarP(&mode, "mode", "m", 0, "請說輸入轉換單字模式")
}
