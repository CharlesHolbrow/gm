package main

// MidiHandler represents any device that listens for midi events
type MidiHandler interface {
	HandleNote(Note)
	HandleMisc(interface{})
	HandleCC(CC)
	HandlePW(PitchWheel)
}
