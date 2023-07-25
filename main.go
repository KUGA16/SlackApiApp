package main

import (
    "fmt"
	"log"
    "os"
)

func main() {
	// SlackAPI関係のトークンを環境変数から取得
	SLACKAPITOKEN := os.Getenv("SLACK_API_TOKEN")
	SLACKCHANNEL1ID := os.Getenv("SLACK_CHANNEL1_ID")
	if SLACKAPITOKEN == "" {
		log.Println("********** Error **********")
		log.Println("FILE : main.go")
		log.Println("CAUSE: SLACK_API_TOKEN環境変数が設定されていません。")
		log.Println("***************************")
		log.Fatal("(´；ω；`)")
	}
	if SLACKCHANNEL1ID == "" {
		log.Println("********** Error **********")
		log.Println("FILE : main.go")
		log.Println("CAUSE: SLACK_CHANNEL1_ID環境変数が設定されていません。")
		log.Println("***************************")
		log.Fatal("(´；ω；`)")
	}


	fmt.Println(SLACKAPITOKEN)
	fmt.Println(SLACKCHANNEL1ID)
}
