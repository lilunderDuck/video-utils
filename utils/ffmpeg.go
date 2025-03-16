package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func getFFmpegPath() string {
	return filepath.Join(GetCurrentDir(), "bin/ffmpeg.exe")
}

func CheckIfFFmpegExist() {
	ffmpegPath := getFFmpegPath()
	_, openError := os.Open(ffmpegPath)
	if openError != nil {
		panic(fmt.Sprintf("ffmpeg not found on %s", ffmpegPath))
	}
}

func RunFFmpeg(withArguments ...string) {
	ffmpegPath := getFFmpegPath()
	command := exec.Command(ffmpegPath, withArguments...)
	stdout, err := command.StdoutPipe()
	if err != nil {
		fmt.Printf("Failed to run ffmpeg. \n %s\n", err)
		return
	}

	command.Stderr = command.Stdout
	if err = command.Start(); err != nil {
		fmt.Printf("Failed to run ffmpeg. \n %s\n", err)
		return
	}

	for { // while true
		lineBuffer := make([]byte, 1024)
		_, err := stdout.Read(lineBuffer)
		fmt.Print(string(lineBuffer))
		if err != nil {
			break
		}
	}
}
