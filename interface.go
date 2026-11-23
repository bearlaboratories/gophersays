package main

import (
	"context"
	"machine"
	"time"
)

type colourType int

type colourEvent struct {
	typ  colourType
	down bool
}

type colours []*colour

func newColourSet(colours ...*colour) (set colours) {
	for _, v := range colours {
		set = append(set, v)
	}

	return
}

func (c *colours) getColour(typ colourType) (colour *colour) {
	for _, v := range *c {
		if v.typ == typ {
			colour = v
		}
	}

	if colour == nil {
		panic("color does not exist!")
	}

	return
}

func (c *colours) waitForAnyButton(down bool) colourType {
	ctx, cf := context.WithCancel(context.Background())
	ch := make(chan colourEvent)
	for _, v := range *c {
		go func() {
			for {
				select {
				case event := <-v.buttonCh:
					if event.down != down {
						continue
					}
					select {
					case ch <- event:
					case <-ctx.Done():
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	for {
		event := <-ch
		if event.down == down {
			cf()
			return event.typ
		}
	}
}

func (c *colours) waitForAll(down bool) {
	for {
		ok := true
		for _, v := range *c {
			if v.button.Get() != !down {
				ok = false
				break
			}
		}

		if ok {
			return
		}

		c.waitForAnyButton(down)
	}
}

func (c *colours) waitForColour(colour colourType, down bool) {
	col := c.getColour(colour)
	for {
		event := <-col.buttonCh
		if event.down == down {
			return
		}
	}
}

func (c *colours) lightsOn() {
	for _, v := range *c {
		v.led.High()
	}
}

func (c *colours) lightsOff() {
	for _, v := range *c {
		v.led.Low()
	}
}

func (c *colours) lightOn(t colourType) {
	for _, v := range *c {
		if v.typ == t {
			v.led.High()
		}
	}
}

func (c *colours) lightOff(t colourType) {
	for _, v := range *c {
		if v.typ == t {
			v.led.Low()
		}
	}
}

type colour struct {
	typ      colourType
	led      machine.Pin
	button   machine.Pin
	buttonCh chan colourEvent
}

func newColour(typ colourType, led machine.Pin, button machine.Pin) *colour {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	ch := make(chan colourEvent, 1)
	/*
		err := button.SetInterrupt(machine.PinToggle, func(pin machine.Pin) {
			select {
				case _ = <-ch:
				default:
				}
				select {
				case ch <- colourEvent{
					typ: typ,
					down:       !pin.Get(),
				}:
				default:
			}
		})

		if err != nil {
			panic(err)
		}
	*/

	// I did not learn that the ESP-32's interrupts doesn't work with TinyGo until way too late.
	// Gotta do this fucking dogshit instead.

	go func() {
		var last *bool
		for {
			high := button.Get()
			if last == nil || *last != high {
				select {
				case _ = <-ch:
				default:
				}

				select {
				case ch <- colourEvent{
					typ:  typ,
					down: !button.Get(),
				}:
				default:
				}

				last = &high
			}

			time.Sleep(1 * time.Millisecond)
		}
	}()

	return &colour{
		typ:      colourType(typ),
		led:      led,
		button:   button,
		buttonCh: ch,
	}
}
