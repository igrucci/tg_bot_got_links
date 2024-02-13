package main

import (
	"flag"
	"log"
	tgClient "tg-bot-test/clients/telegram"
	_ "tg-bot-test/cunsumer/event-consumer"
	event_consumer "tg-bot-test/cunsumer/event-consumer"
	"tg-bot-test/events/telegram"
	"tg-bot-test/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	s := files.New(storagePath)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, "your_tokengit"),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

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

//
//const (
//	tgBotHost   = "api.telegram.org"
//	StoragePath = "storage"
//	batchSize   = 100
//)
//
//func main() {
//
//	// token = flags.Get(token)
//
//	// для общения с апи тг
//	//tgClient := telegram.New(tgBotHost, mustToken())
//
//	eventsProcessor := telegram.New(
//		tgClient.New(tgBotHost, mustToken()),
//		files.New(StoragePath),
//	)
//	log.Print("service started")
//
//	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
//
//	if err := consumer.Start(); err != nil {
//		log.Fatal("service is stopped", err)
//	}
//}
//
//// чтобы получать
//// fetcher = fetcher.New(tgClient)
//
//// чтобы обрабатываьб
//// processor = processor.New(tgClient)
//
////получает события и обрабатывает их
//// consumer =
//
//func mustToken() string {
//	token := flag.String(
//		"token-bot-token",
//		"",
//		"token for access to telegram bot")
//	flag.Parse()
//	if *token == "" {
//		log.Fatal("token is not specified")
//	}
//	return *token
//}
