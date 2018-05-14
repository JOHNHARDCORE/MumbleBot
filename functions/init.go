package functions

import (
	"github.com/MumbleBot/abstract"
	"github.com/MumbleBot/bot"
)

// Bot allows for package level access to the bot
var Bot *bot.MumBot

// GetCommands initializes a map of functions and returns it
func GetCommands() map[string]abstract.ChatFunction {
	return map[string]abstract.ChatFunction{
		".echo":     new(Echo),
		".info":     new(Info),
		".join":     new(Join),
		".spam":     new(Spam),
		".add":      new(Add),
		".contents": new(Contents),
		".start":    new(Start),
		".stop":     new(Stop),
		".volume":   new(Volume),
		".skip":     new(Skip),
		".help":     new(Help),
		".clear":    new(Clear),
	}
}
