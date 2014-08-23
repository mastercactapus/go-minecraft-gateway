package main

import (
	"bufio"
	"fmt"
	"net"
)

// ENCRYPTION = invalid stream!!!

func main() {
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

type Packet struct {
	Length     int32
	LengthData []byte
	Type       int32
	RawData    []byte
	Data       []byte
}

func ReadVarint(reader *bufio.Reader) (int32, []byte, error) {
	shift := byte(0)
	result := int32(0)
	size := 0

	buf := make([]byte, 16)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			return result, buf, err
		}
		buf[size] = b
		size++

		result |= int32(b&0x7f) << shift
		if (b & 0x80) == 0 {
			return result, buf[:size], nil
		}
		shift += 7

		if shift >= 64 {
			panic("varint is too big")
		}
	}
}

func readPacket(srcReader *bufio.Reader, name string) (*Packet, error) {
	packet := new(Packet)

	packetLen, lData, err := ReadVarint(srcReader)
	if err != nil {
		return packet, err
	}
	packet.Length = packetLen
	packet.LengthData = lData
	fmt.Printf("%s: read %d bytes\n", name, len(lData))

	pos := 0
	packet.RawData = make([]byte, packet.Length)
	for {
		read, err := srcReader.Read(packet.RawData[pos:])
		if err != nil {
			return packet, err
		}
		if read == 0 {
			panic("wahh, shits broke")
		}
		pos += read
		if int32(pos) == packet.Length {
			break
		}
	}

	fmt.Printf("%s: read %d bytes, type '%d'\n", name, packet.Length, packet.Type)
	return packet, nil
}

func writePacket(packet *Packet, dest net.Conn, name string) error {
	wrote, err := dest.Write(packet.LengthData)
	if err != nil {
		return err
	}
	fmt.Printf("%s: wrote %d bytes\n", name, wrote)
	if wrote != len(packet.LengthData) {
		panic("wahhh, couldn't write all the data!")
	}

	wrote, err = dest.Write(packet.RawData)
	fmt.Printf("%s: wrote %d bytes\n", name, wrote)
	if err != nil {
		return err
	}
	if int32(wrote) != packet.Length {
		panic("wahhh, couldn't write all the data!")
	}

	return nil
}

func pipePackets(src net.Conn, dest net.Conn, name string) {
	srcReader := bufio.NewReader(src)
	for {

		packet, err := readPacket(srcReader, name)
		if err != nil {
			fmt.Printf("Error piping (reading) data: %s\n", err)
			src.Close()
			dest.Close()
			return
		}

		err = writePacket(packet, dest, name)
		if err != nil {
			fmt.Printf("Error piping (writing) data: %s\n", err)
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

	go pipePackets(conn, srv, "client->server")
	go pipePackets(srv, conn, "server->client")

}
