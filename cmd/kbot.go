/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)
var TeleToken string

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		TeleToken = os.Getenv("TELE_TOKEN")
		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check your TELE_TOKEN environment variable and try again. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
				case "hello":
					err = m.Send(fmt.Sprintf("Hello! I'm kbot %s", appVersion))
			}
			return err

		})
		kbot.Start()
	},
}

func init() {
	err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system environment variables")
		}

	rootCmd.AddCommand(kbotCmd)
}
