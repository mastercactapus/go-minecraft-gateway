package main

import (
	"fmt"
	"github.com/mastercactapus/go-minecraft-gateway/server"
)

func main() {

	s, err := Server.NewServer(":3000")

	if err != nil {
		panic(err)
	}
	fmt.Println(s)

}
