package source

import (
	"github.com/MumbleBot/abstract"
	"github.com/MumbleBot/bot"
)

var Bot *bot.MumBot

func GetSources() []abstract.Source {
	return []abstract.Source{
		new(SCSource),
		new(YoutubeSource),
	}
}
