package main

import (
	"time"

	_ "github.com/usbarmory/crucible/fusemap"
	usbarmory "github.com/usbarmory/tamago/board/usbarmory/mk2"
	imxusb "github.com/usbarmory/tamago/soc/imx6/usb"
)

func main() {
	port := usbarmory.USB1
	device := &imxusb.Device{}
	configureKeyboardDevice(device)
	port.Init()
	port.DeviceMode()
	port.Reset()

	go port.Start(device)
	for {
		usbarmory.LED("white", true)
		time.Sleep(time.Second)
		usbarmory.LED("white", false)
		time.Sleep(time.Second)
	}
}

func configureKeyboardDevice(device *imxusb.Device) {
	// Supported Language Code Zero: English
	device.SetLanguageCodes([]uint16{0x0409})

	// device descriptor
	device.Descriptor = &imxusb.DeviceDescriptor{}
	device.Descriptor.SetDefaults()
	device.Descriptor.DescriptorType = 0x01
	device.Descriptor.DeviceClass = 0x00
	device.Descriptor.DeviceSubClass = 0x00
	device.Descriptor.DeviceProtocol = 0x00
	// http://pid.codes/1209/2702/
	device.Descriptor.VendorId = 0x1209
	device.Descriptor.ProductId = 0x2702
	device.Descriptor.Device = 0x0001
	iManufacturer, _ := device.AddString(`JonnyZ123`)
	device.Descriptor.Manufacturer = iManufacturer
	iProduct, _ := device.AddString(`Jon's Keyboard`)
	device.Descriptor.Product = iProduct
	iSerial, _ := device.AddString(`0.1`)
	device.Descriptor.SerialNumber = iSerial

	// configuration descriptor
	conf := &imxusb.ConfigurationDescriptor{}
	conf.SetDefaults()
	conf.DescriptorType = 0x02
	conf.TotalLength = 0x003B
	conf.NumInterfaces = 0x00      // changed 0 to 1 -> caused "numInterfaces to become 2" and malformed packet
	conf.ConfigurationValue = 0x00 // changed 1 to 0 -> got rid of USB version FIXME in wireshark
	conf.Configuration = 0x00
	// conf.Attributes = 0b10100000

	// -- interface descriptor (keyboard)
	keyboardInterface := &imxusb.InterfaceDescriptor{}
	keyboardInterface.SetDefaults()
	keyboardInterface.DescriptorType = 0x04
	keyboardInterface.InterfaceNumber = 0x00
	keyboardInterface.AlternateSetting = 0x00
	keyboardInterface.NumEndpoints = 0x02
	keyboardInterface.InterfaceClass = 0x03
	keyboardInterface.InterfaceSubClass = 0x01
	keyboardInterface.InterfaceProtocol = 0x01
	keyboardInterface.Interface = 0x00

	// ----- class descriptor
	keyboardClassDescriptor := HIDDescriptor{
		Length:               0x09,
		DescriptorType:       0x05,
		HID:                  0x101,
		CountryCode:          0x00,
		NumDescriptors:       0x01,
		ReportDescriptorType: 0x22,
		DescriptorLength:     0x3f,
	}
	keyboardEndpointDescriptor := imxusb.EndpointDescriptor{
		Length:          0x07,
		DescriptorType:  0x05,
		EndpointAddress: 0b10000001,
		Attributes:      0b00000011,
		MaxPacketSize:   0x0008,
		Interval:        0x0A,
		Function:        HIDTx,
	}

	// Put the descriptors together
	keyboardInterface.ClassDescriptors = append(keyboardInterface.ClassDescriptors, keyboardClassDescriptor.Bytes())
	keyboardInterface.Endpoints = append(keyboardInterface.Endpoints, &keyboardEndpointDescriptor)
	conf.AddInterface(keyboardInterface)
	device.AddConfiguration(conf)

	// device qualifier
	device.Qualifier = &imxusb.DeviceQualifierDescriptor{}
	device.Qualifier.SetDefaults()
	device.Qualifier.NumConfigurations = uint8(len(device.Configurations))
}
