package main

import (
	"fmt"
	"log"
	"time"

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
	myMidiHandler := NewMidiLogger()
	ms, err := MakeMidiStream(midiInputDeviceID, myMidiHandler)
	if err != nil {
		log.Panicln("Error Making Midi Stream", err)
	}
	defer ms.Close()

	out, err := portmidi.NewOutputStream(2, 1024, 0)
	if err != nil {
		log.Fatal(err)
	}
	out.WriteShort(Note{On: true, Note: 64, Vel: 127}.Midi())
	time.Sleep(time.Second)
	out.WriteShort(Note{Note: 64, Vel: 127}.Midi())

	g := sync.WaitGroup{}
	g.Add(1)
	g.Wait()

}
