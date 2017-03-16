package main

import (
	"fmt"
	"github.com/eminom/gstrike/nucleo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//TODO: Give the instructor to go
	nucleo.NewServer().StartServe()
	brk := make(chan os.Signal, 1)
	signal.Notify(brk, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case <-brk:
		break
	}
	fmt.Println("Goodbye")
}
