package videos

import (
	"fmt"
	"os"
	"os/exec"

	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
)

type SoundCloud struct {
	// TODO: figure out what
	title  string
	length int
	poster *gumble.User
	path   string
	url    string
}

func NewSoundCloud(length int, title, path, url string, poster *gumble.User) *SoundCloud {
	return &SoundCloud{
		length: length,
		title:  title,
		path:   path,
		url:    url,
		poster: poster,
	}
}

func (s *SoundCloud) Url() string {
	return s.url
}

func (s *SoundCloud) Poster() *gumble.User {
	return s.poster
}

func (s *SoundCloud) Path() string {
	return s.path
}

func (s *SoundCloud) Remove() {
	os.Remove(s.path)
}

func (s *SoundCloud) Title() string {
	return s.title
}

func (s *SoundCloud) Play() error {
	fmt.Println("Attempting to play: ", s.Title())
	file := gumbleffmpeg.SourceFile(s.Path())
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

func (s *SoundCloud) Download() error {
	if s.Path() != "res/media/temp.mp3" {
		if _, err := os.Stat(s.Path()); err != nil {
			fmt.Println("couldn't find ", s.Path())
			return err
		}
		fmt.Println("found the video")
		return nil
	}
	fmt.Println("Downloading: ", s.Url())
	fmt.Println("Path: ", s.Path())
	args := []string{"-x", "--audio-format", "mp3", "--max-filesize", "10m", "--no-playlist", "--audio-quality", "3", "-o", "res/media/temp.mp3", s.Url()}
	cmd := exec.Command("./youtube-dl", args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
