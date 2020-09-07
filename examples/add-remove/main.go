package main

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/matti/dynamicmultiwriter"
)

func main() {
	pr1, pw1 := io.Pipe()
	pr2, pw2 := io.Pipe()

	dw := dynamicmultiwriter.New()
	dw.Add(pw1, pw2)
	cmd := exec.Command("ping", "-c", "10", "google.com")
	cmd.Stdout = dw

	go reader(1, pr1)
	go reader(2, pr2)

	go func() {
		time.Sleep(time.Second * 1)
		dw.Remove(pw2)
		time.Sleep(time.Second * 3)
		dw.Remove(pw1)
		dw.Add(pw2)
	}()

	cmd.Run()
}

func reader(i int, r io.Reader) {
	p := make([]byte, 4<<20)
	for {
		r.Read(p)
		str := string(p)
		if len(str) > 40 {
			str = str[:40]
		}
		fmt.Printf("r%d: %s\n", i, str)
	}
}
