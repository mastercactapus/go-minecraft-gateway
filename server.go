package main

import (
	"./server"
	"bufio"
	"fmt"
	"net"
)

// ENCRYPTION = invalid stream!!!

func main() {

	s, err := Server.NewServer(":3000")

	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	return
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Printf("Error creating listener: %s\n", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			break
		}

		go handleConnection(conn)
	}

	ln.Close()
}

func pipePackets(src net.Conn, dest net.Conn, name string) {
	srcReader := bufio.NewReader(src)
	for {

		packet, err := NextPacket(srcReader)
		if err != nil {
			fmt.Printf("%s: Error reading data: %s\n", name, err)
			src.Close()
			dest.Close()
			return
		}

		fmt.Printf("%s: payload %d bytes\n", name, len(packet.Payload))
		data := packet.Serialize()
		n, err := dest.Write(data)
		if n != len(data) {
			fmt.Printf("%s: Error writing data, cutoff\n", name)
			src.Close()
			dest.Close()
			return
		}
		if err != nil {
			fmt.Printf("%s: Error writing data: %s\n", name, err)
			src.Close()
			dest.Close()
			return
		}
	}

}

func handleConnection(conn net.Conn) {
	fmt.Printf("Got connection from %s\n", conn.RemoteAddr)
	srv, err := net.Dial("tcp", "127.0.0.1:25566")
	if err != nil {
		fmt.Printf("Could not connect to minecraft server: %s\n", err)
		conn.Close()
		return
	}

	go pipePackets(conn, srv, "C->S")
	go pipePackets(srv, conn, "S->C")

}
