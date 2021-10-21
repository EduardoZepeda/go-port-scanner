package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	fmt.Printf("Scanning %s...\n", *site)
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			wg.Done()
			conn.Close()
			fmt.Printf("[] Port %d is open\n", port)
		}(i)
	}
	wg.Wait()
}
