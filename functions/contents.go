package functions

import (
	"fmt"

	"layeh.com/gumble/gumble"
)

type Contents struct{}

func (c *Contents) GetName() string {
	return ".contents: Displays the current queue.<br>Usage: .contents"
}

func (c *Contents) Exec(ev *gumble.TextMessageEvent) error {
	fmt.Println(Bot.Queue.GetQueue())
	send(Bot.Queue.GetQueue(), ev.Sender)
	return nil
}
