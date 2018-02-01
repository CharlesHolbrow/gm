package gm

import (
	"fmt"

	"github.com/rakyll/portmidi"
)

// PMInfoToString converts portmidi.DeviceInto into a Human Readable string
func PMInfoToString(info *portmidi.DeviceInfo) string {
	var ioType string
	if info.IsInputAvailable && info.IsOutputAvailable {
		ioType = "In/Output Device"
	} else if info.IsInputAvailable {
		ioType = "Input Device "
	} else if info.IsOutputAvailable {
		ioType = "Output Device"
	} else {
		ioType = "No I/O       "
	}
	return fmt.Sprintf("%s - %s - %s", ioType, info.Interface, info.Name)
}

// PrintDevices enumerates available midi devices to the stdout
func PrintDevices() {
	deviceCount := portmidi.CountDevices()
	fmt.Printf("Found %d midi devices\n", deviceCount)
	for i := 0; i < deviceCount; i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		fmt.Println(i, "-", PMInfoToString(info))
	}
	fmt.Println()
}
