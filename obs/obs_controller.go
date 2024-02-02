package obs

import (
	"akai-util/led"
	"akai-util/utils"
	"fmt"
	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/andreykaipov/goobs/api/typedefs"
	"github.com/rakyll/portmidi"
	"sort"
	"time"
)

var client *goobs.Client
var loadedScenes []typedefs.Scene
var currentSceneIndex = 0
var midiStream *portmidi.Stream

func Init(s *portmidi.Stream) {
	midiStream = s
	go func() {
		for {
			if client == nil {
				craetedClient, err := goobs.New("localhost:4469")
				if err == nil {
					fmt.Println("Reconnected to OBS")
					client = craetedClient
					attemtSceneReload()
					led.StartBlinking(0, 48)
				}
			}
			time.Sleep(15 * time.Second)
		}
	}()

	go func() {
		for {
			if client != nil {
				attemtSceneReload()
			}
			time.Sleep(10 * time.Second)
		}
	}()

	for i := 1; i < 8; i++ {
		fmt.Println("Registering scene ", i)
		utils.RegisterMidiNoteOn(50, byte(i), func(note int64, channel int64) {
			defer func() {
				if recover() != nil {
					fmt.Println("Error while changing scene")
				}
			}()

			currentSceneIndex = int(channel) - 1
			fmt.Println("Changing scene to ", currentSceneIndex, " ", loadedScenes[currentSceneIndex].SceneName)

			if client != nil {
				client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{SceneName: &loadedScenes[currentSceneIndex].SceneName})
			}
			renderScene(currentSceneIndex+1, false)
		})
	}

	renderScene(0, false)
}

func renderScene(sceneIndex int, inverted bool) {
	if sceneIndex < 1 {
		sceneIndex = 1
	}
	if sceneIndex > 7 {
		sceneIndex = 7
	}
	for i := 0; i < 7; i++ {
		var i64 int64 = int64(i)
		var on = i == sceneIndex
		if inverted {
			on = !on
		}
		led.SetLed(midiStream, i64, 50, 64, on)
	}
}

func attemtSceneReload() {
	defer func() {
		// recover from panic if one occurred. Set err to nil otherwise.
		if recover() != nil {
			fmt.Println("Error while reloading scenes, killing connection")
			client = nil
			led.StopBlinking(0, 48)
			loadedScenes = []typedefs.Scene{}
			disconnectButIgnoreError()
		}
	}()
	loadedScenes = ReloadScenes()
}

func disconnectButIgnoreError() {
	defer func() {
		// recover from panic if one occurred. Set err to nil otherwise.
		if recover() != nil {
			fmt.Println("Error while disconnecting, ignoring")
		}
	}()
	if client != nil {
		client.Disconnect()
	}
}

func ReloadScenes() []typedefs.Scene {
	respsone, err := client.Scenes.GetSceneList()
	if err != nil {
		return []typedefs.Scene{}
	}

	var scenes []typedefs.Scene

	for i := range respsone.Scenes {
		scenes = append(scenes, *respsone.Scenes[i])
	}

	// sort ascending by sceneindex
	sort.Slice(scenes, func(i, j int) bool {
		return scenes[i].SceneIndex > scenes[j].SceneIndex
	})

	fmt.Println("Loaded scenes: ", len(scenes))
	return scenes
}
