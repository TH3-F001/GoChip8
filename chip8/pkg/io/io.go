package io

// Why combine display and input into one? because most libraries typically handle bot in a semi coupled manner.
// SDL requires a window to take scan codes, and TCell handles keyboard events itself, not externally
// at the end of the day, code doesnt always reflect reality.

type IO interface {
	GetPixels() *[][]byte
	GetPixel(col byte, row byte) byte
	SetPixel(col byte, row byte, lit byte) error
	Refresh() error
	Listen() (byte, error)
	Terminate() error
}
