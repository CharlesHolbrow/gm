package main

import (
	"fmt"
)

// Note on or Note off events
type Note struct {
	Ch   uint8
	Note uint8
	Vel  uint8
	On   bool
}

// Midi return the correct arguments to pass in to portmidi.WriteShort
func (n Note) Midi() (status, b1, b2 int64) {
	if n.On {
		return 0x90 + int64(n.Ch%16), int64(n.Note), int64(n.Vel)
	}
	return 0x80 + int64(n.Ch%16), int64(n.Note), int64(n.Vel)
}

func (n Note) String() string {
	if n.On {
		return fmt.Sprintf("Note On: %d Vel: %d Channel(%x)", n.Note, n.Vel, n.Ch)
	}
	return fmt.Sprintf("Note Off: %d Vel: %d Channel(%x)", n.Note, n.Vel, n.Ch)
}

// CC Event
type CC struct {
	Ch     uint8
	Number uint8
	Value  uint8
}

// Midi return the correct arguments to pass in to portmidi.WriteShort
func (cc CC) Midi() (status, b1, b2 int64) {
	return 0xB0 + int64(cc.Ch%16), int64(cc.Number), int64(cc.Value)
}

func (cc CC) String() string {
	return fmt.Sprintf("CC Number: %d Value: %d Channel(%x)", cc.Number, cc.Value, cc.Ch)
}

// SPP represents a midi song position pointer event. Contains the number of
// 16th notes since the beginning of the song
type SPP int

func (spp SPP) String() string {
	return fmt.Sprintf("Song Position (16th notes): %d", spp)
}

// Clock represents a midi clock tick event
type Clock struct{}

func (c Clock) String() string {
	return "Midi Clock Tick"
}

// Start represents a midi start event
type Start struct{}

func (s Start) String() string {
	return "Midi Start"
}

// Continue represents a midi continue event
type Continue struct{}

func (c Continue) String() string {
	return "Midi Continue"
}

// Stop represents a midi stop event
type Stop struct{}

func (s Stop) String() string {
	return "Midi Stop"
}

// PitchWheel event
type PitchWheel struct {
	Ch    uint8
	Value int
}

func (pw PitchWheel) String() string {
	return fmt.Sprintf("Pitch Wheel %d Channel(%x)", pw.Value, pw.Ch)
}

// "System Common" and "Realtime" MIDI Messages do not specify a channel. To
// parse these messages, we split the status byte into two "nibbles". The
// First nibble is 0xF, the second nibble indicates the type of message AND
// the number of bytes to expect before the next status byte. For example,
// a "Song Position Pointer" (SPP) will be in the format 0xF2 0xNN 0xNN.
// portmidid handles converting the three bytes into an "event" for us, but it
// is our job to convert the two 7 bit values into a meaningful representation.
func parse0xF(nib2 uint8, d1, d2 int64) interface{} {
	switch nib2 {
	case 0x0:
		return "System Exclusive"
	case 0x1:
		return "System Common - MIDI Time Code Quarter Frame."
	case 0x2:
		return SPP(d2*127 + d1) // 16th notes since beginning
	case 0x3:
		return "Sys Commmon Song Select(Song #)	(0-127)	NONE"
	case 0x4:
		return "System Common - undefined	?	?"
	case 0x5:
		return "System Common - undefined	?	?"
	case 0x6:
		return "Sys Common Tune Request	NONE"
	case 0x7:
		return "Sys Commmon End of SysEx (EOX)"
	case 0x8:
		return Clock{}
	case 0x9:
		return "Sys real time undefined"
	case 0xA:
		return Start{}
	case 0xB:
		return Continue{}
	case 0xC:
		return Stop{}
	case 0xD:
		return "Sys real time undefined"
	case 0xE:
		return "Sys real time active sensing"
	case 0xF:
		return "Sys real time sys reset"
	default:
		return "unknown"
	}
}
