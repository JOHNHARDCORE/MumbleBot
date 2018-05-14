package functions

import (
	"layeh.com/gumble/gumble"
)

func send(str string, tar interface{}) {
	switch t := tar.(type) {
	case *gumble.User:
		t.Send(str)
	case *gumble.Channel:
		t.Send(str, false)
	default:
		return
	}
}

func getRandChildChannel(mp gumble.Channels) *gumble.Channel {
	for _, v := range mp {
		return v
	}

	return nil
}
