package led

import (
	"fmt"
	"github.com/rakyll/portmidi"
	"sync"
	"time"
)

type blinkTarget struct {
	channel, key  int64
	offThenRemove bool
}

var blinkingLedIds []blinkTarget
var lock sync.Mutex

func StartBlinking(channel, key int64) {
	lock.Lock()
	defer lock.Unlock()
	blinkingLedIds = append(blinkingLedIds, blinkTarget{channel, key, false})
}

func StopBlinking(channel, key int64) {
	lock.Lock()
	defer lock.Unlock()
	for i := range blinkingLedIds {
		var led = blinkingLedIds[i]
		if led.channel == channel && led.key == key {
			led.offThenRemove = true
			blinkingLedIds[i] = led
		}
	}
}

func StartBlinkLoop(s *portmidi.Stream) {
	go func() {
		var flip = false
		for {
			lock.Lock()
			var maxLights = len(blinkingLedIds)
			for i := range blinkingLedIds {
				var led = blinkingLedIds[i]
				if led.offThenRemove {
					fmt.Println("Removing blinking led", led.channel, led.key)
					SetLed(s, led.channel, led.key, 0, false)
					blinkingLedIds = append(blinkingLedIds[:i], blinkingLedIds[i+1:]...)
					i--
					// would this move i out of bounds?
					if i >= maxLights-1 {
						break
					}
					continue
				}
				SetLed(s, led.channel, led.key, 127, flip)
			}
			lock.Unlock()
			flip = !flip
			time.Sleep(150 * time.Millisecond)
		}
	}()
}
