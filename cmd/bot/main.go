package main

import (
    "os"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "os/signal"
    "syscall"
    "github.com/arturoguerra/mclogger/internal/logger"
)

var (
    TOKEN string
)

func init() {
    TOKEN = os.Getenv("TOKEN")
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

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    dgo.Close()
}
