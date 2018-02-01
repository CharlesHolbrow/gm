package main

import "fmt"

// Note on or Note off events
type Note struct {
	Ch   uint8
	Note uint8
	Vel  uint8
	On   bool
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

func parse0xF(nib2 uint8, d1, d2 int64) interface{} {
	switch nib2 {
	case 0x0:
		return "System Exclusive"
	case 0x1:
		return "System Common  - undefined	?	?"
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
		return "Sys real time undefined "
	case 0xA:
		return Start{}
	case 0xB:
		return Continue{}
	case 0xC:
		return Stop{}
	case 0xD:
		return "Sys real time undefined "
	case 0xE:
		return "Sys real time active sensing "
	case 0xF:
		return "Sys real time sys reset"
	default:
		return "unknown"
	}
}
