package utils

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/sqp/pulseaudio"
)

type PAClient struct {
	*pulseaudio.Client

	playbackStreamsByName map[string][]dbus.ObjectPath
	recordStreamsByName   map[string][]dbus.ObjectPath
	sourcesByName         map[string][]dbus.ObjectPath
	sinksByName           map[string][]dbus.ObjectPath
}

func NewPAClient(c *pulseaudio.Client) *PAClient {
	client := &PAClient{
		Client:                c,
		playbackStreamsByName: make(map[string][]dbus.ObjectPath, 0),
		recordStreamsByName:   make(map[string][]dbus.ObjectPath, 0),
		sourcesByName:         make(map[string][]dbus.ObjectPath, 0),
		sinksByName:           make(map[string][]dbus.ObjectPath, 0),
	}
	return client
}

func (c *PAClient) NewPlaybackStream(path dbus.ObjectPath) {
	c.RefreshStreams()
}

func (c *PAClient) PlaybackStreamRemoved(path dbus.ObjectPath) {
	c.RefreshStreams()
}

func (c *PAClient) RefreshStreams() error {
	playbackStreamsByName := make(map[string][]dbus.ObjectPath, 0)
	recordStreamsByName := make(map[string][]dbus.ObjectPath, 0)
	sinksByName := make(map[string][]dbus.ObjectPath, 0)
	sourcesByName := make(map[string][]dbus.ObjectPath, 0)

	streams, err := c.Core().ListPath("PlaybackStreams")
	if err != nil {
		return err
	}

	for _, streamPath := range streams {
		stream := c.Stream(streamPath)
		props, err := stream.MapString("PropertyList")
		if err != nil {
			return err
		}

		if applicationName, ok := props["application.name"]; ok {
			if _, ok := playbackStreamsByName[applicationName]; ok {
				playbackStreamsByName[applicationName] = append(playbackStreamsByName[applicationName], streamPath)
			} else {
				playbackStreamsByName[applicationName] = []dbus.ObjectPath{streamPath}
			}
		}
	}

	streams, err = c.Core().ListPath("RecordStreams")
	if err != nil {
		return err
	}

	for _, streamPath := range streams {
		stream := c.Stream(streamPath)
		props, err := stream.MapString("PropertyList")
		if err != nil {
			return err
		}

		if applicationName, ok := props["application.name"]; ok {
			if _, ok := recordStreamsByName[applicationName]; ok {
				recordStreamsByName[applicationName] = append(recordStreamsByName[applicationName], streamPath)
			} else {
				recordStreamsByName[applicationName] = []dbus.ObjectPath{streamPath}
			}
		}
	}

	sinks, err := c.Core().ListPath("Sinks")
	if err != nil {
		return err
	}
	for _, sinkPath := range sinks {
		device := c.Device(sinkPath)
		props, err := device.MapString("PropertyList")
		if err != nil {
			panic(err)
		}

		if deviceDescription, ok := props["device.description"]; ok {
			if _, ok := sinksByName[deviceDescription]; ok {
				sinksByName[deviceDescription] = append(sinksByName[deviceDescription], sinkPath)
			} else {
				sinksByName[deviceDescription] = []dbus.ObjectPath{sinkPath}
			}
		}
	}

	sources, err := c.Core().ListPath("Sources")
	if err != nil {
		return err
	}
	for _, sourcePath := range sources {
		device := c.Device(sourcePath)
		props, err := device.MapString("PropertyList")
		if err != nil {
			panic(err)
		}

		if deviceDescription, ok := props["device.description"]; ok {
			if _, ok := sourcesByName[deviceDescription]; ok {
				sourcesByName[deviceDescription] = append(sourcesByName[deviceDescription], sourcePath)
			} else {
				sourcesByName[deviceDescription] = []dbus.ObjectPath{sourcePath}
			}
		}
	}

	c.playbackStreamsByName = playbackStreamsByName
	c.recordStreamsByName = recordStreamsByName
	c.sinksByName = sinksByName
	c.sourcesByName = sourcesByName
	return nil
}

func (c *PAClient) ProcessVolumeAction(appName string, volume float32) error {
	volume = volume / 100
	pa100perc := 65535
	newVol := uint32(volume * float32(pa100perc))

	objs := make([]*pulseaudio.Object, 0)

	if appName == "*" {
		for s := range c.playbackStreamsByName {
			var found = false
			for i := range trackedSinks {
				if trackedSinks[i] == s {
					found = true
				}
			}

			if !found {
				var paths = c.playbackStreamsByName[s]
				for _, streamPath := range paths {
					fmt.Println("Handling untracked "+s, c.Stream(streamPath))
					objs = append(objs, c.Stream(streamPath))
				}
			}
		}
	} else {
		if streamPaths, ok := c.playbackStreamsByName[appName]; ok {
			for _, streamPath := range streamPaths {
				objs = append(objs, c.Stream(streamPath))
			}
		}
	}

	if len(objs) > 0 {
		for _, obj := range objs {
			err := obj.Set("Volume", []uint32{newVol, newVol})
			if err != nil {
				return err
			}
		}
	} else {
		fmt.Printf("Could not find %s by name [%s] to set its volume", "stream", appName)
	}
	return nil
}
