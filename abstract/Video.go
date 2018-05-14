package abstract

import "layeh.com/gumble/gumble"

type Video interface {
	Play() error
	Download() error
	Remove()
	Title() string
	Url() string
	Path() string
	Poster() *gumble.User
}
