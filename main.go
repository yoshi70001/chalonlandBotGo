package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"gopkg.in/hraban/opus.v2"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	println(token)
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
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
		if m.Member.User.Bot {
			return
		}
		voiceConnect, _ := sess.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, true)

		f, err := os.Open("speech_32kbps_fb.wav")
		if err != nil {

		}
		opusParse, err := opus.NewStream(f)
		if err != nil {

		}
		defer s.Close()
		buf := make([]byte, 16384)
		for {
			n, err := opusParse.Read(buf)
			if err == io.EOF {
				break
			} else if err != nil {

			}
			pcm := buf[:n*2]
			voiceConnect.OpusSend <- pcm

			// send pcm to audio device here, or write to a .wav file

		}
		time.Sleep(2 * time.Second)
		voiceConnect.Disconnect()

	})

	sess.Identify.Intents = discordgo.IntentsAll

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El bot esta online")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = sess.Close()
	if err != nil {
		log.Fatal(err)
	}
}
