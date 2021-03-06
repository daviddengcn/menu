// menu.go
package main

import (
	"fmt"
	"os"

	"github.com/daviddengcn/go-colortext"
	"github.com/kylelemons/goat/term"
)

func Select(list ...string) int {
	current := 0

	const (
		UP    = "\x1b\x5b\x41"
		DOWN  = "\x1b\x5b\x42"
		ENTER = "\r"
	)

	for i, line := range list {
		if i == current {
			ct.ChangeColor(ct.Black, false, ct.White, false)
		} else {
			ct.ResetColor()
		}
		fmt.Print(line)
		ct.ResetColor()
		if i < len(list)-1 {
			fmt.Print("\r\n")
		}
	}

	for i := range list {
		if i > 0 {
			fmt.Print(UP)
		}
	}

	raw := make([]byte, 10)
	tty := term.NewRawTTY(os.Stdin)
	tty.SetEcho(nil)
	for {
		n, err := tty.Read(raw)
		if err != nil {
			fmt.Printf("read: %s\n", err)
			return -1
		}
		str := string(raw[:n])
		switch str {
		case term.Interrupt:
			for i := current; i < len(list); i++ {
				fmt.Print("\n")
			}
			fmt.Print("\r")
			return -1
		case UP:
			if current > 0 {
				fmt.Print("\r")
				fmt.Print(list[current])

				current--
				fmt.Print(str)

				fmt.Print("\r")
				ct.ChangeColor(ct.Black, false, ct.White, false)
				fmt.Print(list[current])
				ct.ResetColor()
			}
		case DOWN:
			if current < len(list)-1 {
				fmt.Print("\r")
				fmt.Print(list[current])

				current++
				fmt.Print(str)

				fmt.Print("\r")
				ct.ChangeColor(ct.Black, false, ct.White, false)
				fmt.Print(list[current])
				ct.ResetColor()
			}
		case ENTER:
			for i := current; i < len(list); i++ {
				fmt.Print("\n")
			}
			fmt.Print("\r")
			return current
		}
	}
}

func main() {
	list := os.Args[1:]
	if len(list) == 0 {
		return
	}
	os.Exit(Select(list...))
}
