package functions

import (
	"fmt"

	"github.com/MumbleBot/source"
	"layeh.com/gumble/gumble"
)

type Blowout struct{}

func (b *Blowout) GetName() string {
	return ".blowout: blow it out"
}

func (b *Blowout) Exec(ev *gumble.TextMessageEvent) error {
	send(Blow, ev.Sender.Channel)
	if Bot.Queue.Enabled() {
		send("can't blowout mid song idiot", ev.Sender)
		return fmt.Errorf("can't blowout mid song idiot")
	}
	oldV := Bot.Volume
	var src source.GenericSource
	Bot.Queue.Play(src.NewVideo(0, "blowout.mp3", "res/media/blowout.mp3", "blowout.mp3", ev.Sender))
	Bot.Stream.Volume = .6
	Bot.Volume = oldV
	return nil
}
