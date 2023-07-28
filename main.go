package main

import (
    "fmt"
	"log"
    "os"
	"time"
	"regexp"
	"strconv"

	"github.com/slack-go/slack"
)

func main() {
	// SlackAPI関係のトークンを環境変数から取得
	SLACKAPITOKEN := os.Getenv("SLACK_API_TOKEN")
	SLACKCHANNEL1ID := os.Getenv("SLACK_CHANNEL1_ID")
	SLACK_USER_ID := os.Getenv("SLACK_USER_ID")

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
	if SLACK_USER_ID == "" {
	errorMessages := []string{
		"******************** Error ********************",
		"FILE : main.go",
		"CAUSE: SLACK_USER_ID 環境変数が設定されていません。",
		"***********************************************",
		"(´；ω；`)",
	}
	creatLogFile(errorMessages)
	}

	// SlackAPIクライアントを初期化
	slackApi := slack.New(SLACKAPITOKEN)

	// 今月の開始と終了のUNIXエポック時間を取得
	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, 0).Add(-time.Second)
	firstDayUnix := firstDayOfMonth.Unix()
	lastDayUnix := lastDayOfMonth.Unix()

	// 今月のメッセージ履歴を取得
	historyParams := &slack.GetConversationHistoryParameters{
		ChannelID: SLACKCHANNEL1ID,
		Oldest:    fmt.Sprintf("%d", firstDayUnix),
		Latest:    fmt.Sprintf("%d", lastDayUnix),
	}

	history, error := slackApi.GetConversationHistory(historyParams)
	if error != nil {
		errorMessages := []string{
			"******************** Error ********************",
			"FILE : main.go",
			"CAUSE: " + error.Error(),
			"***********************************************",
			"(´；ω；`)",
		}
		creatLogFile(errorMessages)
	}

	// ユーザーごとにメッセージを取得する
	userMsgs := make(map[string]string)
	for _, allMsg := range history.Messages {
		userID := allMsg.User
		// ユーザーIDがすでに存在する場合は改行で連結、存在しない場合は新しく追加
		if _, ok := userMsgs[userID]; ok {
			userMsgs[userID] += allMsg.Text
		} else {
			userMsgs[userID] = allMsg.Text
		}
	}

	// TODO: userMsgsを引数としてsumNumbersBeforeAsterisk関数を呼び出し、戻り値で全ユーザーの合計マップを取得する（ユーザーID、ユーザー名、合計工数）

	// ユーザーごとに合計工数を取得する
	for userID, userMsg := range userMsgs {
		fmt.Println()
		fmt.Println("-----------------------")
		fmt.Println("User ID:", userID)
		fmt.Println(userMsg)
		fmt.Println("工数合計:", sumNumbersBeforeAsterisk(userMsg))
		fmt.Println("-----------------------")
		fmt.Println()
	}

	//total := sumNumbersBeforeAsterisk(msgs)
	//fmt.Println("合計:", total)
}

// 文章の中から工数を取得し、合計を返却する関数
func sumNumbersBeforeAsterisk(text string) int {
	// 文章内の*の前にある数字を抽出する正規表現パターン
	pattern := `(\d+)h\*`
	re := regexp.MustCompile(pattern)
	// 正規表現で文章内の数字を抽出
	matches := re.FindAllStringSubmatch(text, -1)
	// 抽出した数字を合計
	sum := 0
	for _, match := range matches {
		num, err := strconv.Atoi(match[1])
		if err != nil {
			continue // 数字に変換できない場合は無視する
		}
		sum += num
	}

	return sum
}

// エラーのログ出力とログファイル作成を行う関数
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

// 毎月1回バッチで実行
// やり直しもできるようにする（SlackAPP DMにメッセージを送ることで）