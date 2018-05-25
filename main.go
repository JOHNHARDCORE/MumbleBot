package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/MumbleBot/bot"
	"github.com/MumbleBot/functions"
	"github.com/MumbleBot/source"
	"github.com/MumbleBot/videos"
	_ "layeh.com/gumble/opus"
)

var terminate = make(chan bool)
var ip = flag.String("ip", "", "Server Address, format is ip:port")
var name = flag.String("name", "DJ", "Bot Name (might not allow certain characters)")
var cert = flag.String("cert", "", "Path to cert file (file.pem)")
var key = flag.String("key", "", "Path to key file (file.pem)")

func main() {
	flag.Parse()

	b := bot.NewBot()
	b.Commands = functions.GetCommands()
	b.Sources = source.GetSources()
	functions.Bot = b
	bot.Bot = b
	videos.Bot = b

	if _, err := os.Stat("res/media/temp.mp3"); err == nil {
		os.Remove("res/media/temp.mp3")
	}

	err := b.Connect(*ip, *name, *cert, *key)
	if err != nil {
		panic(err)
	}
	fmt.Println("injected")

	//Main Loop
	/* This doesn't work--ctrl+c out to exit */
	// nvm maybe it does now i just tested it once
	go evaluate()
	for {
		select {
		case <-terminate:
			fmt.Printf("Exiting\n")
			return
		default:
		}
	}

}

func evaluate() {
	var str string
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Scanf("%s", &str)
		if str == "exit" {
			terminate <- true
			return
		}
	}

}
