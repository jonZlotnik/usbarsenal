# usbarsenal

`make CROSS_COMPILE=arm-none-eabi- TARGET=usbarmory imx`
`sudo $HOME/go/bin/armory-boot-usb -i example.imx`

## USB Interface Scraping w/ Wireshark on Manjaro:
`modprobe usbmon`

Filter like this:
`usb.device_address==11`

## View USB device from OS:
  
Watch OS-level USB device logs:
`sudo dmesg -wH | grep usb`

Check USB device structure/validity with:
`sudo pacman -S usbview`

## Reference material

https://hikalium.github.io/os_dev_specs/

and the /notes directory