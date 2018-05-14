package functions

import (
	"fmt"
	"strings"

	"layeh.com/gumble/gumble"
)

type Join struct{}

func (j *Join) GetName() string {
	return ".join: Attempts to make the bot to join the channel you are in. Send in a private message.<br>Usage: .join"
}

func (j *Join) Exec(ev *gumble.TextMessageEvent) error {
	sender := ev.Sender
	self := ev.Client.Self
	if sender == nil {
		return nil
	}
	split := strings.Split(ev.TextMessage.Message, " ")[1:]
	target := strings.Join(split, " ")
	targetChannel := sender.Channel
	currentChannel := self.Channel

	if target != "" {
		targetUsr := ev.Client.Users.Find(target)
		if targetUsr != nil {
			targetChannel = targetUsr.Channel
		}
	}

	if targetChannel == currentChannel {
		return nil
	}
	self.Move(targetChannel)
	send(fmt.Sprintf("Joined Channel %s", self.Channel.Name), sender)

	return nil
}
