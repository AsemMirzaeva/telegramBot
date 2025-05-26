package main

import (
	"flag"
	"log"
	tgClient "telegrambot/clients/telegram"
	event_consumer "telegrambot/consumer/event-consumer"
	"telegrambot/events/telegram"
	"telegrambot/storage/files"
)

// token = flags.Get(token)
// tgClient = telegram.New(token)
// fetcher = fetcher.New(tgClient)
// processor = processor.New(tgClient)
// consumer.Start(fetcher, processor)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
	// token = "8138152240:AAGwa-89PrRGIsDgZxR4__m-ONvnDz800jo"
)

func main() {

	evetsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(evetsProcessor, evetsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
