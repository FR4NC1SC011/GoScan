package cmd

import (
	"fmt"
	"net"
	"sort"
	"time"

	"github.com/fatih/color"
)

func worker(ports, results chan int, address string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", address, p)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
		//color.Cyan("%d open", p)
	}
}

func Scanner(host string) {
	ports := make(chan int, 1000)
	results := make(chan int)

	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, host)
	}

	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(results)
	close(ports)

	sort.Ints(openports)

	color.Cyan("|=|.:Open Ports:.|=|")
	for _, port := range openports {
		color.Green("%d open\n", port)
	}

	fmt.Println("All ports were scanned")
	color.Magenta("|=|.:NMAP Scan Results:.|=|")
	Cmd(openports, host)
	color.Magenta("|| Did you see that we are faster than RustScan :) ||")
}
