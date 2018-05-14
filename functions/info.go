package functions

import (
	"fmt"
	"strings"

	"layeh.com/gumble/gumble"
)

type Info struct{}

func (i *Info) GetName() string {
	return ".info: Returns information about selected user.<br>Usage: .info [user]"
}

func (i *Info) Exec(ev *gumble.TextMessageEvent) error {
	arr := strings.Split(ev.TextMessage.Message, " ")[1:]
	usr := strings.Join(arr, " ")

	user := ev.Client.Users.Find(usr)
	sender := ev.Sender
	if user != nil {
		send(fmt.Sprintf("\nInfo for user: %s<br>", user.Name), sender)
		send(fmt.Sprintf("UID: %d<br>", user.UserID), sender)
		send(fmt.Sprintf("Channel: %s<br>", user.Channel.Name), sender)
		send(fmt.Sprintf("Muted? %v<br>", user.Muted), sender)
		send(fmt.Sprintf("Deafened? %v<br>", user.Deafened), sender)
		send(fmt.Sprintf("Comment: %s<br>", user.Comment), sender)
	} else {
		send("User not found\n", sender)
		return fmt.Errorf("%s not found", usr)
	}

	return nil
}
