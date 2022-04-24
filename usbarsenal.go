package main

import (
	"time"

	usbarmory "github.com/usbarmory/tamago/board/f-secure/usbarmory/mark-two"
)

func main() {
	for {
		usbarmory.LED("white", true)
		time.Sleep(time.Second)
		usbarmory.LED("white", false)
		time.Sleep(time.Second)
	}
}
