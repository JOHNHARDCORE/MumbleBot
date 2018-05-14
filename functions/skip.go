package functions

import (
	"fmt"

	"layeh.com/gumble/gumble"
)

type Skip struct{}

func (s *Skip) GetName() string {
	return ".skip: voteskip the current video.<br>Usage: .skip"
}

func (s *Skip) Exec(ev *gumble.TextMessageEvent) error {
	send(fmt.Sprintln(Bot.Queue.Skip(ev.Sender)), ev.Sender.Channel)

	return nil
}
