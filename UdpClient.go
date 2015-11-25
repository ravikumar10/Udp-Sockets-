package main

import (
    "net"
    "fmt"
    "bytes"
    "log"
)

func main() {
    var buf bytes.Buffer
    conn, err := net.Dial("udp", "127.0.0.1:2000")
    if err != nil {
        log.Fatalf("Fatal error: %s\n", err)

    }
    for i := 0; i < 10000; i++ {        
        fmt.Fprintf(&buf, "This is packet number %d\n", i)
        if _, err = buf.WriteTo(conn); err != nil {
            log.Printf("Error: %s\n",err)
        }
        buf.Reset()
    }
	fmt.Scanln()
}
