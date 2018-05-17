# MumbleBot
Mumble Bot is a bot that allows for users to playback youtube audio, among other things. Heavily WIP

## Features
Supports Youtube Playback

## Requirements
<ul>
  <li><a href=https://github.com/rg3/youtube-dl>Youtube-dl</a></li>
  <li>ffmpeg</li>
</ul>
Pre-compiled binaries are included for Windows<br>

#### Building:

<ul>
  <li><a href=https://github.com/layeh/gumble>Gumble</a></li>
  <li><a href=https://github.com/layeh/gopus>Gopus</a></li>
  <li><a href=https://github.com/kkdai/youtube>Youtube</a></li>
  <li><span>Go 1.5 or later (any version should be fine?)</span></li> 
</ul>


Windows users need mingw-64 to compile
## Installation
Install Requirements
```
go get github.com/Mattrlearned/MumbleBot
```
```
Place youtube-dl in same directory as the compiled binary
```
## Building
Navigate to folder with main.go
```
go build
```
```
Windows users: layeh.com/gopus/opus_nonshared.go will need mingw added to target platform (line 5)
```
## Usage
Launch the exe from a terminal (such as cmd or bash or what have you).<br>
To exit: type 'exit' into the terminal and send
## Options
```
Flags:
-ip: Server Address, format is ip:port
-name: Name of bot
-cert: Location of cert file (.pem). Required for registration
-key: Location of key file (.pem). Required for registration
```

## Commands
All commands are prefixed with a '.'
To issue a command: either send message to the channel the bot is in, or directly message it
```
Commands:
.help: list all commands and their usage
```
