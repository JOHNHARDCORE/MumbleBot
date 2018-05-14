package functions

import (
	"bytes"
	"math/rand"

	"layeh.com/gumble/gumble"
)

type Spam struct{}

func (s *Spam) GetName() string {
	return ".spam: generate a random string of chararcters. its not really useful.<br>Usage: .spam"
}

func (s *Spam) Exec(ev *gumble.TextMessageEvent) error {
	n := (rand.Int() % 31) + 1
	var res bytes.Buffer
	//generates a random string of n characters from A-Z
	for i := 0; i < n; i++ {
		res.WriteByte(byte(rand.Int()%('Z'-'A')) + 'A')
	}

	send(res.String(), ev.Sender)
	return nil
}
