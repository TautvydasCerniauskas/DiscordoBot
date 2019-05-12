package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Token Discord bot token
var Token = "NTc2NDg2NDUyNjUzOTgxNjk3.XNXNvw.uL5hcU42pnz74cyQ9cDSboepSNk"

func main() {
	d, err := discordgo.New("Bot " + Token)
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
