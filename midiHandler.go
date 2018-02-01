package main

import (
	"fmt"
	"time"
)

const blurCC = 33 // when we receive a cc ask client to blur
const str1CC = 38
const stretchChannel = 1 // lowest four notes on this channel engage stretcher.

type MidiHandler interface {
	handleNote(Note)
	handleMisc(interface{})
	handleCC(CC)
	handlePW(PitchWheel)
}

type midiHandler struct {
	keys [16][128]uint8
	ccs  [16][128]uint8
	// how many notes are on on this channel?
	isPlaying         bool
	sixteenthsElapsed int

	// There are 24 ticks per quarter, so there are 6 per 16th. beatPosition
	beatPosition int

	start time.Time
}

func makeMidiHandler() *midiHandler {
	mh := &midiHandler{
		sixteenthsElapsed: 0, // is this even needed?
	}
	return mh
}

func (mh *midiHandler) handleNote(n Note) {
	if n.Vel == 0 {
		n.On = false
	}
	notes := mh.keys[int(n.Ch)]

	// Update our note maps maps
	if n.On {
		notes[n.Note] = n.Vel
	} else {
		notes[n.Note] = 0
	}
	fmt.Printf("%s - %d.%d (%v)\n", n, mh.sixteenthsElapsed, mh.beatPosition, time.Since(mh.start))
}

// Handle sync messages inclueing start, stop, continue, spp
func (mh *midiHandler) handleMisc(event interface{}) {
	switch t := event.(type) {
	case Clock:
		if !mh.isPlaying {
			break
		}
		mh.beatPosition = (mh.beatPosition + 1) % 6
		if mh.beatPosition == 0 {
			mh.sixteenthsElapsed++
			mh.onSixteenth()
		}
		if mh.sixteenthsElapsed%16 == 0 {
			mh.onWhole()
		}
	case SPP:
		mh.beatPosition = 0
		mh.sixteenthsElapsed = int(t) // increment on the next tick
		// When we move the cursor in reaper, reaper smartly sends note off events
		// for active notes in midi items. However, it does not send note off
		// events for notes held on the keyboard.
	case Start:
		mh.start = time.Now()
		mh.beatPosition = 0
		mh.sixteenthsElapsed = 0
		mh.isPlaying = true
	case Continue:
		mh.isPlaying = true
	case Stop:
		mh.isPlaying = false
	default:
		fmt.Println("midiHandler.handleMisc received", t)
	}
	fmt.Printf("%s - %d.%d (%v)\n", event, mh.sixteenthsElapsed, mh.beatPosition, time.Since(mh.start))
}

func (mks *midiHandler) handleCC(cc CC) {
	mks.ccs[cc.Ch][cc.Number] = cc.Value
	switch cc.Number {
	case 64:
		if cc.Value >= 64 {
			// pedal down
		} else {
			// pedal up
		}
	}
}

func (mks *midiHandler) onSixteenth() {
}

func (mh *midiHandler) onWhole() {
}

func (mks *midiHandler) handlePW(pw PitchWheel) {
}
