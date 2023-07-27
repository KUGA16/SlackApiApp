package main

import (
    // "fmt"
	"log"
    "os"
	"github.com/slack-go/slack"
)

func main() {
	// SlackAPI関係のトークンを環境変数から取得
	SLACKAPITOKEN := os.Getenv("SLACK_API_TOKEN")
	SLACKCHANNEL1ID := os.Getenv("SLACK_CHANNEL1_ID")

	if SLACKAPITOKEN == "" {
		errorMessages := []string{
			"******************** Error ********************",
			"FILE : main.go",
			"CAUSE: SLACK_API_TOKEN 環境変数が設定されていません。",
			"***********************************************",
			"(´；ω；`)",
		}
		creatLogFile(errorMessages)
	}
	if SLACKCHANNEL1ID == "" {
		errorMessages := []string{
			"******************** Error ********************",
			"FILE : main.go",
			"CAUSE: SLACK_CHANNEL1_ID 環境変数が設定されていません。",
			"***********************************************",
			"(´；ω；`)",
		}
		creatLogFile(errorMessages)
	}

	// SlackAPIクライアントを初期化
	api := slack.New(SLACKAPITOKEN)

	// メッセージ履歴を取得します
	historyParams := &slack.GetConversationHistoryParameters{
		ChannelID: SLACKCHANNEL1ID,
		Count:     10, // 取得するメッセージの数を指定します（ここでは10件）
	}

	messages, err := api.GetConversationHistory(historyParams)
	if err != nil {
		log.Fatalf("メッセージ履歴の取得に失敗しました: %v", err)
	}

	// 取得したメッセージを表示します
	for _, message := range messages.Messages {
		fmt.Printf("ユーザー: %s, メッセージ: %s\n", message.User, message.Text)
	}
}

func creatLogFile(messages []string) {
	for index, message := range messages {
		// TODO: 受け取ったメッセージ配列を1行づつファイルに書き込む


		// もし配列要素の最後のループの場合、異なるログを出力する
		if index == len(messages)-1 {
			log.Fatal(message)
		}
		log.Println(message)
	}
}