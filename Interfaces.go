package main

// MidiHandler represents any device that listens for midi events
type MidiHandler interface {
	handleNote(Note)
	handleMisc(interface{})
	handleCC(CC)
	handlePW(PitchWheel)
}
