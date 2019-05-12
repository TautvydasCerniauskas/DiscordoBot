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
	channels := getChannels(d)
	err = d.Open()
	if err != nil {
		fmt.Println("Unable to establish connection", err)
	}

	messages := listMessages(d, channels[4].ID)
	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i].Content)
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	d.Close()
}

func getGuilds(s *discordgo.Session) []*discordgo.UserGuild {
	guilds, err := s.UserGuilds(100, "", "")
	if err != nil {
		log.Fatal(err)
	}
	return guilds
}

func getChannels(s *discordgo.Session) []*discordgo.Channel {
	for _, guild := range getGuilds(s) {
		channels, err := s.GuildChannels(guild.ID)
		if err != nil {
			log.Fatal(err)
		}
		return channels
	}
	return nil
}

func listMessages(s *discordgo.Session, channelID string) []*discordgo.Message {
	messages, _ := s.ChannelMessages(channelID, 100, "", "", "")
	return messages
}

func getUserStatus(d *discordgo.Session, m *discordgo.MessageCreate, u *discordgo.User) []*discordgo.User {
	return nil
}
