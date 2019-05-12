package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/greenavenue/discord_bot/lib"
)

func handleCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Username == "greenavenue" {
		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}

		if m.Content == "pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}

		strTok := strings.SplitN(m.Content, " ", 2)
		if strTok[0] == "!md5" {
			s.ChannelMessageSend(m.ChannelID, lib.MD5Hash(strTok[1]))
		}
	}

}
