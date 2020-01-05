package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/bwmarrin/discordgo"
    "os/signal"
    "syscall"
    "github.com/arturoguerra/mclogger/internal/logger"
    "github.com/arturoguerra/mclogger/internal/rcon"
)

var (
    TOKEN string
    CMD_CHANNEL string
)

func init() {
    TOKEN = os.Getenv("TOKEN")
    CMD_CHANNEL = os.Getenv("MC_COMMANDS_CHANNEL")
    if TOKEN == "" {
        panic("TOKEN NOT FOUND")
    }
}

func main() {
    dgo, err := discordgo.New("Bot " + TOKEN)
    if err != nil {
        fmt.Println(err)
        return
    }

    err = dgo.Open()
    if err != nil {
        fmt.Println(err)
        return
    }

    logger.New(dgo)
    dgo.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate) {
        if m.ChannelID == CMD_CHANNEL && strings.HasPrefix(m.Content, "-rcon")  {
           rcon.New(s, m)
        }
    })

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    dgo.Close()
}
