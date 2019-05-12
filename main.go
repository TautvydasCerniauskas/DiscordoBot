package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file", err)
	}
	token := os.Getenv("TOKEN")
	d, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Failed to create discord session", err)
	}

	d.AddHandler(handleCmd)
	err = d.Open()
	if err != nil {
		fmt.Println("Unable to establish connection", err)
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	d.Close()
}
