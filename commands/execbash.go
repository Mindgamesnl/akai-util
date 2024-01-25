package commands

import (
	"akai-util/led"
	"fmt"
	"os/exec"
)

func Script(workingDir, scriptName string) func(int64, int64) {
	return func(note int64, channel int64) {
		led.StartBlinking(channel, note)

		// execute the script
		fmt.Println("Executing script: ", scriptName, workingDir)

		cmd := exec.Command("bash", scriptName)
		cmd.Dir = workingDir
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing script: ", err)
		}
		fmt.Println(string(out))
		led.StopBlinking(channel, note)
	}
}
