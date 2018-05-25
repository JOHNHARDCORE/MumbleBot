package videos

import (
	"fmt"
	"os"
	"os/exec"

	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
)

// Video holds the information about a video
type Video struct {
	// TODO: figure out what
	title  string
	length int
	poster *gumble.User
	path   string
	url    string
	// currentOffset int
}

func NewVideo(length int, title, path, url string, poster *gumble.User) *Video {
	return &Video{
		length: length,
		title:  title,
		path:   path,
		url:    url,
		poster: poster,
	}
}
func (v *Video) Play() error {
	fmt.Println("Attempting to play: ", v.Title())
	file := gumbleffmpeg.SourceFile(v.Path())
	stream := gumbleffmpeg.New(Bot.Client, file)

	Bot.Stream = stream
	err := stream.Play()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(Bot.Volume)
	stream.Volume = Bot.Volume
	stream.Wait()

	return nil
}

func (v *Video) Download() error {
	if v.Path() != "res/media/temp.mp3" {
		if _, err := os.Stat(v.Path()); err != nil {
			fmt.Println("couldn't find ", v.Path())
			return err
		}
		fmt.Println("found the video")
		return nil
	}
	fmt.Println("Downloading: ", v.Url())
	fmt.Println("Path: ", v.Path())
	args := []string{"-x", "--audio-format", "mp3", "--max-filesize", "10m", "--no-playlist", "--audio-quality", "3", "-o", "res/media/temp.mp3", v.Url()}
	cmd := exec.Command("./youtube-dl", args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func (v *Video) Remove() {
	os.Remove(v.path)
}

func (v *Video) Title() string {
	return v.title
}

func (v *Video) Length() int {
	return v.length
}

func (v *Video) Poster() *gumble.User {
	return v.poster
}

func (v *Video) Path() string {
	return v.path
}

func (v *Video) Url() string {
	return v.url
}
