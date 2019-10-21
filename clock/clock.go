// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/djangulo/clockwall/timezone"
)

var (
	port  string
	tzDir string
)

func init() {
	const (
		portUsage    = "port to listen to"
		defaultPort  = "8000"
		tzDirUsage   = "timezone info directory"
		defaultTzDir = "/usr/share/zoneinfo"
	)
	flag.StringVar(&port, "port", defaultPort, portUsage)
	flag.StringVar(&port, "p", defaultPort, portUsage+" (shorthand)")

	flag.StringVar(&tzDir, "zoneinfo", defaultTzDir, portUsage)
	flag.StringVar(&tzDir, "z", defaultTzDir, portUsage+" (shorthand)")
}

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	location, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}

	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("Mon Jan 2 15:04:05\n"))
		// _, err := io.WriteString(c, time.Now().Format("Mon Jan 2 15:04:05 "+tz+"\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()

	tz := os.Getenv("TZ")
	if tz == "" {
		fmt.Fprint(os.Stderr, "TZ not set, exitting\n")
		os.Exit(1)
	}

	if os.PathSeparator == '\\' {
		tz = strings.ReplaceAll(tz, "\\", "/")
	}

	zoneinfo, err := timezone.SystemTimezones(tzDir)
	if err != nil {
		panic(err)
	}

	if !zoneinfo.Validate(tz) {
		fmt.Fprintf(os.Stderr, "invalid timezone provided: %q\n", tz)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "listening on localhost:%s\n", port)

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, tz) // handle connections concurrently
	}
}
