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

// Channel struct ChannelType int
type Channel struct {
	// The ID of the channel.
	ID string `json:"id"`

	// The ID of the guild to which the channel belongs, if it is in a guild.
	// Else, this ID is empty (e.g. DM channels).
	GuildID string `json:"guild_id"`

	// The name of the channel.
	Name string `json:"name"`

	// The type of the channel.
	Type discordgo.ChannelType `json:"type"`
}

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
	m := &discordgo.MessageCreate{}
	u := &discordgo.User{}

	getUserStatus(d, m, u)

	d.AddHandler(handleCmd)
	err = d.Open()
	if err != nil {
		fmt.Println("Unable to establish connection", err)
	}
	channels := getSliceOfStructs(d)
	for i := range channels {
		chaan := channels[i]
		fmt.Println(*chaan)
	}
	messages := listMessages(d, m, channels[4].ID)
	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i].Content)
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	d.Close()
}

func getSliceOfStructs(d *discordgo.Session) []*Channel {
	chnls := []*Channel{}
	guilds, err := d.UserGuilds(100, "", "")
	if err != nil {
		log.Fatal(err)
	}
	for _, guild := range guilds {
		fmt.Println(guild.ID, guild.Name)
		channels, err := d.GuildChannels(guild.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, chaan := range channels {
			tempChan := new(Channel)
			tempChan.ID = chaan.ID
			tempChan.GuildID = chaan.GuildID
			tempChan.Name = chaan.Name
			tempChan.Type = chaan.Type
			// fmt.Println(chaan.ID, chaan.Name, chaan.Type, chaan.Messages)
			chnls = append(chnls, tempChan)

		}
	}
	return chnls
}

func listMessages(s *discordgo.Session, m *discordgo.MessageCreate, channelID string) []*discordgo.Message {
	messages, _ := s.ChannelMessages(channelID, 100, "", "", "")

	return messages
}

func getUserStatus(d *discordgo.Session, m *discordgo.MessageCreate, u *discordgo.User) []*discordgo.User {
	return nil
}
