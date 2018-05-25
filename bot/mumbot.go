package bot

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/MumbleBot/abstract"
	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
	"layeh.com/gumble/gumbleutil"
)

var Bot *MumBot

// MumBot smokes hella weed and holds all the information in one cute struct
type MumBot struct {
	/* TODO: figure out what the hell goes in here*/
	Commands map[string]abstract.ChatFunction
	Sources  []abstract.Source
	Client   *gumble.Client
	Cfg      *gumble.Config
	TLS      *tls.Config
	Queue    *Queue
	Stream   *gumbleffmpeg.Stream
	Volume   float32
}

// NewBot initializes the bot
func NewBot() *MumBot {
	return &MumBot{
		TLS:      new(tls.Config),
		Commands: make(map[string]abstract.ChatFunction, 0),
		Queue:    newQueue(),
		Stream:   new(gumbleffmpeg.Stream),
		Volume:   float32(.4),
	}
}

// Connect initializes the bot and connects to the server
func (bot *MumBot) Connect(ip, name, cert, key string) error {
	listeners := gumbleutil.Listener{
		Connect:     bot.connectHandler,
		TextMessage: bot.messageHandler,
		Disconnect:  bot.disconnectHandler,
	}
	bot.Cfg = gumble.NewConfig()
	bot.Cfg.Attach(listeners)
	bot.Cfg.Username = name
	// refactor this
	if cert != "" && key != "" {
		skip := false
		if _, err := os.Stat(cert); os.IsNotExist(err) {
			fmt.Println("Could not find key. Use openssl to enable registration")
			skip = true
		}
		if _, err := os.Stat(key); os.IsNotExist(err) {
			fmt.Println("Could not find key. Use openssl to enable registration")
			skip = true
		}

		if !skip {
			cert, err := tls.LoadX509KeyPair(cert, key)
			if err != nil {
				return err
			}
			bot.TLS.InsecureSkipVerify = true
			bot.TLS.Certificates = append(bot.TLS.Certificates, cert)
		}
	}
	client, err := gumble.DialWithDialer(new(net.Dialer), ip, bot.Cfg, bot.TLS)
	if err != nil {
		return err
	}

	bot.Client = client
	return nil

}

func (b *MumBot) connectHandler(e *gumble.ConnectEvent) {
	fmt.Printf("Welcome Message: %s\n", *(e.WelcomeMessage))
	fmt.Printf("Max Bitrate: %d\n", *(e.MaximumBitrate))
	//Request stats for every connected user
	for _, v := range e.Client.Users {
		v.RequestStats()
	}
}

func (b *MumBot) disconnectHandler(e *gumble.DisconnectEvent) {
	switch e.Type {
	case gumble.DisconnectKicked:
		fmt.Printf("Kicked for Reason: %s\nExiting",
			e.String)
	case gumble.DisconnectBanned:
		fmt.Printf("Banned for Reason: %s\nExiting",
			e.String)
	case gumble.DisconnectUser:
		fmt.Printf("Disconnected\nExiting")
	}

	os.Exit(0)
	/* TODO: handle reconnect maybe? */
}

func (b *MumBot) messageHandler(e *gumble.TextMessageEvent) {
	if e.TextMessage.Sender == nil {
		fmt.Printf("sender not found\n")
		return
	}
	if e.TextMessage.Message[0] == '<' {
		fmt.Printf("%s sent non-message to channel\n", e.TextMessage.Sender.Name)
		return
	}
	fmt.Printf("Recieved Message from %s: %s\n",
		e.TextMessage.Sender.Name,
		e.TextMessage.Message)
	fmt.Printf("Parsing...\n")
	b.parse(e)

}

func (b *MumBot) parse(ev *gumble.TextMessageEvent) {
	//check to see if its a bot request
	if ev.TextMessage.Message[0] != '.' {
		return
	}

	name := strings.Split(ev.TextMessage.Message, " ")

	fn, ok := b.Commands[name[0]]
	if ok != false {
		fn.Exec(ev)
	} else {
		fmt.Printf("Command not Found\n")
	}
}
