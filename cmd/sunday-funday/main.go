package main

import (
	"fmt"
	"os"

	"github.com/logikone/sunday-funday/server"
)

func main() {
	s := server.Server{}

	if err := s.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
