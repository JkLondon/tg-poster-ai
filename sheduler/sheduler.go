package sheduler

import (
	"cum-ai/openai"
	"cum-ai/tg"
	"math/rand"
	"time"
)

type Worker struct {
	openaiApi *openai.OpenAI
	tgBot     *tg.TgBot
	timeDelta time.Duration
}

func NewWorker(openaiApi *openai.OpenAI, tgBot *tg.TgBot, timeDelta time.Duration) *Worker {
	return &Worker{openaiApi: openaiApi, tgBot: tgBot, timeDelta: timeDelta}
}

func (w *Worker) Post(text string) {
	println("start gen")
	news, err := w.openaiApi.CreateNews(text)
	if err != nil {
		println("openaiApi", err.Error())
		return
	}
	println(news)
	err = w.tgBot.PostShitAboutTown(news)
	if err != nil {
		println("tgBot", err.Error())
		return
	}
}

func (w *Worker) Work() {
	towns := []string{"Авдеевка", "Верхние кринжепки", "Работино", "Клещеевка"}
	go w.Post(towns[rand.Intn(len(towns))])
	ticker := time.Tick(w.timeDelta)
	for range ticker {
		go w.Post(towns[rand.Intn(len(towns))])
	}
}
