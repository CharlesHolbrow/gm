package main

import (
	"encoding/json"
	"math/rand"
	"os"
)

// Glyphs is a map collection of glyphs by name string
type Glyphs struct {
	byName map[string][]byte
	byInt  [][]byte
}

// ReadGlyphsDotJSON a file 'glyphs.json` and return a collectino of glyphs
func ReadGlyphsDotJSON() (Glyphs, error) {
	glyphs := Glyphs{
		byName: make(map[string][]byte, 32),
	}

	file, err := os.Open("glyphs.json")
	if err != nil {
		return glyphs, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&glyphs.byName)
	if err != nil {
		return glyphs, err
	}

	// Make the array for storing glyphs by integer. Do not include a glyph named
	// empty (if it exists)
	length := len(glyphs.byName)
	if _, ok := glyphs.byName["empty"]; ok {
		length--
	}

	glyphs.byInt = make([][]byte, length)
	index := 0
	for name, bytes := range glyphs.byName {
		if name != "empty" {
			glyphs.byInt[index] = bytes
			index++
		}
	}

	return glyphs, nil
}

// Get a glyph with it's Name
func (g Glyphs) Get(name string) []byte {
	data, ok := g.byName[name]
	if !ok {
		return nil
	}
	copied := make([]byte, len(data))
	copy(copied, data)
	return copied
}

// GetByInt returns a glyph indexed by it's interger.
func (g Glyphs) GetByInt(i int) []byte {
	if i >= len(g.byInt) || i < 0 {
		return g.Get("empty")
	}
	data := g.byInt[i]
	copied := make([]byte, len(data))
	copy(copied, data)
	return copied
}

// Length returns thenumber of Glyphs
func (g Glyphs) Length() int {
	return len(g.byInt)
}

// Random gets a random Glpph
func (g Glyphs) Random() []byte {
	return g.GetByInt(rand.Intn(len(g.byInt)))
}
