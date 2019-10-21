// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

type Client struct {
	country string
	city    string
	ch      chan []byte
	conn    *net.Conn
}

var mu sync.Mutex

func (c *Client) Receive() {
	// defer wg.Done()
	for {
		buf := make([]byte, 512)
		n, err := (*(*c).conn).Read(buf)
		if err == io.EOF {
			break
		}
		// _, err = w.Write(tmp[:n])
		// if err != nil {
		// 	panic(err)
		// }
		// w.Write([]byte(c.name + "\t"))
		// w.Write(buf[:n])
		c.ch <- buf[:n]
		// fmt.Fprintf(tw, format, c.name, string(buf[:n]))
		// tw.Flush()
		// mu.Unlock()
	}

}

func NewClient(country, city, address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	ch := make(chan []byte, 1)
	return &Client{country, city, ch, &conn}, nil
}

func CloseAll(clients []*Client) {
	for _, client := range clients {
		(*(*client).conn).Close()
	}
}

const format = "%-15v\t%-15v\t%v"

var wg sync.WaitGroup

func main() {
	clients := make([]*Client, 0)
	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, "=")
		var country, city string
		if strings.Contains(split[0], ",") {
			cc := strings.Split(split[0], ",")
			country, city = cc[1], cc[0]
		} else {
			country = split[0]
		}
		client, err := NewClient(country, city, split[1])
		if err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}

	sort.SliceStable(clients, func(i, j int) bool { return (*clients[i]).city < (*clients[j]).city })
	sort.SliceStable(clients, func(i, j int) bool { return (*clients[i]).country < (*clients[j]).country })

	defer CloseAll(clients)
	for _, client := range clients {
		// wg.Add(1)
		go client.Receive()
	}

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	for {
		mu.Lock()
		time.Sleep(1 * time.Second)
		fmt.Fprintf(tw, format+"\n", "Country", "City", "Time")
		fmt.Fprintf(tw, format+"\n", "-------", "---", "----")
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		for _, client := range clients {
			buf := <-client.ch
			fmt.Fprintf(tw, format, client.country, client.city, string(buf))
		}
		mu.Unlock()
		tw.Flush()
	}
	// wg.Wait()

	// for {
	// 	PrintAll(tw, clients)
	// 	time.Sleep(1 * time.Second)
	// }

}

// func PrintAll(tw *tabwriter.Writer, clients []*Client) {
// 	wg.Add(1)

// 	for _, client := range clients {
// 		go client.Receive(tw)
// 		tw.Flush()
// 	}
// }
