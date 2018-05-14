package functions

import (
	"layeh.com/gumble/gumble"
)

type Clear struct{}

func (c *Clear) GetName() string {
	return ".clear: clears the current playlist (thanks kcd)"
}

func (c *Clear) Exec(ev *gumble.TextMessageEvent) error {
	Bot.Queue.Clear()
	return nil
}
