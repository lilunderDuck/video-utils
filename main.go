package main

import (
	"os"
	"video-utils/cmd"
	"video-utils/tools"
	"video-utils/utils"
)

func main() {
	utils.CheckIfFFmpegExist()
	os.Mkdir("out", 0777)

	stuff := cmd.New(&cmd.VideoToolsMapping{
		"split-video": tools.SplitVideo,
		"merge-video": tools.MergeVideo,
	})
	if stuff == nil {
		return
	}

	stuff.Run()
}
