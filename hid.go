package main

import (
	"bytes"
	"encoding/binary"
)

type HIDDescriptor struct {
	Length               uint8
	DescriptorType       uint8
	HID                  uint16
	CountryCode          uint8
	NumDescriptors       uint8
	ReportDescriptorType uint8
	DescriptorLength     uint16
}

func (d *HIDDescriptor) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, d)
	return buf.Bytes()
}

func HIDTx(_ []byte, lastErr error) (in []byte, err error) {
	resBytes := [8]byte{
		0, 0, 1, 0, 0, 1, 0, 0,
	}
	in = resBytes[:]
	err = nil
	return
}
