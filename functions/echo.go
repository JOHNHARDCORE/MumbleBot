package functions

import (
	"layeh.com/gumble/gumble"
)

type Echo struct{}

func (e *Echo) GetName() string {
	return ".echo: Tells the bot to relay a message back to you.<br>Usage: .echo [message]"
}

func (e *Echo) Exec(ev *gumble.TextMessageEvent) error {
	sender := ev.Sender
	send(ev.TextMessage.Message, sender)

	return nil

}
