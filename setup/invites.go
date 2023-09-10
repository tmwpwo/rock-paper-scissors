package handlers

import "github.com/bwmarrin/discordgo"

// Formats the user in a readable format
func formatUser(u *discordgo.User) string {
	return u.Username + "#" + u.Discriminator
}

// Generic message format for errors
func ErrorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}

// Generic message format for successful operations
func successMessage(title string, message string) string {
	return "✅  **" + title + "**\n" + message
}

func sendDirectInvite(s *discordgo.Session, m *discordgo.MessageCreate, recipient *discordgo.User) {
	if m.Author.ID == recipient.ID {
		s.ChannelMessageSend(m.ChannelID, ErrorMessage("Invalid recipient", "Cannot play against yourself!"))
				return
	}

	if recipient.Bot {
		s.ChannelMessageSend(m.ChannelID, ErrorMessage("Invalid recipient", "Cannot play against bot!"))
		return
	}

	dm, err := s.UserChannelCreate(recipient.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ErrorMessage("Bot error", "Error creating direct message."))
		return
	}

	invite, err := s.ChannelMessageSendEmbed(dm.ID, &discordgo.MessageEmbed{
		Title:       "Rock Paper Scissors game invite from " + formatUser(m.Author),
		Description: "Click the  ✅  to accept this invitation, or the  ❌  to deny.",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "invite:" + m.Author.ID,
		},
	})

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ErrorMessage("Bot error", "Error sending invite."))
		return
	}

	s.MessageReactionAdd(dm.ID, invite.ID, "✅")
	s.MessageReactionAdd(dm.ID, invite.ID, "❌")

	s.ChannelMessageSend(m.ChannelID, successMessage("Success", "Invite sent to "+formatUser(recipient)+"!"))
}

func InviteCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	recipients := m.Mentions
	if len(recipients) == 1 {
		sendDirectInvite(s, m, recipients[0])
	} else if len(recipients) > 1 {
		s.ChannelMessageSend(m.ChannelID, ErrorMessage("Invalid invite", "Cannot invite multiple players!"))
	}
}
