package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/MumbleBot/bot"
	"github.com/MumbleBot/functions"
	"github.com/MumbleBot/video"
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
	functions.Bot = b
	bot.Bot = b
	video.Bot = b

	if _, err := os.Stat("temp.mp3"); err == nil {
		os.Remove("temp.mp3")
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
