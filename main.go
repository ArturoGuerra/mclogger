package main

import (
    "io"
    "fmt"
    "os"

    "github.com/papertrail/go-tail/follower"
)

func main() {
    t, _ := follower.New("/data/logs/latest.log", follower.Config{
        Whence: io.SeekEnd,
        Offset: 0,
        Reopen: true,
    })

    for line := range t.Lines() {
        fmt.Println(line)
    }

    if t.Err() != nil {
        fmt.Fprintln(os.Stderr, t.Err())
    }
}
