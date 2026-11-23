package main

import (
	"time"

	"tinygo.org/x/drivers/buzzer"
)

func PlayEndMusic() {
	buzz.Tone(buzzer.C4, 0.3)
	buzz.Tone(buzzer.E4, 0.3)
	buzz.Tone(buzzer.G4, 0.3)
	buzz.Tone(buzzer.C5, 0.3)
	buzz.Tone(buzzer.E5, 0.3)
	buzz.Tone(buzzer.G5, 0.3)
	buzz.Tone(buzzer.C4, 0.3)
	buzz.Tone(buzzer.Eb4, 0.3)
	buzz.Tone(buzzer.Ab4, 0.3)
	buzz.Tone(buzzer.C5, 0.3)
	buzz.Tone(buzzer.Eb5, 0.3)
	buzz.Tone(buzzer.Ab5, 0.3)
	buzz.Tone(buzzer.D4, 0.3)
	buzz.Tone(buzzer.F4, 0.3)
	buzz.Tone(buzzer.Bb4, 0.3)
	buzz.Tone(buzzer.D5, 0.3)
	buzz.Tone(buzzer.F5, 0.3)
	buzz.Tone(buzzer.Bb5, 0.3)
	buzz.Tone(buzzer.Bb5, 0.2)
	time.Sleep(100 * time.Millisecond)
	buzz.Tone(buzzer.Bb5, 0.2)
	time.Sleep(100 * time.Millisecond)
	buzz.Tone(buzzer.Bb5, 0.2)
	buzz.Tone(buzzer.C6, 1)
}

func PlayLose() {
	buzz.Tone(buzzer.B4, 0.3)
	buzz.Tone(buzzer.F5, 0.3)
	time.Sleep(300 * time.Millisecond)
	buzz.Tone(buzzer.F5, 0.3)
	buzz.Tone(buzzer.F5, 0.3)
	buzz.Tone(buzzer.E5, 0.3)
	buzz.Tone(buzzer.D5, 0.3)
	time.Sleep(100 * time.Millisecond)
	buzz.Tone(buzzer.D5, 0.2)
	time.Sleep(50 * time.Millisecond)
	buzz.Tone(buzzer.C5, 0.2)
	time.Sleep(300 * time.Millisecond)
	buzz.Tone(buzzer.E4, 0.3)
	buzz.Tone(buzzer.C4, 0.9)
}
