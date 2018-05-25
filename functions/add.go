package functions

import (
	"fmt"
	"regexp"

	"github.com/MumbleBot/abstract"
	"github.com/kkdai/youtube"
	"layeh.com/gumble/gumble"
)

type Add struct{}

func (a *Add) GetName() string {
	return ".add: Adds a video to the queue.<br>\tUsage: .add [url]"
}

func (a *Add) Exec(ev *gumble.TextMessageEvent) error {
	url := ev.TextMessage.Message[4:]
	fmt.Println("URL: ", url)
	var src abstract.Source
	var matches []string
	for i := range Bot.Sources {
		fmt.Println("trying ", i)
		fmt.Println(Bot.Sources[i].Regex())
		valid := regexp.MustCompile(Bot.Sources[i].Regex())
		matches = valid.FindStringSubmatch(url)
		if len(matches) > 1 {
			fmt.Println("found a match")
			src = Bot.Sources[i]
			break
		}
	}

	if len(matches) == 0 {
		send("Invalid video URL. Try stripping html from link if it's valid (Send message and send via Source Text tab)", ev.Sender)
		return fmt.Errorf("Invalid video URL. Try stripping html from link if it's valid (Send message and send via Source Text tab)")
	}

	y := youtube.NewYoutube(false)
	fmt.Println("Matched: ", matches[0])
	path := "res/media/temp.mp3"
	title := matches[0] + ".mp3"
	id := matches[0]

	err := y.DecodeURL(matches[0])
	if err == nil {
		fmt.Println("retrieved video data")
		title = y.GetTitle()
	} else {
		fmt.Println("couldn't retrieve data or whatever")
		y = nil
	}
	video := src.NewVideo(0, title, path, id, ev.Sender)
	err = Bot.Queue.Add(video)

	if err != nil {
		send(fmt.Sprintln(err), ev.Sender)
		return err
	}
	send(fmt.Sprintf("%s Added %s to queue", ev.Sender.Name, title), ev.Sender.Channel)

	return nil
}
