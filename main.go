package main

import (
	"discord/setup"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	session, err := discordgo.New(Code)
	if err != nil {
		log.Println(err)
	}

	session.AddHandler(handlers.GeneralHandler)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	fmt.Println("online...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
