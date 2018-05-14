package functions

import (
	"fmt"

	"layeh.com/gumble/gumble"
)

type Start struct{}

func (s *Start) GetName() string {
	return ".start: Tells the bot to start playing videos.<br>Usage: .start"
}

func (s *Start) Exec(ev *gumble.TextMessageEvent) error {
	fmt.Println("Attempting to start...")
	go Bot.Queue.Start()
	/*
		if vid == nil {
			fmt.Println("Error: no videos in queue")
			send("Error: no videos in queue", ev.Sender)
		}
	*/
	return nil
}
