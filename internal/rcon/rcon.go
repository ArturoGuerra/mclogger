package rcon

import (
    mcrcon "github.com/Kelwing/mc-rcon"
    "github.com/bwmarrin/discordgo"
    "strings"
    "fmt"
    "os"
)

var (
    Host string
    Port string
    Password string
)

func init() {
    Host = os.Getenv("RCON_HOST")
    Port = os.Getenv("RCON_PORT")
    Password = os.Getenv("RCON_PASSWORD")
}

func getcommand(message string) string {
    return strings.TrimPrefix(message, "/")
}

func New(session *discordgo.Session, message *discordgo.MessageCreate) {
    conn := new(mcrcon.MCConn)
    err := conn.Open(fmt.Sprintf("%s:%s", Host, Port), Password)
    if err != nil {
        fmt.Println("Error opening connection to mc server")
        return
    }


    defer conn.Close()

    err = conn.Authenticate()
    if err != nil {
        fmt.Println("Error authenticating with mc server")
        return
    }

    command := getcommand(message.Content)
    resp, err := conn.SendCommand(command)
    if err != nil {
        session.ChannelMessageSend(message.ChannelID, err.Error())
    } else {
        var msg string
        if resp == "" {
            msg = "Command executed successfully!"
        } else {
            msg = resp
        }

        session.ChannelMessageSend(message.ChannelID, msg)
    }
}
