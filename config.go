package main

import (
	"machine"
)

const gameLength = 12 // Amount of lights until the game is won

// Define your colour types
const (
	ylw colourType = iota
	red
	green
)

// Aliases
var (
	ylwLED   machine.Pin = machine.GPIO13
	redLED               = machine.GPIO12
	greenLED             = machine.GPIO27

	ylwButton   = machine.GPIO14
	redButton   = machine.GPIO26
	greenButton = machine.GPIO25
)

// Add colours, voltage pulldown is expected for buttons
var set = newColourSet(
	newColour(ylw, ylwLED, ylwButton),
	newColour(red, redLED, redButton),
	newColour(green, greenLED, greenButton),
)

// This global is directly used for audio
var buzzerPin = machine.GPIO33
