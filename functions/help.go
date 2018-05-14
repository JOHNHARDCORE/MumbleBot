package functions

import (
	"layeh.com/gumble/gumble"
)

type Help struct{}

func (h *Help) GetName() string {
	return ".help: Lists available commands and their use.<br>Usage: .help"
}

func (h *Help) Exec(ev *gumble.TextMessageEvent) error {
	cmds := Bot.Commands
	for _, c := range cmds {
		send(c.GetName(), ev.Sender)
	}
	return nil
}
