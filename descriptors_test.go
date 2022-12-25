package main

import (
	"encoding/hex"
	"testing"

	imxusb "github.com/usbarmory/tamago/soc/imx6/usb"
)

func Test_configureKeyboardDevice(t *testing.T) {
	device := &imxusb.Device{}
	configureKeyboardDevice(device)
	descBytes := device.Descriptor.Bytes()
	t.Log(hex.EncodeToString(descBytes))
}
