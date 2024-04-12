package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    pingChannel := make(chan string)
    pongChannel := make(chan string)

    wg.Add(2)

    go func() {
        defer wg.Done()
        for {
            select {
            case message := <-pingChannel:
                fmt.Println("Pong", message)
                pongChannel <- "ping"
            case <-pongChannel:
                fmt.Println("Ping")
                pingChannel <- "pong"
            }
        }
    }()

    go func() {
        defer wg.Done()
        pingChannel <- "ping"
    }()

    wg.Wait()
}
