package main

import (
	"fmt"
	"flag"
	"log"
	"log/syslog"
	"time"
)

func main() {
	var count int

	// get flags and parse
	host := flag.String("host", "localhost", "Host name")
	proto := flag.String("proto", "udp", "Protocol (udp/tcp)")
	port := flag.Int("port", 514, "Port")
	interval := flag.Int("interval", 5, "Interval (seconds)")
	flag.Parse()

	// dial syslog
	syslog, err := syslog.Dial(*proto, fmt.Sprintf("%s:%d", *host, *port), syslog.LOG_ERR, "syslogpinger")
	if err != nil {
		log.Fatal("failed to dial syslog")
	}

	// loop and log
	count = 0
	for {
		syslog.Alert(fmt.Sprintf("ping #%d", count))
		time.Sleep(time.Duration(*interval) * time.Second)
		count++
	}
}
