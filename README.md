# Go Program with Concurrency for "Ping-Pong"

The following Go program demonstrates the use of concurrency, channels, and goroutines to alternately display the words "ping" and "pong":

```go
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
```

## Explanation:

**Package Imports:**

- **fmt:** For text formatting and printing.
- **sync:** For synchronization of goroutines with `sync.WaitGroup`.

**Channel Definition:**

- **pingChannel:** Channel for sending and receiving the "ping" message.
- **pongChannel:** Channel for sending and receiving the "pong" message.

**Goroutine Creation:**

- **go func() {...}:** Creates two goroutines using the `go` keyword.

**Synchronization:**

- **wg.Add(2):** Informs the `sync.WaitGroup` that there are 2 goroutines to be monitored.
- **defer wg.Done():** Decrements the `sync.WaitGroup` counter when the goroutine finishes.

**Channel Communication:**

- **select {...}:** Allows the goroutine to block until a message is available on one of the channels.
  - **case message := <-pingChannel:** Receives a message from `pingChannel` and stores it in the variable `message`.
    - **fmt.Println("Pong", message):** Prints "Pong" followed by the received message.
    - **pongChannel <- "ping":** Sends the "ping" message to `pongChannel`.
  - **case <-pongChannel:** Receives a message from `pongChannel`.
    - **fmt.Println("Ping"):** Prints "Ping".
    - **pingChannel <- "pong":** Sends the "pong" message to `pingChannel`.

**Start of Communication:**

- **go func() { pingChannel <- "ping" }():** Starts the first goroutine by sending "ping" to `pingChannel`, initiating communication.

**Waits for Goroutines to Finish:**

- **wg.Wait():** Blocks the main program until all goroutines (2) have finished their execution.

This program demonstrates how to use channels and goroutines to implement concurrent communication between different parts of the program. The goroutines communicate by alternating the "ping" and "pong" messages through channels, displaying them on the console.
