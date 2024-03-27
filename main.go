package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	prefix := "."
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		argsaux := strings.Split(m.Content, "")
		// args := strings.Split(m.Content, " ")
		if argsaux[0] != prefix {
			return
		}
	})

	sess.Identify.Intents = discordgo.IntentsAll

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()
	fmt.Println("El bot esta online")
}
