package logger

import (
    "io"
    "fmt"
    "os"
    "time"
    "strings"
    "github.com/papertrail/go-tail/follower"
    "github.com/bwmarrin/discordgo"
)


var (
    fconfig follower.Config
    log_channel string
    files []string
)

func init() {
    fconfig = follower.Config{
        Whence: io.SeekEnd,
        Offset: 0,
        Reopen: true,
    }

    rawfiles := os.Getenv("FILE_PATHS")
    log_channel = os.Getenv("MC_LOGS_CHANNEL")
    files = strings.Split(rawfiles, ",")
}

func New(session *discordgo.Session) {
    logs := make(chan follower.Line)

    for _, filename := range files {
        go Log(filename, logs)
    }

    for line := range logs {
        session.ChannelMessageSend(log_channel, line.String())
    }
}

func Log(filename string, logs chan follower.Line) {
    for {
        t, err := follower.New(filename, fconfig)
        if err != nil {
            time.Sleep(20 * time.Second)
        } else {
            for line := range t.Lines() {
                logs <- line
            }

            if t.Err() != nil {
                fmt.Fprintln(os.Stderr, t.Err())
            }
        }
    }
}
