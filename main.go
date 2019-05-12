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

	guilds, err := d.UserGuilds(100, "", "")
	if err != nil {
		log.Fatal(err)
	}
	userInfo, err := d.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	d.AddHandler(listMessages)
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

func listMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	messages, _ := s.ChannelMessages(m.ChannelID, 100, "", "", "")
	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i].Content)
	}
}
