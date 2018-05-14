package abstract

import "layeh.com/gumble/gumble"

type ChatFunction interface {
	GetName() string
	Exec(*gumble.TextMessageEvent) error
}
