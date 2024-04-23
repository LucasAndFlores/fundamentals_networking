package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// you can see the communication using sudo tcpdump -n -v -i (use ifconfig to see which profile you can use)

func main() {
	hostAndPort := os.Args[1]

	if len(hostAndPort) == 1 {
		fmt.Println("Please, provide host:port to connect to")
		os.Exit(1)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", hostAndPort)

	if err != nil {
		log.Fatalf("An error occurred when resolve udp address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		log.Fatalf("An error occurred when dial up: %v", err)
	}

	// Setup a channel to listen for Ctrl+C signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	// Start a goroutine to listen for the signal
	go func() {
		<-sig
		fmt.Println("\nCtrl+C received. Exiting.")
		conn.Close()
		os.Exit(0)
	}()

	// Continuous loop for sending messages
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		scanner.Scan()
		message := scanner.Text()
		if strings.TrimSpace(message) == "" {
			continue // Skip empty messages
		}

		_, err := conn.Write([]byte(message))

		if err != nil {
			log.Fatalf("An error occurred when writing to udp connection: %v", err)
		}

	}

}
