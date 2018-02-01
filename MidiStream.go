package main

import (
	"fmt"

	"github.com/rakyll/portmidi"
)

// MidiStream wraps a portmidi.Stream
type MidiStream struct {
	pmStream    *portmidi.Stream // portmidi stream
	midiHandler MidiHandler
}

// MakeMidiStream creates a MidiStream object that is listening to a portmidi
// device identified by deviceID.
func MakeMidiStream(deviceID int, handler MidiHandler) (*MidiStream, error) {
	var ms *MidiStream

	in, err := portmidi.NewInputStream(portmidi.DeviceID(deviceID), 1024)
	if err != nil {
		return ms, err
	}

	ms = &MidiStream{
		pmStream:    in,
		midiHandler: handler,
	}

	go func() {
		for event := range in.Listen() {
			ms.handlePortmidiEvent(event)
		}
	}()

	return ms, nil
}

// Close and stop piping input to channels
func (ms MidiStream) Close() {
	ms.pmStream.Close()
}

// Send pormidi events to their appropriate go channels
func (ms MidiStream) handlePortmidiEvent(e portmidi.Event) {
	nib1 := uint8((e.Status & 0xF0) >> 4)
	nib2 := uint8(e.Status & 0x0F)
	switch nib1 {
	case 0x8:
		ms.midiHandler.handleNote(Note{On: false, Ch: nib2, Note: uint8(e.Data1), Vel: uint8(e.Data2)})
	case 0x9:
		ms.midiHandler.handleNote(Note{On: true, Ch: nib2, Note: uint8(e.Data1), Vel: uint8(e.Data2)})
	case 0xE:
		ms.midiHandler.handlePW(PitchWheel{Ch: nib2, Value: int(e.Data2*127 + e.Data1 - 8128)})
	case 0xF:
		ms.midiHandler.handleMisc(parse0xF(nib2, e.Data1, e.Data2))
	case 0xB:
		ms.midiHandler.handleCC(CC{Ch: nib2, Number: uint8(e.Data1), Value: uint8(e.Data2)})
	default:
		fmt.Println("Warning. Unhandled Midi event:", e)
	}
}
