package main

import (
	"cum-ai/openai"
	"cum-ai/sheduler"
	"cum-ai/tg"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	openAICLient := openai.NewOpenAIClient(os.Getenv("OPENAI"))
	chatID, err := strconv.Atoi(os.Getenv("CHATID"))
	if err != nil {
		log.Fatalf("parse chat id")
	}
	tgBot := tg.NewTgBot(os.Getenv("TGTOKEN"), int64(chatID))
	worker := sheduler.NewWorker(openAICLient, tgBot, 3*time.Minute)
	worker.Work()
}
