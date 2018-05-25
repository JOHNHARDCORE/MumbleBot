package functions

import (
	"layeh.com/gumble/gumble"
)

type Forest struct{}

func (f *Forest) GetName() string {
	return ".forest: implies ur being manipulated and asks everyone if they want to play the forest.<br>usage: .forest"
}

func (f *Forest) Exec(ev *gumble.TextMessageEvent) error {
	send(Frst, ev.Sender.Channel)
	return nil
}
