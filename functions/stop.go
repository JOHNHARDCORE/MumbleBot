package functions

import (
	"layeh.com/gumble/gumble"
)

type Stop struct{}

func (s *Stop) GetName() string {
	return ".stop: Tells the bot to stop playing videos.<br>Usage: .stop"
}

func (s *Stop) Exec(ev *gumble.TextMessageEvent) error {
	Bot.Queue.Stop()
	return nil
}
