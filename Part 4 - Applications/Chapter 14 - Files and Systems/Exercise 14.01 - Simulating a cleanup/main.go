package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CTRL-C")
				fmt.Println("Cleanup is occuring")
				cleanUp()
				done <- struct{}{}
			case syscall.SIGKILL:
				fmt.Println()
				fmt.Println("Someone pressed CTRL-Z")
				fmt.Println("Cleanup is occuring")
				cleanUp()
				done <- struct{}{}
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught(ctrl-z, ctrl-c)")
	done <- struct{}{}
	fmt.Println("Out of here")
}

func cleanUp() {
	fmt.Println("Simulating clean up")
	for i := range 10 {
		fmt.Println("Deleting files... Not really.", i)
		time.Sleep(1 * time.Second)
	}
}
