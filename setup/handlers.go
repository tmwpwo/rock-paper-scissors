package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)




const prefix string = "!gobot"

func GeneralHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		// m.GuildID is a server unique identifier
		// if it's equal to zero, it means that it comes from DM's
		if m.GuildID == "" {

			AnswerHandler(s, m)

		}

		if !strings.HasPrefix(m.Content, prefix) {
			return
		}

		// Get the arguments
		args := strings.Split(m.Content, " ")[1:]
		// Ensure valid command
		if len(args) == 0 {
			s.ChannelMessageSend(m.ChannelID, ErrorMessage("Command missing", "D:"))
			return
		}
		log.Println(args)
		switch args[0] {
			case "siemanko":
				SiemankoHandler(s, m)
			case "wyzwanie":
				PromptHander(s, m) 
			case "invite":
				InviteCommandHandler(s, m) 
		}
}

func PromptHander(s *discordgo.Session, m *discordgo.MessageCreate) {

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Println(err)
	}

	if _, exists := Response[channel.ID]; !exists {
		Response[channel.ID] = Answers{
			OriginalChannel: m.ChannelID,
			Option:          "",
		}
		s.ChannelMessageSend(channel.ID, "rock, paper, scissors?")

	} else {
		s.ChannelMessageSend(channel.ID, "hmmmmmmm?")
	}
}

func SiemankoHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	author := discordgo.MessageEmbedAuthor{
		Name: "Bot mowi:",
		URL:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}
	nice := discordgo.MessageEmbed{
		Title:  "hello!",
		Author: &author,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &nice)
}

func AnswerHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	answers, exists := Response[m.ChannelID]
	if !exists {
		return
	}
	if answers.Option == "" {
		answers.Option = m.Content
	}

	log.Println(answers.OriginalChannel)
	s.ChannelMessageSend(m.ChannelID, m.Content)

}
