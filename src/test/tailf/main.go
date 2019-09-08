package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		// Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile("./logs/test.log", config)
	if err != nil {
		fmt.Println("init tail failed, err: ", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("file close reopen filename: %s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		fmt.Println("msg: ", msg.Text)
	}
}
