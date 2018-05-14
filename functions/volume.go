package functions

import (
	"fmt"
	"strconv"
	"strings"

	"layeh.com/gumble/gumble"
)

type Volume struct{}

func (v *Volume) GetName() string {
	return ".volume: Sets the volume of the bot.<br>Usage: .volume [.01 to 1.0]"
}

func (v *Volume) Exec(ev *gumble.TextMessageEvent) error {
	num := strings.Split(ev.TextMessage.Message, " ")
	if len(num) < 2 {
		send("invalid volume given\nValid ranges for volume is .01 to .6", ev.Sender)
		return fmt.Errorf("Invalid Volume given")
	}
	vol, err := strconv.ParseFloat(num[1], 32)
	fmt.Println(num[1])
	send(fmt.Sprintf("setting volume to: %f", vol), ev.Sender)
	if err != nil {
		send("invalid volume given", ev.Sender)
		return fmt.Errorf("Invalid Volume given")
	}

	if vol <= .01 || vol >= 0.61 {
		send("invalid volume given\nValid ranges for volume is .01 to 0.6", ev.Sender)
		return fmt.Errorf("Invalid Volume given")
	}

	send(fmt.Sprintf("%s set volume to %f", ev.Sender.Name, vol), ev.Sender.Channel)
	if Bot.Stream == nil {
		Bot.Volume = float32(vol)
	} else {
		Bot.Stream.Volume = float32(vol)
	}
	return nil
}
