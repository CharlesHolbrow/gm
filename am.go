package main

import (
	"fmt"
	"log"

	"sync"

	"github.com/rakyll/portmidi"
)

const midiInputDeviceID = 0

func main() {
	// Star portmidi and print devices
	portmidi.Initialize()
	deviceCount := portmidi.CountDevices()
	fmt.Printf("Found %d midi devices\n", deviceCount)
	for i := 0; i < deviceCount; i++ {
		info := portmidi.Info(portmidi.DeviceID(i))
		fmt.Println(i, "-", info)
	}

	fmt.Printf("Listening for MIDI on %d - %v\n", midiInputDeviceID, *portmidi.Info(midiInputDeviceID))
	myMidiHandler := makeMidiHandler()
	ms, err := MakeMidiStream(midiInputDeviceID, myMidiHandler)
	if err != nil {
		log.Panicln("Error Making Midi Stream", err)
	}
	defer ms.Close()
	g := sync.WaitGroup{}
	g.Add(1)
	g.Wait()

}
