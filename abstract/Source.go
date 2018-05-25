package abstract

import (
	"layeh.com/gumble/gumble"
)

type Source interface {
	Init()
	Regex() string
	NewVideo(length int, title, path, url string, poster *gumble.User) Video
}
