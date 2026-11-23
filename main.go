package main

import (
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

var buzz = buzzer.New(buzzerPin)

func init() {
	buzzerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

const (
	Start = iota
	Playing
	Listening
	Failed
	Won
)

var ticker = time.NewTicker(1 * time.Second)

func GeneratePattern(set colours, length int) []colourType {
	var types []colourType
	for _, v := range set {
		types = append(types, v.typ)
	}

	pat := make([]colourType, length)
	for i := range pat {
		pat[i] = colourType(types[rand.Intn(len(types))])
	}

	return pat
}

func main() {
	var buttons []colourType
	var progress int
	var responseIndex int
	state := Start

	for {
		switch state {
		case Start:
			waitCh := make(chan struct{})
			go func() {
				set.waitForAnyButton(true)
				var send struct{}
				waitCh <- send
			}()

			for {
				select {
				case time := <-ticker.C:
					if time.Unix()%2 == 1 {
						set.lightsOn()
					} else {
						set.lightsOff()
					}
				case <-waitCh:
					buttons = GeneratePattern(set, gameLength)
					progress = 0
					responseIndex = 0

					set.lightsOff()
					state = Playing
				}

				if state != Start {
					break
				}
			}
		case Playing:
			set.waitForAll(false)
			time.Sleep(500 * time.Millisecond)
			for i, v := range buttons {
				set.lightOn(v)
				time.Sleep(200 * time.Millisecond)
				set.lightOff(v)
				time.Sleep(200 * time.Millisecond)

				if i == progress {
					state = Listening
					break
				}

			}
			set.waitForAll(false)
		case Listening:
			go func() { buzz.Tone(buzzer.F5, 0.3) }()
			for {
				button := set.waitForAnyButton(true)

				if button != buttons[responseIndex] {
					state = Failed
					break
				} else if responseIndex == (len(buttons) - 1) {
					state = Won
					break
				} else if responseIndex == progress {
					responseIndex = 0
					progress++
					state = Playing
					break
				} else {
					responseIndex++
				}
			}
		case Failed:
			go func() {
				PlayLose()
			}()
			for range 8 {
				set.lightsOff()
				time.Sleep(200 * time.Millisecond)
				set.lightsOn()
				time.Sleep(200 * time.Millisecond)
			}

			state = Start
		case Won:
			go func() {
				PlayEndMusic()
			}()
			for range 15 {
				for _, v := range set {
					set.lightsOff()
					v.led.High()
					time.Sleep(100 * time.Millisecond)
				}
			}

			state = Start
		}
	}
}
