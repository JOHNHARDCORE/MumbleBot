package source

import (
	"github.com/MumbleBot/abstract"
	"github.com/MumbleBot/videos"
	"layeh.com/gumble/gumble"
)

type GenericSource struct{}

func (g *GenericSource) Init() {

}
func (g *GenericSource) NewVideo(length int, title, path, url string, poster *gumble.User) abstract.Video {
	return videos.NewVideo(length, title, path, url, poster)
}
func (g *GenericSource) Regex() string {
	return ""
}

/*
func (g *GenericSource) Download(url string) error {
	if _, err := os.Stat(s.Path()); err != nil {
		fmt.Println("couldn't find ", s.Path())
		return err
	} else {
		fmt.Println("found the video")
		return nil
	}
	args := []string{"-x", "--audio-format", "mp3", "--max-filesize", "10m", "--no-playlist", "--audio-quality", "3", "-o", "res/media/temp.mp3", url}
	cmd := exec.Command("./youtube-dl", args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
*/
